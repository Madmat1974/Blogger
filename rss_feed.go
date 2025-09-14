package main

import (
	"context"
	"database/sql"
	"encoding/xml"
	"fmt"
	"html"
	"io"
	"net/http"
	"strings"
	"time"

	"github.com/Madmat1974/Gator/internal/database"

	"github.com/google/uuid"
)

type RSSFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description string    `xml:"description"`
		Item        []RSSItem `xml:"item"`
	} `xml:"channel"`
}

type RSSItem struct {
	Title       string `xml:"title"`
	Link        string `xml:"link"`
	Description string `xml:"description"`
	PubDate     string `xml:"pubDate"`
}

func fetchFeed(ctx context.Context, feedURL string) (*RSSFeed, error) {
	req, err := http.NewRequestWithContext(ctx, "GET", feedURL, nil)
	if err != nil {
		return nil, err
	}
	req.Header.Set("User-Agent", "gator")
	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status %d from %s", resp.StatusCode, feedURL)
	}
	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}
	var feed RSSFeed
	if err := xml.Unmarshal(body, &feed); err != nil {
		return nil, err
	}

	feed.Channel.Title = html.UnescapeString(feed.Channel.Title)
	feed.Channel.Description = html.UnescapeString(feed.Channel.Description)
	for i := range feed.Channel.Item {
		feed.Channel.Item[i].Title = html.UnescapeString(feed.Channel.Item[i].Title)
		feed.Channel.Item[i].Description = html.UnescapeString(feed.Channel.Item[i].Description)
	}
	return &feed, nil
}

func postFromRSSItem(item RSSItem, feedID uuid.UUID) database.CreatePostParams {
	return database.CreatePostParams{
		ID:          uuid.New(),
		Title:       toNullString(item.Title),
		Url:         strings.TrimSpace(item.Link),
		Description: toNullString(item.Description),
		PublishedAt: parsePubDate(item.PubDate),
		FeedID:      feedID,
	}
}

// helper functions for postFromRSSItem
func toNullString(s string) sql.NullString {
	s = strings.TrimSpace(s)
	if s == "" {
		return sql.NullString{Valid: false}
	}
	return sql.NullString{String: s, Valid: true}
}

func parsePubDate(s string) sql.NullTime {
	s = strings.TrimSpace(s)
	if s == "" {
		return sql.NullTime{Valid: false}
	}
	layouts := []string{
		time.RFC1123Z, time.RFC1123,
		time.RFC822Z, time.RFC822,
		time.RFC3339, "Mon, 02 Jan 2006 15:04:05 MST",
	}
	for _, layout := range layouts {
		if t, err := time.Parse(layout, s); err == nil {
			return sql.NullTime{Time: t.UTC(), Valid: true}
		}
	}
	return sql.NullTime{Valid: false}
}
