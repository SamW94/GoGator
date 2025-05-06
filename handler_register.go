package main

import (
	"context"
	"fmt"
	"log"
	"time"

	"github.com/SamW94/blogo-aggregator/internal/database"
	"github.com/google/uuid"
)

func handlerRegister(s *state, cmd command) error {
	if len(cmd.arguments) == 0 {
		return fmt.Errorf("no arguments supplied with register command - the register handler expects a username in format 'gator register <username>")
	}

	name := cmd.arguments[0]

	userParams := database.CreateUserParams{
		ID:        uuid.New(),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
		Name:      name,
	}

	_, err := s.db.CreateUser(context.Background(), userParams)
	if err != nil {
		return fmt.Errorf("error calling the database.CreateUser() function: %w", err)
	}

	err = handlerLogin(s, cmd)
	if err != nil {
		return fmt.Errorf("error calling the handlerLogin() function from handlerRegister: %w", err)
	}

	fmt.Printf("User %s was created in postgres DB.", name)
	log.Printf("user created in Postgres DB with parameter: %v", userParams)

	return nil
}
