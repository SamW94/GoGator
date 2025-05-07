package main

import (
	"context"
	"fmt"

	"github.com/SamW94/blogo-aggregator/internal/database"
)

func middlewareLoggedIn(handler func(s *state, cmd command, user database.User) error) func(*state, command) error {
	return func(s *state, cmd command) error {
		context := context.Background()
		userStruct, err := s.db.GetUser(context, s.config.CurrentUsername)
		if err != nil {
			return fmt.Errorf("error retrieving current user from database.GetUser() function: %w", err)
		}
		err = handler(s, cmd, userStruct)
		if err != nil {
			return err
		}
		return nil
	}
}
