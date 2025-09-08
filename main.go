package main

import (
	"Blogger/internal/config"
	"fmt"
	"os"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	s := &state{cfg: &cfg}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}
	cmds.register("login", handlerLogin)
	if len(os.Args) < 2 {
		fmt.Println("wrong input")
		os.Exit(1)
	}
	cmd := command{name: os.Args[1], args: os.Args[2:]}

	if err := cmds.run(s, cmd); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
