package main

import (
	"context"
	"fmt"
	"gator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	ctx := context.Background()

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("please login first")
	}
	ffs, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("could not list follows: %w", err)
	}

	for _, ff := range ffs {
		fmt.Println(ff.FeedName)
	}
	return nil
}
