package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/database"
	"github.com/SamW94/blogo-aggregator/internal/rss"
	"github.com/google/uuid"
)

func createPostInDB(s *state, rssItem rss.RSSItem, feedURL string) (database.CreatePostRow, error) {
	context := context.Background()

	title := rssItem.Title
	url := rssItem.Link
	description := rssItem.Description
	publishedat := rssItem.PubDate
	publishedatParsed, err := time.Parse("2025-05-07 16:13:04.981876", publishedat)

	parentFeed, err := s.db.GetFeed(context, feedURL)
	if err != nil {
		return database.CreatePostRow{}, fmt.Errorf("error calling database.GetFeed() function: %w", err)
	}

	feedID := parentFeed.ID

	createPostParams := database.CreatePostParams{
		ID:          uuid.New(),
		CreatedAt:   time.Now(),
		UpdatedAt:   time.Now(),
		Title:       title,
		Url:         url,
		Description: description,
		PublishedAt: publishedatParsed,
		FeedID:      feedID,
	}

	s.db.CreatePost(context, createPostParams)

	return database.CreatePostRow{}, nil
}
