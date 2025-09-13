package main

import (
	"context"
	"database/sql"
	"errors"
	"fmt"
	"time"

	"github.com/lib/pq"
)

func handlerAgg(s *state, cmd command) error {
	ctx := context.Background()

	if len(cmd.args) != 1 {
		return fmt.Errorf("usage: %s <time_between_reqs> (e.g. 10s, 1m, 1h)", cmd.name)
	}

	d, err := time.ParseDuration(cmd.args[0])
	if err != nil {
		return fmt.Errorf("invalid duration: %w", err)
	}

	fmt.Printf("Collecting feeds every %s\n", d)
	ticker := time.NewTicker(d)
	for ; ; <-ticker.C {
		if err := scrapeFeeds(ctx, s); err != nil {
			fmt.Printf("scrape error: %v\n", err)
		}
	}
}

func scrapeFeeds(ctx context.Context, s *state) error {
	feed, err := s.db.GetNextFeedToFetch(ctx)
	if err == sql.ErrNoRows {
		fmt.Println("no feeds to fetch")
		return nil
	}
	if err != nil {
		return err
	}
	fmt.Printf("fetching: name=%q url=%q\n", feed.Name, feed.Url)

	if err := s.db.MarkFeedFetched(ctx, feed.ID); err != nil {
		return err
	}

	f, err := fetchFeed(ctx, feed.Url)
	if err != nil {
		return err
	}

	for _, it := range f.Channel.Item {
		params := postFromRSSItem(it, feed.ID)
		if _, err := s.db.CreatePost(ctx, params); err != nil {
			var pqErr *pq.Error
			if errors.As(err, &pqErr) && pqErr.Code == "23505" { // unique_violation
				continue
			}
			fmt.Printf("create post error feed=%s url=%s err=%v\n", feed.ID, params.Url, err)
			continue
		}
	}
	return nil
}
