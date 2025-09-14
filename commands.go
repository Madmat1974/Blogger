package main

import (
	"context"
	"fmt"
	"gator/internal/config"
	"gator/internal/database"
)

type state struct {
	db  *database.Queries
	cfg *config.Config
}

type command struct {
	name string
	args []string
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

// wrapper middleware function
func middlewareLoggedIn(
	handler func(s *state, cmd command, user database.User) error,
) func(s *state, cmd command) error {
	return func(s *state, cmd command) error {
		if s.cfg.CurrentUserName == "" {
			return fmt.Errorf("please login first")
		}

		user, err := s.db.GetUser(context.Background(), s.cfg.CurrentUserName)
		if err != nil {
			return fmt.Errorf("get user %q: %w", s.cfg.CurrentUserName, err)
		}
		return handler(s, cmd, user)
	}
}
