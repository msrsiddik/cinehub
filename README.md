# cinehub

**cinehub** is a Go service built on the **dvdrental** (PostgreSQL sample DB) that provides both **REST** (Fiber + GORM) and **GraphQL** (gqlgen) APIs.  
The goal is to demonstrate how a single domain model can power both API styles, with production-grade project structure, configuration, and developer experience.

---

## Features
- ⚡ **REST API** — Fast routing with Fiber, ORM with GORM  
- 🔎 **GraphQL API** — Schema-first GraphQL powered by gqlgen  
- 🧭 Shared domain/models — REST and GraphQL use the same service/repository layers  
- 📦 Clean architecture layout (service → repository → db)  
- 🛡️ Graceful shutdown, health check endpoint 

---

## Tech Stack
- Go 1.24.5+
- Fiber v2
- GORM + Postgres driver
- gqlgen
- PostgreSQL (dvdrental)

---

## Project Layout
```
cinehub/
├── main-module/
│   ├── cmd/
│   │   └── server/               # main.go (entrypoint)
│   ├── go.mod
│   └── go.sum
├── graphql-module/
│   ├── go.mod
│   ├── go.sum
│   ├── gqlgen.yml
│   ├── tools.go
│   ├── graph/
│   └── server/
├── restapi-module/
│   ├── go.mod
│   ├── go.sum
│   ├── restapi/
│   │   ├── server.go
│   │   └── ...
│   ├── docs/                     # generated Swagger docs
│   └── main.go
├── entities-module/
│   ├── go.mod
│   ├── go.sum
│   ├── database/
│   │   └── db_init.go
│   ├── model/
│   │   ├── actor.gen.go
│   │   ├── address.gen.go
│   │   ├── ...
│   └── query/
│       └── ...
├── tmp/                          # build artifacts, logs
│   ├── build-errors.log
│   └── main
├── .air.toml                     # Air hot reload config
├── .env
├── .gitignore
├── go.work
├── go.work.sum
├── README.md
```

---

## Prerequisites
- Go 1.24.5+
- PostgreSQL 13+ (local or Docker)
- dvdrental sample DB dump (`.tar` format)

### Import dvdrental
```bash
createdb dvdrental
pg_restore -U postgres -d dvdrental -1 dvdrental.tar
```

---

## Configuration
Example `.env`:

```
SERVER_PORT=3000

DB_HOST=192.168.0.101
DB_PORT=5432
DB_USER=postgres
DB_PASSWORD=postgres
DB_NAME=dvdrental
DB_SSL=disable

GRAPHQL_PLAYGROUND=true
```

---

## Run
```bash
go mod tidy
go run ./main-module/cmd/server
# or with hot reload
air
```

Server runs at: `http://localhost:3000`

- GraphQL Playground path: `/graphql`
- GraphQL endpoint: `/query`
- Health check: `/healthz`

---

---

## Generate Swagger Documentation

To generate or update the Swagger documentation for the REST API module, run the following command from the project root:

```bash
swag init -g restapi-module/restapi/server.go -o restapi-module/docs
```

**Instructions:**
- Make sure you are in the root directory of the project (where `go.work` is located).
- The `-g` flag specifies the main Go file containing your Swagger annotations.
- The `-o` flag specifies the output directory for the generated Swagger docs.
- After running this command, the Swagger UI will be available at:

---

## Credits
- dvdrental sample DB (PostgreSQL)  
- Fiber, GORM, gqlgen communities  