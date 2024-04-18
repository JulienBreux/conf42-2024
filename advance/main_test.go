package main

import (
	"context"
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

	endpoint, err := microcksContainer.HttpEndpoint(ctx)
	require.NoError(t, err)

	baseApiUrl, err := microcksContainer.RestMockEndpoint(ctx, "CounterAPI", "0.0.1")
	require.NoError(t, err)
	require.Equal(t, endpoint+"/rest/CounterAPI/0.0.1", baseApiUrl)

	v, err := increment(baseApiUrl)
	require.NoError(t, err)
	require.Equal(t, int64(20), v)
}
