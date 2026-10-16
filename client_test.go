package contextual_test

import (
	"context"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/psyb0t/go-contextual-ai"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

const (
	testAPIKey      = "test-key"
	testAgentID     = "6f1c2a4e-0b3d-4c5e-8f9a-1b2c3d4e5f60"
	testDatastoreID = "0a1b2c3d-4e5f-4061-8293-a4b5c6d7e8f9"
)

func newTestClient(
	t *testing.T,
	handler http.HandlerFunc,
) *contextual.ClientWithResponses {
	t.Helper()

	srv := httptest.NewServer(handler)
	t.Cleanup(srv.Close)

	auth := contextual.WithRequestEditorFn(
		func(_ context.Context, req *http.Request) error {
			req.Header.Set("Authorization", "Bearer "+testAPIKey)

			return nil
		},
	)

	client, err := contextual.NewClientWithResponses(srv.URL, auth)
	require.NoError(t, err)

	return client
}

func TestListDatastores(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodGet, r.Method)
		assert.Equal(t, "/datastores", r.URL.Path)
		assert.Equal(t, "5", r.URL.Query().Get("limit"))
		assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"total_count":1,"datastores":[{"id":"` + testDatastoreID +
				`","name":"docs","datastore_type":"UNSTRUCTURED",` +
				`"created_at":"2026-01-02T03:04:05Z"}]}`,
		))
	})

	limit := 5

	res, err := client.ListDatastoresDatastoresGetWithResponse(
		context.Background(),
		&contextual.ListDatastoresDatastoresGetParams{Limit: &limit},
	)
	require.NoError(t, err)
	require.Equal(t, http.StatusOK, res.StatusCode())
	require.NotNil(t, res.JSON200)

	assert.Equal(t, 1, res.JSON200.TotalCount)
	require.Len(t, res.JSON200.Datastores, 1)

	ds := res.JSON200.Datastores[0]
	assert.Equal(t, testDatastoreID, ds.Id.String())
	assert.Equal(t, "docs", ds.Name)
}

func TestCreateAgent(t *testing.T) {
	t.Parallel()

	client := newTestClient(t, func(w http.ResponseWriter, r *http.Request) {
		assert.Equal(t, http.MethodPost, r.Method)
		assert.Equal(t, "/agents", r.URL.Path)
		assert.Equal(t, "application/json", r.Header.Get("Content-Type"))
		assert.Equal(t, "Bearer "+testAPIKey, r.Header.Get("Authorization"))

		var body contextual.CreateAgentRequest

		assert.NoError(t, json.NewDecoder(r.Body).Decode(&body))
		assert.Equal(t, "support", body.Name)

		w.Header().Set("Content-Type", "application/json")
		_, _ = w.Write([]byte(
			`{"id":"` + testAgentID + `","datastore_ids":["` +
				testDatastoreID + `"]}`,
		))
	})

	res, err := client.CreateAgentAgentsPostWithResponse(
		context.Background(),
		contextual.CreateAgentRequest{Name: "support"},
	)
	require.NoError(t, err)
	require.NotNil(t, res.JSON200)

	assert.Equal(t, testAgentID, res.JSON200.Id.String())
	require.Len(t, res.JSON200.DatastoreIds, 1)
	assert.Equal(t, testDatastoreID, res.JSON200.DatastoreIds[0].String())
}
