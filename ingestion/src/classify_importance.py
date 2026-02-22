"""
Event importance scoring/classification.

This is a stub. Implement your own scoring logic based on:
- Number of mentions
- Goldstein scale
- Actor significance
- Geographic relevance
- Temporal factors
"""


def score_event(event: dict) -> float:
    """
    Calculate importance score for an event.

    Args:
        event: Event dict with fields like mentions, goldstein_scale, tone

    Returns:
        Importance score between 0.0 and 1.0
    """
    # Stub implementation - replace with actual scoring logic
    mentions = event.get("mentions", 0)
    goldstein = abs(event.get("goldstein_scale", 0))

    # Simple weighted combination (placeholder)
    mention_score = min(mentions / 100, 1.0)
    impact_score = goldstein / 10.0

    score = (mention_score * 0.6) + (impact_score * 0.4)

    return min(max(score, 0.0), 1.0)


def classify_event_type(event: dict) -> str:
    """
    Classify event into a category.

    Args:
        event: Event dict with event_code field

    Returns:
        Category string
    """
    event_code = str(event.get("event_code", ""))

    # CAMEO event code classification (simplified)
    if event_code.startswith("01"):
        return "statement"
    elif event_code.startswith("02"):
        return "appeal"
    elif event_code.startswith("03"):
        return "cooperation"
    elif event_code.startswith("04"):
        return "consultation"
    elif event_code.startswith("05"):
        return "diplomacy"
    elif event_code.startswith("1"):
        return "conflict_verbal"
    elif event_code.startswith("2"):
        return "conflict_physical"
    else:
        return "other"
