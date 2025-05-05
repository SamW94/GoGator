package main

import (
	"fmt"

	"github.com/SamW94/blogo-aggregator/internal/config"
)

func main() {
	configFile, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	configFile.SetUser("Sam")

	configFile, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Println(configFile)
}
