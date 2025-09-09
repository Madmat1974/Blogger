package main

import (
	"Blogger/internal/config"
	"Blogger/internal/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"os"

	"github.com/google/uuid"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
}

func handlerRegister(s *state, cmd command) error {
	// go
	ctx := context.Background()
	username := cmd.args[0]
	_, err := s.db.GetUser(ctx, username)
	if err == nil {
		// found user
		fmt.Println("user already exists")
		os.Exit(1)
	}
	if !errors.Is(err, sql.ErrNoRows) {
		// unexpected error
		return err
	}

	// proceed to create user...

	// 2) create
	now := time.Now()
	_, err = s.db.CreateUser(context.Background(), database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: now,
		UpdatedAt: now,
		Name:      username,
	})
	if err != nil {
		return err
	}

	// 3) set current user
	if err := s.cfg.SetUser(username); err != nil {
		return err
	}

	fmt.Println("user created:", username)
	return nil
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.args) < 1 {
		return fmt.Errorf("need a single argument for username")
	}
	username := cmd.args[0]

	_, err := s.db.GetUser(context.Background(), username)
	if err != nil {
		return fmt.Errorf("user %q not found", username)
	}

	if err := s.cfg.SetUser(username); err != nil {
		return err
	}
	fmt.Println("user set to", username)
	return nil
}

type commands struct {
	handlers map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	if c.handlers == nil {
		c.handlers = make(map[string]func(*state, command) error)
	}
	c.handlers[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	h, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("unknown command: %s", cmd.name)
	}
	return h(s, cmd)
}
