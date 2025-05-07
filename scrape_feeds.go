package main

import (
	"context"
	"database/sql"
	"fmt"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/database"
	"github.com/SamW94/blogo-aggregator/internal/rss"
)

func scrapeFeeds(s *state, rssClient *rss.Client) error {
	context := context.Background()

	feed, err := s.db.GetNextFeedToFetch(context)
	if err != nil {
		return fmt.Errorf("error when calling database.GetNextFeedToFetch() function: %w", err)
	}

	feedID := feed.ID
	nullTime := sql.NullTime{
		Time:  time.Now(),
		Valid: true,
	}

	markFetchedParams := database.MarkFeedFetchedParams{
		LastFetchedAt: nullTime,
		ID:            feedID,
	}

	feed, err = s.db.MarkFeedFetched(context, markFetchedParams)
	if err != nil {
		return fmt.Errorf("error when calling database.MarkFeedFetched() function: %w", err)
	}

	rssFeed, err := rssClient.FetchFeed(context, feed.Url)
	if err != nil {
		return fmt.Errorf("error calling rss.FetchFeed() function with URL %s: %w", feed.Url, err)
	}

	for i := 0; len(rssFeed.Channel.Item) > i; i++ {
		fmt.Println(rssFeed.Channel.Item[i].Title)
	}

	return nil
}
