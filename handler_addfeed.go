package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("not enough arguments supplied - the addfeed handler expects 2 arguments in format 'gator addfeed <name> <url>")
	}

	context := context.Background()
	userStruct, err := s.db.GetUser(context, s.config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error retrieving current user from database.GetUser() function: %w", err)
	}

	userID := userStruct.ID
	name := cmd.arguments[0]
	url := cmd.arguments[1]

	feedParams := database.CreateFeedParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
		Url:       url,
		UserID:    userID,
	}

	feed, err := s.db.CreateFeed(context, feedParams)
	if err != nil {
		return fmt.Errorf("error calling database.CreateFeed() function: %w", err)
	}

	fmt.Println(feed)
	return nil
}
