package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {

	ctx := context.Background()

	rows, err := s.db.GetFeeds(ctx)
	if err != nil {
		return err
	}
	if len(rows) == 0 {
		fmt.Println("no feeds found")
		return nil
	}
	for _, r := range rows {
		fmt.Printf("%s - %s (by %s)\n", r.FeedsName, r.FeedsUrl, r.UserName)
	}
	return nil
}
