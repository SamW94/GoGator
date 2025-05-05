package main

import (
	"fmt"
	"os"

	"github.com/SamW94/blogo-aggregator/internal/config"
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

	err = commands.run(&appState, inputCommand)
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
