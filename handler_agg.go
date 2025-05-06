package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	rssClient := rss.NewClient(5 * time.Second)
	rssFeed, err := rssClient.FetchFeed(context.Background(), "https://www.wagslane.dev/index.xml")
	if err != nil {
		return fmt.Errorf("error calling the rss.FetchFeed() function: %w", err)
	}

	fmt.Println(*rssFeed)
	return nil
}
