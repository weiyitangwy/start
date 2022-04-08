package main

import (
	"fmt"
	"io"
	"net/http"
	"testing"

	"github.com/ory/dockertest/v3"
	"github.com/stretchr/testify/require"
)

func TestMain(t *testing.T) {

	pool, err := dockertest.NewPool("")
	require.NoError(t, err, "could not connect to Docker")
	fmt.Println(pool)

	resource, err := pool.Run("helloserver", "latest", []string{})
	require.NoError(t, err, "could not start container")

	t.Cleanup(func() {
		require.NoError(t, pool.Purge(resource), "failed to remove container")
	})

	var resp *http.Response

	err = pool.Retry(func() error {
		fmt.Println("pool", resource.GetPort("8080/tcp"))
		resp, err = http.Get("http://localhost:8080/")
		fmt.Println("pool", err, resp)
		if err != nil {
			t.Log("container not ready, waiting...")
			return err
		}
		return nil
	})
	require.NoError(t, err, "HTTP error")
	defer resp.Body.Close()

	require.Equal(t, http.StatusOK, resp.StatusCode, "HTTP status code")

	body, err := io.ReadAll(resp.Body)
	require.NoError(t, err, "failed to read HTTP body")

	// Finally, test the business requirement!
	require.Contains(t, string(body), "World", "does not respond with love?")
}
