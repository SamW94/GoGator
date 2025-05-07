package main

import (
	"github.com/SamW94/GoGator/internal/config"
	"github.com/SamW94/GoGator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
