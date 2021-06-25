package main

import (
	"fmt"
	"github.com/yudwig/echo-sample/app/driver/db/rdb"
	"github.com/yudwig/echo-sample/app/driver/echo/server"
	"os"
)

func help() {
	fmt.Println("Input command.")
	fmt.Println("- serve: Run server")
	fmt.Println("- migrate: Run migration")
}

func main() {
	if len(os.Args) < 2 {
		help()
		os.Exit(1)
	}
	cmd := os.Args[1]
	switch cmd {
	case "serve":
		server.Run()
	case "migrate":
		db, _ := rdb.NewDb()
		rdb.Migrate(db)
	default:
		help()
	}
}
