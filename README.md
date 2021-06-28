# echo-clean-api
A sample REST API based on clean architecture using GO framework echo.

## Features
echo-clean-api is demonstration of CRUD API for a simple user management system.
To test API also includes simple REPL CLI client tool.


## How to run

#### DB migration

By default, app uses local mysql DB which dns is "root@tcp(127.0.0.1:3306)/echo". So first, create database 'echo' before migration.
After that, run migration command.

```sh
go run main.go migrate
```

#### Run server

Run echo server on "localhost:8000" by this command.

```sh
go run main.go server
```

#### Run client shell

Run API client tool by this command.

```sh
go run main.go client
```

When started shell, command list are displayed as follows.

```sh
% go run main.go client
Client shell start.

Commands:
  clear       clear the screen
  create      create user. input: > create {name}
  delete      delete user. input: > delete {id}
  exit        exit the program
  get         get user list. input: > get
  help        display help
  update      update user name. input: > update {id} {name}

>>>
```

## Directory structure
The application has the following directory structure.

```
app
├── adapter
│   ├── controller
│   ├── interactor
│   ├── presenter
│   └── repository
├── core
│   ├── domain
│   │   └── entity
│   │       └── user
│   ├── presentation
│   │   └── response
│   └── usecase
└── driver
    ├── client
    ├── db
    │   ├── mockdb
    │   ├── rdb
    │   └── seed
    ├── echo
    │   ├── routes
    │   └── server
    └── log
```

Based on the clean architecture, it has the following dependency directions.

```
 ┌────────┐      ┌─────────┐      ┌────────┐  
 │ driver │ ───> │ adapter │ ───> │  core  │ 
 └────────┘      └─────────┘      └────────┘ 
```

