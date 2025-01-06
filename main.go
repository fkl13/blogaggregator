package main

import (
	"fmt"

	"github.com/fkl13/boot.dev/blogaggregator/internal/config"
)

func main() {
	cfg, err := config.Read()
	if err != nil {
		fmt.Printf("could not read config: %v\n", err)
		return
	}

	err = cfg.SetUser("fabian")
	if err != nil {
		fmt.Printf("could not write config file: %v\n", err)
		return
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Printf("could not read config: %v\n", err)
		return
	}
	fmt.Println(cfg)
}
