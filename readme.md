## Folder Organization

├── bin (compiled application binaries)
├── cmd (application specific code for API: server, reading and writing HTTP requests, and managing authentication)
│ └── api
│   └── main.go
├── internal (basically any code which isn’t application-specific and can potentially be reused)
├── migrations(sql migration files)
├── remote (config files and setup scripts)
├── go.mod
└── Makefile
