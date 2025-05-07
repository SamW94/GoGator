package main

import (
	"context"
	"fmt"
	"strings"
	"time"

	"github.com/SamW94/GoGator/internal/database"
	"github.com/SamW94/GoGator/internal/rss"
	"github.com/google/uuid"
)

func createPostInDB(s *state, rssItem rss.RSSItem, feedURL string) (database.CreatePostRow, error) {
	context := context.Background()
	publishedatParsed, err := time.Parse("Mon, 02 Jan 2006 15:04:05 MST", rssItem.PubDate)

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
