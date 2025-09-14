package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/Madmat1974/Gator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {

	var limit int32 = 2

	if len(cmd.args) == 1 {
		num, err := strconv.ParseInt(cmd.args[0], 10, 32)
		if err != nil {
			return fmt.Errorf("invalid limit: %w", err)
		}
		limit = int32(num)
	}

	posts, err := s.db.GetPosts(context.Background(), database.GetPostsParams{
		UserID: user.ID,
		Limit:  int32(limit),
	})
	if err != nil {
		return fmt.Errorf("couldn't get posts: %w", err)
	}

	for _, post := range posts {
		fmt.Printf("Title: %v\n", post.Title)
		fmt.Printf("URL: %s\n", post.Url)
		fmt.Printf("Feed: %s\n", post.FeedName)

		if post.PublishedAt.Valid {
			fmt.Printf("Published: %s\n", post.PublishedAt.Time)
		}

		fmt.Println("---")
	}
	return nil
}
