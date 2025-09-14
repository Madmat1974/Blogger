package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"

	"github.com/Madmat1974/Gator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	ctx := context.Background()
	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: unfollow <url>")
	}
	url := cmd.args[0]
	feed, err := s.db.GetFeedByURL(ctx, url) //get feed id via url
	if err != nil {
		return fmt.Errorf("feed not found for user: %s", url)
	}
	err = s.db.Unfollow(ctx, database.UnfollowParams{
		UserID: user.ID,
		FeedID: feed.ID,
	})
	if err != nil {
		if errors.Is(err, sql.ErrNoRows) {
			return fmt.Errorf("feed not found: %s", url)
		}
		return err
	}
	fmt.Printf("unfollowed: %s\n", url)
	return nil
}
