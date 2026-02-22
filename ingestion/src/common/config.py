from pydantic_settings import BaseSettings


class Config(BaseSettings):
    """Configuration loaded from environment variables."""

    neo4j_uri: str = "neo4j://localhost:7687"
    neo4j_user: str = "neo4j"
    neo4j_password: str

    bigquery_project: str = ""

    class Config:
        env_file = "../.env"
        env_file_encoding = "utf-8"


def get_config() -> Config:
    """Load configuration from environment."""
    return Config()
