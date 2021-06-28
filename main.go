package main

import (
	"fmt"
	"github.com/yudwig/echo-clean-api/app/driver/db/rdb"
	"github.com/yudwig/echo-clean-api/app/driver/echo/server"
	"github.com/yudwig/echo-clean-api/cmd"
	"os"
)

func help() {
	fmt.Println("Input command.")
	fmt.Println("- server:  Run server")
	fmt.Println("- client:  Run client repl shell")
	fmt.Println("- migrate: Migrate database")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	mode := os.Args[1]
	switch mode {
	case "server":
		server.Run()
	case "migrate":
		db, _ := rdb.NewDb()
		rdb.Migrate(db)
	case "client":
		cmd.RunShell()
	default:
		help()
	}
}
