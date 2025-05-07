package main

import (
	"fmt"
	"time"

	"github.com/SamW94/GoGator/internal/rss"
)

func handlerAgg(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no arguments supplied with agg command - the agg handler expects the username in format 'agg <duration>, e.g. agg 1h")
	}

	time_between_reqs, err := time.ParseDuration(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error parsing duration from provided string %s: %w", cmd.arguments[0], err)
	}

	fmt.Printf("Collecting feeds every %v\n", time_between_reqs)
	rssClient := rss.NewClient(5 * time.Second)
	ticker := time.NewTicker(time_between_reqs)
	for {
		scrapeFeeds(s, &rssClient)
		<-ticker.C
	}
}
