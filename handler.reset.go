package main

import (
	"context"
	"fmt"
	"os"
)

func handlerReset(s *state, cmd command) error {
	err := s.db.DeleteUsers(context.Background())
	if err != nil {
		fmt.Println("unable to reset")
		os.Exit(1)
	}
	if err == nil {
		fmt.Println("reset was successful")
		os.Exit(0)
	}
	return err
}
