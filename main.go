package main

import (
	"Blogger/internal/config"
	"Blogger/internal/database"
	"database/sql"
	"fmt"
	"os"

	_ "github.com/lib/pq"
)

func main() {

	cfg, err := config.Read()
	if err != nil {
		fmt.Println("error reading config:", err)
		os.Exit(1)
	}

	db, err := sql.Open("postgres", cfg.DbURL)
	if err != nil {
		fmt.Println("error opening db:", err)
		os.Exit(1)
	}
	defer db.Close()

	dbQueries := database.New(db)

	s := &state{
		db:  dbQueries,
		cfg: &cfg}

	cmds := &commands{
		handlers: make(map[string]func(*state, command) error),
	}

	cmds.register("reset", handlerReset)
	cmds.register("register", handlerRegister)
	cmds.register("users", handlerGetUsers)

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
