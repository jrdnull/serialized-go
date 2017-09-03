package serialized

import (
	"bytes"
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httptest"
	"testing"
)

func TestListReactions(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte(`{
  "reactions": [
    {
      "id": "be278b27-8687-42b4-a502-164a6702797c",
      "name": "PaymentProcessedEmailReaction",
      "feed": "payment",
      "eventType": "PaymentProcessed",
      "action": {
        "httpMethod": "POST",
        "targetUri": "https://your-webhook",
        "body": "A new payment was processed",
        "actionType": "HTTP"
      }
    }
  ]
}`))
	}))

	c := NewClient(
		WithBaseURL(ts.URL),
	)

	r, err := c.ListReactions()
	if err != nil {
		t.Fatal(err)
	}

	if len(r) != 1 {
		t.Errorf("unexpected number of reactions = %d; want = %d", len(r), 1)
	}
}

func TestCreateReaction(t *testing.T) {
	ts := httptest.NewServer(http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		want := []byte(`{"id":"be278b27-8687-42b4-a502-164a6702797c","name":"PaymentProcessedEmailReaction","feed":"payment","eventType":"PaymentProcessed","action":{"httpMethod":"POST","targetUri":"https://your-webhook","body":"A new payment was processed","actionType":"HTTP"}}`)

		got, err := ioutil.ReadAll(r.Body)
		if err != nil {
			t.Fatal(err)
		}

		got = bytes.TrimSpace(got)

		if !bytes.Equal(got, want) {
			var wantIndent bytes.Buffer
			json.Indent(&wantIndent, want, "", "\t")

			var gotIndent bytes.Buffer
			json.Indent(&gotIndent, got, "", "\t")

			t.Errorf("unexpected request body =\n%s\n\nwant =\n%s", got, want)
		}

		w.WriteHeader(http.StatusCreated)
	}))

	c := NewClient(
		WithBaseURL(ts.URL),
	)

	r := Reaction{
		ID:        "be278b27-8687-42b4-a502-164a6702797c",
		Name:      "PaymentProcessedEmailReaction",
		Feed:      "payment",
		EventType: "PaymentProcessed",
		Action: Action{
			HTTPMethod: "POST",
			TargetURI:  "https://your-webhook",
			Body:       "A new payment was processed",
			ActionType: "HTTP",
		},
	}

	if err := c.CreateReaction(r); err != nil {
		t.Fatal(err)
	}
}
