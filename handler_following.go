package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	ctx := context.Background()

	if s.cfg.CurrentUserName == "" {
		return fmt.Errorf("please login first")
	}
	user, err := s.db.GetUser(ctx, s.cfg.CurrentUserName)
	if err != nil {
		return fmt.Errorf("current user not found: %w", err)
	}

	ffs, err := s.db.GetFeedFollowsForUser(ctx, user.ID)
	if err != nil {
		return fmt.Errorf("could not list follows: /w", err)
	}

	for _, ff := range ffs {
		fmt.Println(ff.FeedName)
	}
	return nil
}
