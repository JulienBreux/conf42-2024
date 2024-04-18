package main

import (
	"context"
	"testing"

	"github.com/redis/go-redis/v9"
	"github.com/stretchr/testify/require"
	"github.com/testcontainers/testcontainers-go"
	redisModule "github.com/testcontainers/testcontainers-go/modules/redis"
)

func TestIncrement(t *testing.T) {
	ctx := context.Background()

	redisContainer, err := redisModule.RunContainer(ctx, testcontainers.WithImage("redis:6"))
	require.NoError(t, err)

	endpoint, err := redisContainer.Endpoint(ctx, "")
	require.NoError(t, err)

	rdb := redis.NewClient(&redis.Options{
		Addr: endpoint,
	})

	_, err = rdb.Set(ctx, counterKey, "19", 0).Result()
	require.NoError(t, err)

	v := increment(ctx, rdb)
	require.Equal(t, int64(20), v)
}
