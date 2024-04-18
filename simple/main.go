package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

const counterKey = "counter"

func main() {
	ctx := context.Background()

	// Redis
	rdb := redis.NewClient(&redis.Options{})

	// HTTP Server
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler(ctx, rdb))
	http.ListenAndServe(":3333", mux)
}

func rootHandler(ctx context.Context, rdb *redis.Client) func(http.ResponseWriter, *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c := increment(ctx, rdb)
		fmt.Fprint(w, c)
	}
	return fn
}

func increment(ctx context.Context, rdb *redis.Client) int64 {
	c, _ := rdb.Incr(ctx, counterKey).Result()
	return c
}
