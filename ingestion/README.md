# Ingestion (Python)

Offline ingestion jobs for fetching GDELT events from BigQuery and writing to Neo4j.

## Prerequisites

- Python 3.11+
- Neo4j instance running
- GCP credentials (for BigQuery access)

## Setup

```bash
python -m venv .venv
source .venv/bin/activate
pip install -e .
```

## Configuration

Set environment variables (or use `../.env`):

```bash
export NEO4J_URI=neo4j://localhost:7687
export NEO4J_USER=neo4j
export NEO4J_PASSWORD=your-password
export BIGQUERY_PROJECT=your-gcp-project
```

For BigQuery, also configure GCP credentials:
```bash
export GOOGLE_APPLICATION_CREDENTIALS=/path/to/service-account.json
```

## Run

```bash
python src/ingest_example.py
```

## Scripts

| Script | Description |
|--------|-------------|
| `ingest_example.py` | Fetch GDELT events from BigQuery, write to Neo4j |
| `classify_importance.py` | Event importance scoring utilities |

## Directory Structure

```
ingestion/
├── src/
│   ├── common/
│   │   ├── config.py         # Environment config
│   │   └── neo4j_client.py   # Neo4j write client
│   ├── ingest_example.py     # Main ingestion script
│   └── classify_importance.py # Scoring logic
└── pyproject.toml
```
