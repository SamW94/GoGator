package main

import (
	"context"
	"fmt"
	"strconv"

	"github.com/SamW94/blogo-aggregator/internal/database"
)

func handlerBrowse(s *state, cmd command, user database.User) error {
	var limit int32 = 2
	if len(cmd.arguments) > 0 {
		parsedLimit, err := strconv.Atoi(cmd.arguments[0])
		if err != nil {
			return fmt.Errorf("error converting command argument %s to integer for limit: %w", cmd.arguments[0], err)
		}
		limit = int32(parsedLimit)
	} else {
		fmt.Println("no limit value provided as argument, defaulting limit to 2.")
	}

	posts, err := s.db.GetPostsForUser(context.Background(), database.GetPostsForUserParams{
		UserID: user.ID,
		Limit:  limit,
	})
	if err != nil {
		return fmt.Errorf("error getting posts for user: %w", err)
	}

	fmt.Println("Recent posts from your feeds:")
	for i, post := range posts {
		fmt.Printf("%d. %s\n", i+1, post.PostTitle)
		fmt.Printf("	%s\n", post.PostUrl)
		fmt.Printf("	Published: %s\n\n", post.PostPublishedAt.Format("Jan 2, 2006"))
	}
	return nil
}
