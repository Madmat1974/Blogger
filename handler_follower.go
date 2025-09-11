package main

import (
	"Blogger/internal/database"
	"context"
	"fmt"

	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	ctx := context.Background()
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <url>", cmd.name)
	}
	url := cmd.args[0]
	// check if there is a current user logged in
	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("please login first")
	}
	//check database as well
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	f, err := s.db.GetFeedByURL(ctx, url)
	if err != nil {
		return fmt.Errorf("feed not found for usr: %s", url)
	}

	ff, err := s.db.CreateFeedFollow(ctx, database.CreateFeedFollowParams{
		ID:     uuid.New(),
		UserID: user.ID,
		FeedID: f.ID,
	})
	if err != nil {
		return fmt.Errorf("could not create follow: %w", err)
	}

	fmt.Printf("Following Feed: %s    Feed Follower: %s\n", ff.FeedName, ff.UserName)
	return nil
}
