# News Graph

Graph visualization of news stories and events, powered by GDELT data and Neo4j.

## Architecture

```
┌─────────────────────────────────────────────────────────────────┐
│                         Browser                                  │
│  ┌─────────────────────────────────────────────────────────┐    │
│  │                    Canvas (#engine)                         │    │
│  │         Rendered by Rust/Wasm (future: WebGPU)          │    │
│  └─────────────────────────────────────────────────────────┘    │
│                              ▲                                   │
│                    __INITIAL_DATA__                              │
└──────────────────────────────┬──────────────────────────────────┘
                               │
                     ┌─────────┴─────────┐
                     │   Go Backend      │
                     │   (net/http)      │
                     └─────────┬─────────┘
                               │
                     ┌─────────┴─────────┐
                     │      Neo4j        │
                     │  (graph database) │
                     └─────────┬─────────┘
                               │
                     ┌─────────┴─────────┐
                     │  Python Ingestion │
                     │  (GDELT/BigQuery) │
                     └───────────────────┘
```

## Directory Structure

```
news/
├── backend/          # Go HTTP server
│   ├── main.go
│   └── internal/
│       ├── templates/
│       └── neo4j/
├── frontend/         # Vite + TypeScript
│   ├── src/
│   └── engine/       # Rust/Wasm crate
└── ingestion/        # Python ingestion jobs
    └── src/
```

## Quick Start

### Prerequisites

- Go 1.21+
- Node.js 20.19+
- Rust + wasm-pack
- Python 3.11+
- Neo4j (local or cloud)

### Setup

1. Configure environment:
   ```bash
   cp .env.example .env
   # Edit .env with your Neo4j credentials
   ```

2. Build frontend:
   ```bash
   cd frontend
   npm install
   npm run build
   ```

3. Run backend:
   ```bash
   cd backend
   source ../.env
   go mod tidy
   go run main.go
   ```

4. Open `http://localhost:8080`

## Development

See individual READMEs in each directory:
- [backend/README.md](backend/README.md) - Go server
- [frontend/README.md](frontend/README.md) - Vite + TypeScript
- [frontend/engine/README.md](frontend/engine/README.md) - Rust/Wasm
- [ingestion/README.md](ingestion/README.md) - Python jobs

## Environment Variables

| Variable | Default | Description |
|----------|---------|-------------|
| `PORT` | `8080` | Backend server port |
| `STATIC_DIR` | `../frontend/dist` | Path to built frontend |
| `NEO4J_URI` | `neo4j://localhost:7687` | Neo4j connection URI |
| `NEO4J_USER` | `neo4j` | Neo4j username |
| `NEO4J_PASSWORD` | (required) | Neo4j password |
| `BIGQUERY_PROJECT` | - | GCP project for GDELT queries |
