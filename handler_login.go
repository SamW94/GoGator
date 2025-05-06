package main

import (
	"context"
	"fmt"
)

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no arguments supplied with login command - the login handler expects the username in format 'gator login <username>")
	}

	err := s.config.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("error when calling the config.SetUser() function: %w", err)
	}

	_, err = s.db.GetUser(context.Background(), cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("username does not exist in the postgres database: %w", err)
	}

	fmt.Printf("Username set as: %s!\n", cmd.arguments[0])
	return nil
}
