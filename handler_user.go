package main

import (
	"Blogger/internal/database"
	"context"
	"database/sql"
	"errors"
	"fmt"
	"os"
	"time"

	"github.com/google/uuid"
)

func handlerGetUsers(s *state, cmd command) error {
	curr := s.cfg.CurrentUserName

	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("get users: %w", err)
	}

	for _, u := range users {
		line := "* " + u.Name
		if u.Name == curr {
			line += " (current)"
		}
		fmt.Println(line)
	}
	return nil
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
