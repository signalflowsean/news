# Backend (Go)

HTTP server that serves the story page with data from Neo4j.

## Prerequisites

- Go 1.21+ (`brew install go`)
- Neo4j instance running

## Setup

1. Copy environment file and configure:
   ```bash
   cp ../.env.example ../.env
   # Edit ../.env with your Neo4j credentials
   ```

2. Install dependencies:
   ```bash
   go mod tidy
   ```

## Run

```bash
source ../.env && go run main.go
```

Server starts at `http://localhost:8080` (or `$PORT`).

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Server port |
| `STATIC_DIR` | `../frontend/dist` | Path to built frontend assets |
<!-- TODO -->
| `NEO4J_URI` | `neo4j://localhost:7687` | Neo4j connection URI |
| `NEO4J_USER` | `neo4j` | Neo4j username |
| `NEO4J_PASSWORD` | (required) | Neo4j password |
