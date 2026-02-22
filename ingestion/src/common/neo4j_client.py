from typing import Any
from neo4j import GraphDatabase


class Neo4jClient:
    """Client for reading/writing story and event data to Neo4j."""

    def __init__(self, uri: str, user: str, password: str):
        self._driver = GraphDatabase.driver(uri, auth=(user, password))
        self._driver.verify_connectivity()

    def close(self) -> None:
        self._driver.close()

    def __enter__(self) -> "Neo4jClient":
        return self

    def __exit__(self, *args: Any) -> None:
        self.close()

    def upsert_event(self, event_data: dict[str, Any]) -> None:
        """
        Upsert an event node in Neo4j.

        Args:
            event_data: Dict with keys: id, title, date, location, etc.
        """
        with self._driver.session() as session:
            session.run(
                """
                MERGE (e:Event {id: $id})
                SET e.title = $title,
                    e.date = $date,
                    e.location = $location,
                    e.importance = $importance,
                    e.updatedAt = datetime()
                """,
                **event_data,
            )

    def upsert_story(self, story_data: dict[str, Any]) -> None:
        """
        Upsert a story node in Neo4j.

        Args:
            story_data: Dict with keys: id, title, summary, etc.
        """
        with self._driver.session() as session:
            session.run(
                """
                MERGE (s:Story {id: $id})
                SET s.title = $title,
                    s.summary = $summary,
                    s.createdAt = coalesce(s.createdAt, datetime()),
                    s.updatedAt = datetime()
                """,
                **story_data,
            )

    def link_event_to_story(self, event_id: str, story_id: str) -> None:
        """Create relationship between event and story."""
        with self._driver.session() as session:
            session.run(
                """
                MATCH (e:Event {id: $event_id})
                MATCH (s:Story {id: $story_id})
                MERGE (e)-[:PART_OF]->(s)
                """,
                event_id=event_id,
                story_id=story_id,
            )
