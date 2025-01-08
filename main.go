package main

import (
	"fmt"
	"os"

	"github.com/fkl13/boot.dev/blogaggregator/internal/config"
)

type state struct {
	cfg *config.Config
}

type command struct {
	name      string
	arguments []string
}

type commands struct {
	cmdMap map[string]func(*state, command) error
}

func (c *commands) register(name string, f func(*state, command) error) {
	c.cmdMap[name] = f
}

func (c *commands) run(s *state, cmd command) error {
	f, ok := c.cmdMap[cmd.name]
	if !ok {
		return fmt.Errorf("command not found")
	}
	return f(s, cmd)
}

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("could not read config: %v\n", err)
		return
	}

	s := &state{cfg: &cfg}
	cmds := commands{cmdMap: map[string]func(*state, command) error{}}
	cmds.cmdMap["login"] = handlerLogin

	if len(os.Args) < 3 {
		fmt.Println("Too few arguments")
		os.Exit(1)
	}

	cmd := command{
		name:      os.Args[1],
		arguments: os.Args[2:],
	}

	if cmd.name == "login" {
		err := cmds.cmdMap["login"](s, cmd)
		if err != nil {
			fmt.Printf("Command %s failed\n", cmd.name)
			return
		}
	}

	fmt.Println(s.cfg)
}

func handlerLogin(s *state, cmd command) error {
	if len(cmd.arguments) != 1 {
		return fmt.Errorf("usage: %s <name>", cmd.name)
	}

	err := s.cfg.SetUser(cmd.arguments[0])
	if err != nil {
		return fmt.Errorf("couldn't set current user: %w", err)
	}

	return nil
}
