package conf42

import (
	"context"
	"io"
	"net/http"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestRunContainer(t *testing.T) {
	ctx := context.Background()

	conf42Container, err := RunContainer(ctx,
		WithIndex("testdata/index.html"),
	)
	require.NoError(t, err)
	t.Cleanup(func() {
		if err := conf42Container.Terminate(ctx); err != nil {
			t.Fatalf("failed to terminate container: %s", err)
		}
	})

	content, err := fetch(conf42Container.URI)
	require.NoError(t, err)
	require.Equal(t, "Conf42", content) // The value "Conf42" comes from "testdata/index.html"
}

func fetch(url string) (string, error) {
	req, err := http.NewRequest(http.MethodGet, url, nil)
	if err != nil {
		return "", err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return "", err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return "", err
	}

	return string(resBody), nil
}
