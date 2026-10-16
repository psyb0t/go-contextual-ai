package contextual_test

import (
	"context"
	"net/http"
	"os"

	"github.com/psyb0t/go-contextual-ai"
)

// This is the README usage example. It talks to the live API, so it has no
// Output comment and only has to compile.
//
//nolint:testableexamples // needs real credentials, compile check only
func Example() {
	auth := contextual.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set(
				"Authorization",
				"Bearer "+os.Getenv("CONTEXTUAL_API_KEY"),
			)

			return nil
		},
	)

	client, err := contextual.NewClientWithResponses(
		"https://api.contextual.ai/v1",
		auth,
	)
	if err != nil {
		panic(err)
	}

	res, err := client.ListDatastoresDatastoresGetWithResponse(
		context.Background(),
		&contextual.ListDatastoresDatastoresGetParams{},
	)
	if err != nil {
		panic(err)
	}

	_ = res.JSON200
}
