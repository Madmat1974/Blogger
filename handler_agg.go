package main

import (
	"context"
	"fmt"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()
	f, err := fetchFeed(ctx, "https://wagslane.dev/index.xml")
	if err != nil {
		return err
	}
	fmt.Printf("%+v\n", f)
	return nil
}
