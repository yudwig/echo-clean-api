package cmd

import (
	"errors"
	"github.com/abiosoft/ishell"
	"github.com/yudwig/echo-sample/app/driver/client"
)

func RunShell() {
	sh := ishell.New()
	sh.Println("Client start.")
	sh.AddCmd(&ishell.Cmd{
		Name:      "get",
		Aliases:   nil,
		Help:      "get users",
		LongHelp:  "",
		Completer: nil,
		Func: func(c *ishell.Context) {
			client.GetAllUsers()
		},
	})
	sh.AddCmd(&ishell.Cmd{
		Name:      "create",
		Aliases:   nil,
		Help:      "create user",
		LongHelp:  "",
		Completer: nil,
		Func: func(c *ishell.Context) {
			var err error
			if len(c.Args) > 1 {
				err = errors.New("too many args")
			}
			if len(c.Args) == 0 {
				err = errors.New("empty name")
			}
			if err != nil {
				sh.Println("error: ", err.Error())
				sh.Println("e.g) > create john")
				return
			}
			client.CreateUser(c.Args[0])
		},
	})
	sh.AddCmd(&ishell.Cmd{
		Name:      "update",
		Aliases:   nil,
		Help:      "update user name",
		LongHelp:  "",
		Completer: nil,
		Func: func(c *ishell.Context) {
			var err error
			if len(c.Args) > 2 {
				err = errors.New("input id and name")
			}
			if err != nil {
				sh.Println("error: ", err.Error())
				sh.Println("e.g) > update 2 jasmine")
				return
			}

			client.UpdateUser(c.Args[0], c.Args[1])
		},
	})
	sh.AddCmd(&ishell.Cmd{
		Name:      "delete",
		Aliases:   nil,
		Help:      "delete user",
		LongHelp:  "",
		Completer: nil,
		Func: func(c *ishell.Context) {
			var err error
			if len(c.Args) != 1 {
				err = errors.New("input user id")
			}
			if err != nil {
				sh.Println("error: ", err.Error())
				sh.Println("e.g) > delete 1")
				return
			}
			client.DeleteUser(c.Args[0])
		},
	})
	sh.Run()
}
