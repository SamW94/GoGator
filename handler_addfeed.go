package main

import (
	"context"
	"fmt"
	"time"

	"github.com/SamW94/GoGator/internal/database"
	"github.com/google/uuid"
)

func handlerAddfeed(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) < 2 {
		return fmt.Errorf("not enough arguments supplied - the addfeed handler expects 2 arguments in format 'gator addfeed <name> <url>")
	}

	userID := user.ID
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

	feed, err := s.db.CreateFeed(context.Background(), feedParams)
	if err != nil {
		return fmt.Errorf("error calling database.CreateFeed() function: %w", err)
	}

	followParams := database.CreateFeedFollowParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		UserID:    userID,
		FeedID:    feed.ID,
	}

	_, err = s.db.CreateFeedFollow(context.Background(), followParams)
	if err != nil {
		return fmt.Errorf("error calling the database.CreateFeedFollow() function: %w", err)
	}

	fmt.Println(feed)
	return nil
}
