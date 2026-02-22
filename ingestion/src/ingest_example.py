"""
Example ingestion script for GDELT events from BigQuery.

This is a stub showing the intended flow. Implement actual queries
and processing logic as needed.
"""

from google.cloud import bigquery

from common.config import get_config
from common.neo4j_client import Neo4jClient


# Example GDELT query - modify based on your needs
GDELT_QUERY = """
SELECT
    GLOBALEVENTID as id,
    SQLDATE as date,
    Actor1Name as actor1,
    Actor2Name as actor2,
    EventCode as event_code,
    GoldsteinScale as goldstein_scale,
    NumMentions as mentions,
    AvgTone as tone,
    ActionGeo_FullName as location
FROM `gdelt-bq.gdeltv2.events`
WHERE SQLDATE >= @start_date
    AND NumMentions >= @min_mentions
ORDER BY NumMentions DESC
LIMIT @limit
"""


def fetch_gdelt_events(
    client: bigquery.Client,
    start_date: str,
    min_mentions: int = 10,
    limit: int = 100,
) -> list[dict]:
    """
    Fetch GDELT events from BigQuery.

    Args:
        client: BigQuery client
        start_date: YYYYMMDD format
        min_mentions: Minimum mentions threshold
        limit: Max events to fetch

    Returns:
        List of event dicts
    """
    job_config = bigquery.QueryJobConfig(
        query_parameters=[
            bigquery.ScalarQueryParameter("start_date", "STRING", start_date),
            bigquery.ScalarQueryParameter("min_mentions", "INT64", min_mentions),
            bigquery.ScalarQueryParameter("limit", "INT64", limit),
        ]
    )

    query_job = client.query(GDELT_QUERY, job_config=job_config)
    results = query_job.result()

    return [dict(row) for row in results]


def main() -> None:
    config = get_config()

    print("Loading GDELT events from BigQuery...")
    # Stub: In production, create actual BigQuery client
    # bq_client = bigquery.Client(project=config.bigquery_project)
    # events = fetch_gdelt_events(bq_client, start_date="20240101")

    # For now, use stub data
    events = [
        {
            "id": "example-event-1",
            "title": "Example Event",
            "date": "2024-01-01",
            "location": "Washington, DC",
            "importance": 0.5,
        }
    ]

    print(f"Fetched {len(events)} events")

    print("Writing events to Neo4j...")
    with Neo4jClient(
        config.neo4j_uri,
        config.neo4j_user,
        config.neo4j_password,
    ) as neo4j:
        for event in events:
            neo4j.upsert_event(event)
            print(f"  - {event['id']}")

    print("Done.")


if __name__ == "__main__":
    main()
