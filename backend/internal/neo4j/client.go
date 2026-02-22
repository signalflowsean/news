package neo4j

import (
	"context"
	"time"

	"github.com/neo4j/neo4j-go-driver/v5/neo4j"
)

type Story struct {
	ID        string    `json:"id"`
	Title     string    `json:"title"`
	Summary   string    `json:"summary,omitempty"`
	CreatedAt time.Time `json:"createdAt"`
}

type Client struct {
	driver neo4j.DriverWithContext
}

func NewClient(ctx context.Context, uri, user, password string) (*Client, error) {
	driver, err := neo4j.NewDriverWithContext(uri, neo4j.BasicAuth(user, password, ""))
	if err != nil {
		return nil, err
	}

	if err := driver.VerifyConnectivity(ctx); err != nil {
		driver.Close(ctx)
		return nil, err
	}

	return &Client{driver: driver}, nil
}

func (c *Client) GetStories(ctx context.Context) ([]Story, error) {
	session := c.driver.NewSession(ctx, neo4j.SessionConfig{
		AccessMode: neo4j.AccessModeRead,
	})
	defer session.Close(ctx)

	result, err := session.Run(ctx,
		`MATCH (s:Story)
		 RETURN s.id AS id, s.title AS title, s.summary AS summary, s.createdAt AS createdAt
		 ORDER BY s.createdAt DESC
		 LIMIT 100`,
		nil,
	)
	if err != nil {
		return nil, err
	}

	var stories []Story
	for result.Next(ctx) {
		record := result.Record()

		id, _ := record.Get("id")
		title, _ := record.Get("title")
		summary, _ := record.Get("summary")
		createdAt, _ := record.Get("createdAt")

		story := Story{
			ID:    toString(id),
			Title: toString(title),
		}

		if summary != nil {
			story.Summary = toString(summary)
		}

		if createdAt != nil {
			if t, ok := createdAt.(time.Time); ok {
				story.CreatedAt = t
			}
		}

		stories = append(stories, story)
	}

	if err := result.Err(); err != nil {
		return nil, err
	}

	if stories == nil {
		stories = []Story{}
	}

	return stories, nil
}

func (c *Client) Close(ctx context.Context) error {
	return c.driver.Close(ctx)
}

func toString(v interface{}) string {
	if v == nil {
		return ""
	}
	if s, ok := v.(string); ok {
		return s
	}
	return ""
}
