package main

import (
	"context"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/database"
	"github.com/SamW94/blogo-aggregator/internal/rss"
	"github.com/google/uuid"
)

func createPostInDB(s *state, rssItem rss.RSSItem, feedURL string) (database.CreatePostRow, error) {
	context := context.Background()

	timeFormats := []string{
		"Mon Jan 02 15:04:05 2006",
		"Mon Jan 02 15:04:05 MST 2006",
		"Mon Jan 02 15:04:05 -0700 2006",
		"02 Jan 06 15:04 MST",
		"02 Jan 06 15:04 -0700",
		"Monday, 02-Jan-06 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 MST",
		"Mon, 02 Jan 2006 15:04:05 -0700",
		"2006-01-02T15:04:05Z07:00",
		"2006-01-02T15:04:05.999999999Z07:00",
		"3:04PM",
	}

	var publishedatParsed time.Time
	var err error
	for _, format := range timeFormats {
		publishedatParsed, err = time.Parse(format, rssItem.PubDate)
		if err == nil {
			log.Printf("Parsed date '%s' using format '%s'", rssItem.PubDate, format)
			break
		}
	}

	if err != nil {
		return database.CreatePostRow{}, fmt.Errorf("unable to parse date format: %s: %w", rssItem.PubDate, err)
	}

	parentFeed, err := s.db.GetFeed(context, feedURL)
	if err != nil {
		return database.CreatePostRow{}, fmt.Errorf("error calling database.GetFeed() function: %w", err)
	}

	createPostParams := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       rssItem.Title,
		Url:         rssItem.Link,
		Description: rssItem.Description,
		PublishedAt: publishedatParsed,
		FeedID:      parentFeed.ID,
	}

	post, err := s.db.CreatePost(context, createPostParams)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") || strings.Contains(err.Error(), "unique constraint") {
			return database.CreatePostRow{}, nil
		}
		return database.CreatePostRow{}, fmt.Errorf("error calling the database.CreatePost() function: %w", err)
	}
	return post, nil
}
