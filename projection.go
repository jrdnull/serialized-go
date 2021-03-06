package serialized

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

// Projection represents the model used to present your event data.
type Projection struct {
	ID   string          `json:"projectionId,omitempty"`
	Data json.RawMessage `json:"data,omitempty"`
}

// ProjectionDefinition contains the logic for how to build a projection.
type ProjectionDefinition struct {
	Name     string          `json:"projectionName,omitempty"`
	Feed     string          `json:"feedName,omitempty"`
	Handlers []*EventHandler `json:"handlers,omitempty"`
}

// EventHandler contains the functions for modifying the current state of a projection.
type EventHandler struct {
	EventType string      `json:"eventType,omitempty"`
	Functions []*Function `json:"functions,omitempty"`
}

// Function contains the templates for modifying projections.
type Function struct {
	Function       string      `json:"function,omitempty"`
	TargetSelector string      `json:"targetSelector,omitempty"`
	EventSelector  string      `json:"eventSelector,omitempty"`
	TargetFilter   string      `json:"targetFilter,omitempty"`
	EventFilter    string      `json:"eventFilter,omitempty"`
	RawData        interface{} `json:"rawData,omitempty"`
}

// ListProjectionDefinitions returns all definitions.
func (c *Client) ListProjectionDefinitions(ctx context.Context) ([]*ProjectionDefinition, error) {
	req, err := c.newRequest("GET", "/projections/definitions", nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Definitions []*ProjectionDefinition `json:"definitions"`
	}

	resp, err := c.do(ctx, req, &response)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return response.Definitions, err
}

// CreateProjectionDefinition creates a new reaction definition.
func (c *Client) CreateProjectionDefinition(ctx context.Context, d *ProjectionDefinition) error {
	req, err := c.newRequest("POST", "/projections/definitions", d)
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req, nil)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return err
}

// ProjectionDefinition returns a projection definition by name.
func (c *Client) ProjectionDefinition(ctx context.Context, name string) (*ProjectionDefinition, error) {
	req, err := c.newRequest("GET", "/projections/definitions/"+name, nil)
	if err != nil {
		return nil, err
	}

	var pd ProjectionDefinition

	resp, err := c.do(ctx, req, &pd)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return &pd, err
}

// DeleteProjectionDefinition deletes a projection definition.
func (c *Client) DeleteProjectionDefinition(ctx context.Context, name string) error {
	req, err := c.newRequest("DELETE", "/projections/definitions/"+name, nil)
	if err != nil {
		return err
	}

	resp, err := c.do(ctx, req, nil)
	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return err
}

// SingleProjection returns a single projection for the given aggregate.
func (c *Client) SingleProjection(ctx context.Context, projName, aggID string) (*Projection, error) {
	req, err := c.newRequest("GET", "/projections/single/"+projName+"/"+aggID, nil)
	if err != nil {
		return nil, err
	}

	var proj Projection

	resp, err := c.do(ctx, req, &proj)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return &proj, err
}

// ListSingleProjections returns all single projections.
func (c *Client) ListSingleProjections(ctx context.Context, name string) ([]*Projection, error) {
	req, err := c.newRequest("GET", "/projections/single/"+name, nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Projections []*Projection `json:"projections"`
	}

	resp, err := c.do(ctx, req, &response)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return response.Projections, err
}

// AggregatedProjection returns an aggregated projection for the given aggregate.
func (c *Client) AggregatedProjection(ctx context.Context, name string) (*Projection, error) {
	req, err := c.newRequest("GET", "/projections/aggregated/"+name, nil)
	if err != nil {
		return nil, err
	}

	var proj Projection

	resp, err := c.do(ctx, req, &proj)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return &proj, err
}

// ListAggregatedProjections returns all single projections.
func (c *Client) ListAggregatedProjections(ctx context.Context) ([]*Projection, error) {
	req, err := c.newRequest("GET", "/projections/aggregated", nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Projections []*Projection `json:"projections"`
	}

	resp, err := c.do(ctx, req, &response)
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return response.Projections, err
}
