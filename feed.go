package serialized

import (
	"context"
	"fmt"
	"net/http"
	"net/url"
	"strconv"
)

// Feed holds a Serialized.io feed.
type Feed struct {
	Entries []*FeedEntry `json:"entries"`
	HasMore bool         `json:"hasMore"`
}

// FeedEntry holds a Serialized.io feed entry.
type FeedEntry struct {
	SequenceNumber int64
	AggregateID    string
	Timestamp      int64
	Events         []*Event
}

// Feeds returns all feed types.
func (c *Client) Feeds(ctx context.Context) ([]string, error) {
	req, err := c.newRequest("GET", "/feeds/", nil)
	if err != nil {
		return nil, err
	}

	var response struct {
		Feeds []string `json:"feeds"`
	}

	resp, err := c.do(ctx, req, &response)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return response.Feeds, nil
}

// Feed returns the feed for a given aggregate.
func (c *Client) Feed(ctx context.Context, name string, since int64) (*Feed, error) {
	u := &url.URL{
		Path: "/feeds/" + name,
	}

	if since > 0 {
		vs := make(url.Values)
		vs.Set("since", fmt.Sprintf("%d", since))
		u.RawQuery = vs.Encode()
	}

	req, err := c.newRequest("GET", u.String(), nil)
	if err != nil {
		return nil, err
	}

	f := new(Feed)
	resp, err := c.do(ctx, req, &f)
	if err != nil {
		return nil, err
	}

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	return f, nil
}

// FeedSequenceNumber returns current sequence number at head for a given feed.
func (c *Client) FeedSequenceNumber(ctx context.Context, feedName string) (int64, error) {
	req, err := c.newRequest("HEAD", "/feeds/"+feedName, nil)
	if err != nil {
		return 0, err
	}

	resp, err := c.do(ctx, req, nil)
	if err != nil {
		return 0, err
	}

	if resp.StatusCode != http.StatusOK {
		return 0, fmt.Errorf("unexpected status code: %d", resp.StatusCode)
	}

	seqstr := resp.Header.Get("Current-Sequence-Number")
	seq, err := strconv.ParseInt(seqstr, 10, 64)
	if err != nil {
		return 0, err
	}

	return seq, nil
}
