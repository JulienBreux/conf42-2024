package main

import (
	"context"
	"fmt"
	"testing"

	"github.com/stretchr/testify/require"
	microcksModule "microcks.io/testcontainers-go"
)

func TestIncrement(t *testing.T) {
	ctx := context.Background()

	microcksContainer, err := microcksModule.RunContainer(ctx)
	require.NoError(t, err)
	_, err = microcksContainer.ImportAsMainArtifact(ctx, "./testdata/counter-api.yaml")
	require.NoError(t, err)

	uri, err := microcksCounterURI(ctx, microcksContainer)
	require.NoError(t, err)

	v, err := increment(uri)
	require.NoError(t, err)
	require.Equal(t, int64(20), v)
}

func microcksCounterURI(ctx context.Context, microcksContainer *microcksModule.MicrocksContainer) (string, error) {
	ip, err := microcksContainer.Host(ctx)
	if err != nil {
		return "", err
	}

	mappedPort, err := microcksContainer.MappedPort(ctx, "8080")
	if err != nil {
		return "", err
	}

	uri := fmt.Sprintf("http://%s:%s/rest/Counter+API/1.0.0/", ip, mappedPort.Port())
	return uri, nil
}
