package main

import (
	"context"
	"fmt"
)

func handlerUsers(s *state, cmd command) error {
	users, err := s.db.GetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error calling the database.GetUsers() function: %w", err)
	}

	if len(users) == 0 {
		fmt.Println("There are currently no users configured in the database. To register one run 'gator register <username>'")
	}

	for _, user := range users {
		// if the currently logged-in user print in format * <user> (current)
		// else print in format * <user>
		if user.Name == s.config.CurrentUsername {
			fmt.Printf("* %s (current)\n", user.Name)
		} else {
			fmt.Printf("* %s\n", user.Name)
		}
	}

	return nil
}
