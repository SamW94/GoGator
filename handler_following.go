package main

import (
	"context"
	"fmt"
)

func handlerFollowing(s *state, cmd command) error {
	context := context.Background()

	userStruct, err := s.db.GetUser(context, s.config.CurrentUsername)
	if err != nil {
		return fmt.Errorf("error retrieving current user from database.GetUser() function: %w", err)
	}

	userID := userStruct.ID

	feedFollows, err := s.db.GetFeedFollowsForUser(context, userID)
	if err != nil {
		return fmt.Errorf("error calling the database.GetFeedFollowsForUser() function: %w", err)
	}

	if len(feedFollows) == 0 {
		fmt.Printf("User %s is not following any feeds\n", s.config.CurrentUsername)
	} else {
		fmt.Printf("User %s is following these feeds:\n", s.config.CurrentUsername)
		for _, feedFollow := range feedFollows {
			fmt.Println(feedFollow.FeedName)
		}
	}

	return nil
}
