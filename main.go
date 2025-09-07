package main

import (
	"Blogger/internal/config"
	"fmt"
)

func main() {
	user := "Martin"

	cfg, err := config.Read()
	if err != nil {
		fmt.Println(err)
	}

	err = cfg.SetUser(user)
	if err != nil {
		fmt.Println(err)
	}

	cfg, err = config.Read()
	if err != nil {
		fmt.Println(err)
	}

	fmt.Printf("%+v\n", cfg)

}
