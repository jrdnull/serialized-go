package serialized

import (
	"fmt"
	"net/http"
)

// Reaction holds a Serialized.io Reaction.
type Reaction struct {
	ID        string `json:"id"`
	Name      string `json:"name"`
	Feed      string `json:"feed"`
	EventType string `json:"eventType"`
	Action    Action `json:"action"`
}

// Action holds a Serialized.io Action.
type Action struct {
	HTTPMethod string `json:"httpMethod"`
	TargetURI  string `json:"targetUri"`
	Body       string `json:"body"`
	ActionType string `json:"actionType"`
}

// CreateReaction registers a new reaction.
func (c *Client) CreateReaction(r Reaction) error {
	req, err := c.newRequest("POST", "/reactions", r)
	if err != nil {
		return err
	}

	resp, err := c.do(req, nil)
	if resp.StatusCode != http.StatusCreated {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return err
}

// ListReactions returns all registered reactions.
func (c *Client) ListReactions() ([]Reaction, error) {
	req, err := c.newRequest("GET", "/reactions", nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Reactions []Reaction `json:"reactions"`
	}

	resp, err := c.do(req, &response)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return response.Reactions, err
}
