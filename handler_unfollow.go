package main

import (
	"context"
	"fmt"

	"github.com/SamW94/blogo-aggregator/internal/database"
)

func handlerUnfollow(s *state, cmd command, user database.User) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no arguments supplied with unfollow command - the unfollow handler expects a URL in format 'gator unfollow <url>")
	}

	userID := user.ID
	url := cmd.arguments[0]

	feedStruct, err := s.db.GetFeed(context.Background(), url)
	if err != nil {
		return fmt.Errorf("error retrieving feed with URL %s from database.GetFeed(): %w", url, err)
	}

	feedID := feedStruct.ID

	followDeleteParams := database.DeleteFeedFollowParams{
		UserID: userID,
		FeedID: feedID,
	}

	feedFollow, err := s.db.DeleteFeedFollow(context.TODO(), followDeleteParams)
	if err != nil {
		return fmt.Errorf("error calling the database.DeleteFeedFollow() function: %w", err)
	}

	fmt.Println(feedFollow)

	return nil
}
