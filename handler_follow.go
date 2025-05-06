package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerFollow(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no arguments supplied with follow command - the follow handler expects a URL in format 'gator follow <url>")
	}

	url := cmd.arguments[0]
	context := context.Background()

	userStruct, err := s.db.GetUser(context, s.config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error retrieving current user from database.GetUser() function: %w", err)
	}

	userID := userStruct.ID

	feedStruct, err := s.db.GetFeed(context, url)
	if err != nil {
		return fmt.Errorf("error retrieving feed with URL %s from database.GetFeed(): %w", url, err)
	}

	feedID := feedStruct.ID

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feedID,
	}

	_, err = s.db.CreateFeedFollow(context, followParams)
	if err != nil {
		return fmt.Errorf("error calling the database.CreateFeedFollow() function: %w", err)
	}

	fmt.Printf("Feed Name: %s\nCurrent User: %s\n", feedStruct.Name, userStruct.Name)

	return nil
}
