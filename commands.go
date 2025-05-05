package main

import (
	"fmt"
)

func (c *commands) run(s *state, cmd command) error {
	_, ok := c.handlers[cmd.name]
	if !ok {
		return fmt.Errorf("command does not exist")
	}

	err := c.handlers[cmd.name](s, cmd)
	if err != nil {
		return fmt.Errorf("error running the %s command:\n %w", cmd.name, err)
	}
	return nil
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.handlers[name] = f
}
