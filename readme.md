## Folder Organization

```
├── bin (compiled application binaries)
├── cmd (application specific code for API: server, reading and writing HTTP requests, and managing authentication)
│ └── api
│   └── main.go
├── internal (basically any code which isn’t application-specific and can potentially be reused)
├── migrations(sql migration files)
├── remote (config files and setup scripts)
├── go.mod
└── Makefile
```

## Endpoints

| path              | method | description                       |
|-------------------|--------|-----------------------------------|
| /v1/health        | GET    | health check endpoint for the api |
| /v1/character     | POST   | adds a new character              |
| /v1/character/:id | GET    | returns the character             |

## Migrate commands

- create new migration sql file

 ```sh
 migrate create -seq -ext=.sql -dir=./migrations command_name 
 ```

 - run migration sql file
 ```sh
migrate -path=./migrations -database=postgres://follow:pass@localhost:9092/follow?sslmode=disable up
 ```

