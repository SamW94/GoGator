package main

import (
	"context"
	"fmt"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.ResetUsers(context.Background())
	if err != nil {
		return fmt.Errorf("error calling the database.ResetUsers() function: %w", err)
	}

	fmt.Println("users table has been successfully reset")
	return nil
}
