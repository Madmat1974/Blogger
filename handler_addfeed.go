package main

import (
	"Blogger/internal/database"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerAddFeed(s *state, cmd command) error {
	ctx := context.Background()

	if len(cmd.args) != 2 {
		return fmt.Errorf("usage: addfeed <name> <url>")
	}
	name := cmd.args[0]
	url := cmd.args[1]

	u, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("get user %q: %w", s.cfg.CurrentUserName, err)
	}

	id := uuid.New()

	f, err := s.db.CreateFeed(ctx, database.CreateFeedParams{
		ID:     id,
		Name:   name,
		Url:    url,
		UserID: u.ID,
	})
	if err != nil {
		return fmt.Errorf("create feed: %w", err)
	}

	//auto-following
	_, err = s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: u.ID,
		FeedID: f.ID,
	})
	if err != nil {
		return fmt.Errorf("cound not auto-follow feed: %w", err)
	}

	fmt.Println("ID:", f.ID)
	fmt.Println("CreatedAt:", f.CreatedAt)
	fmt.Println("UpdatedAt:", f.UpdatedAt)
	fmt.Println("Name:", f.Name)
	fmt.Println("URL:", f.Url)
	fmt.Println("UserID:", f.UserID)

	return nil
}
