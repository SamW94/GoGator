package main

import (
	"context"
	"fmt"

	"github.com/SamW94/GoGator/internal/database"
)

func handlerFollowing(s *state, cmd command, user database.User) error {
	userID := user.ID

	feedFollows, err := s.db.GetFeedFollowsForUser(context.Background(), userID)
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
