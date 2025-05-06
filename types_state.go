package main

import (
	"github.com/SamW94/blogo-aggregator/internal/config"
	"github.com/SamW94/blogo-aggregator/internal/database"
)

type state struct {
	db     *database.Queries
	config *config.Config
}
