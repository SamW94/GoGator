package main

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/SamW94/blogo-aggregator/internal/config"
	"github.com/SamW94/blogo-aggregator/internal/database"
	_ "github.com/lib/pq"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println("error calling the config.Read() function:\n %w", err)
	}
	appState := state{
		config: &configFile,
	}

	commands := commands{
		handlers: make(map[string]func(*state, command) error),
	}

	commands.register("login", handlerLogin)
	commands.register("register", handlerRegister)
	commands.register("reset", handlerReset)
	commands.register("users", handlerUsers)
	commands.register("agg", handlerAgg)
	commands.register("addfeed", middlewareLoggedIn(handlerAddfeed))
	commands.register("feeds", handlerFeeds)
	commands.register("follow", middlewareLoggedIn(handlerFollow))
	commands.register("following", middlewareLoggedIn(handlerFollowing))
	commands.register("unfollow", middlewareLoggedIn(handlerUnfollow))

	inputArguments := os.Args
	if len(inputArguments) < 2 {
		fmt.Println("no arguments provided - please provide at least one")
		fmt.Println(inputArguments[0])
		os.Exit(1)
	}

	inputCommand := command{
		name:      inputArguments[1],
		arguments: inputArguments[2:],
	}

	db, err := sql.Open("postgres", appState.config.DBUrl)
	if err != nil {
		fmt.Printf("error opening connection to postgres DB: %v", err)
	}

	dbQueries := database.New(db)
	appState.db = dbQueries

	err = commands.run(&appState, inputCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
