# go-starter

A production-ready HTTP server starter project built with modern Go.

## Stack

| Layer | Technology |
|-------|------------|
| Language | Go 1.26 |
| Router | [chi](https://github.com/go-chi/chi) v5 |
| Config | [caarlos0/env](https://github.com/caarlos0/env) |
| Logging | `log/slog` (stdlib) |
| Database | PostgreSQL 17 + [pgx](https://github.com/jackc/pgx) v5 |
| Testing | `testing` + [testify](https://github.com/stretchr/testify) |
| Container | Docker (multi-stage) + distroless |
| CI | GitHub Actions |
| Lint | golangci-lint |

## Project Structure

```
.
├── cmd/server/          # Application entry point
├── internal/
│   ├── config/          # Environment variables & logger
│   ├── database/        # PostgreSQL connection (pgx)
│   ├── handler/         # HTTP handlers
│   └── server/          # Router & HTTP server
├── scripts/init.sql       # PostgreSQL seed script
├── .github/workflows/   # CI pipeline
├── .vscode/             # VS Code settings
├── Dockerfile
├── docker-compose.yml
└── Makefile
```

## Quick Start

### Requirements

- Go 1.26+
- Docker & Docker Compose (optional)

### Run Locally

```bash
# Download dependencies
go mod download

# Start the server
make run
# or
go run ./cmd/server
```

The server starts at `http://localhost:8080` by default.

### Run with Docker

PostgreSQL and the API start together:

```bash
# Build & run (postgres + api)
make docker-up

# Build only
make docker-build

# Stop
make docker-down

# PostgreSQL shell
make docker-psql
```

Docker Compose starts:

- **postgres** — PostgreSQL 17 (`localhost:5432`)
- **api** — Go HTTP server (`localhost:8080`), starts after postgres is ready

### Local development + Docker PostgreSQL

Run only the database in Docker and start the API locally:

```bash
docker compose up postgres -d
cp .env.example .env
make run
```

## API Endpoints

| Method | Path | Description |
|--------|------|-------------|
| `GET` | `/health` | Health check (includes DB status) |
| `GET` | `/api/v1/hello?name=Go` | Greeting endpoint |
| `GET` | `/api/v1/status` | Application info |
| `GET` | `/api/v1/db` | PostgreSQL connection info |

### Example Requests

```bash
curl http://localhost:8080/health
# {"status":"ok","database":"ok"}

curl http://localhost:8080/api/v1/hello?name=Go
# {"message":"Hello, Go!"}

curl http://localhost:8080/api/v1/status
# {"app":"go-starter","version":"dev","env":"development"}

curl http://localhost:8080/api/v1/db
# {"connected":true,"version":"PostgreSQL 17...","greetings":["World","Go"]}
```

## Configuration

The app is configured via environment variables. Copy `.env.example` to get started:

```bash
cp .env.example .env
```

| Variable | Default | Description |
|----------|---------|-------------|
| `APP_NAME` | `go-starter` | Application name |
| `APP_VERSION` | `dev` | Version string |
| `ENV` | `development` | Environment (`development` / `production`) |
| `PORT` | `8080` | Port to listen on |
| `LOG_LEVEL` | `info` | Log level (`debug`, `info`, `warn`, `error`) |
| `DATABASE_URL` | _(empty)_ | PostgreSQL connection string |
| `POSTGRES_USER` | `postgres` | Docker Compose DB user |
| `POSTGRES_PASSWORD` | `postgres` | Docker Compose DB password |
| `POSTGRES_DB` | `go_starter` | Docker Compose DB name |
| `POSTGRES_PORT` | `5432` | PostgreSQL host port |

> **Note:** Logs are written in text format in `development` and JSON format in `production`.

## Tests

```bash
# Run all tests
make test

# Coverage report
make test-cover
# generates coverage.html
```

## Lint

```bash
# requires golangci-lint: https://golangci-lint.run/welcome/install/
make lint
```

## VS Code

The project includes ready-to-use settings in `.vscode/`:

- **Launch Server** — debug with environment variables
- **Debug Tests** — debug handler package tests
- Go format-on-save, organize imports

Recommended extensions are listed in `.vscode/extensions.json`.

## CI/CD

The GitHub Actions pipeline runs on every push and PR:

1. **Test** — unit tests with race detector + coverage
2. **Lint** — golangci-lint
3. **Build** — compile binary
4. **Docker** — verify image build

Dependabot opens weekly PRs for Go module and GitHub Actions updates.

## Makefile Commands

| Command | Description |
|---------|-------------|
| `make run` | Run the server locally |
| `make build` | Build binary (`bin/server`) |
| `make test` | Run tests |
| `make test-cover` | Generate coverage report |
| `make lint` | Run golangci-lint |
| `make tidy` | Run go mod tidy |
| `make docker-up` | Start with Docker Compose (postgres + api) |
| `make docker-down` | Stop Docker Compose |
| `make docker-psql` | Open PostgreSQL shell |

## License

Apache License 2.0 — see [LICENSE](LICENSE) for details.
