package main

import (
	"context"
	"fmt"
)

func handlerFeeds(s *state, cmd command) error {
	context := context.Background()
	feeds, err := s.db.GetFeeds(context)
	if err != nil {
		return fmt.Errorf("error calling the database.GetFields() function: %w", err)
	}

	for _, feed := range feeds {
		fmt.Printf("Feed name: %s\nFeed URL: %s\nCreated by user: %s\n", feed.Name, feed.Url, feed.Name_2)
	}

	return nil
}
