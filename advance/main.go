package main

import (
	"fmt"
	"io"
	"net/http"
	"strconv"
)

const server = "http://localhost:3333"

func main() {
	// HTTP Server
	mux := http.NewServeMux()
	mux.HandleFunc("/", rootHandler(server))
	http.ListenAndServe(":4444", mux)
}

func rootHandler(server string) func(http.ResponseWriter, *http.Request) {
	fn := func(w http.ResponseWriter, r *http.Request) {
		c, err := increment(server)
		fmt.Fprintf(w, "Number of visitors: %d\n", c)
		if err != nil {
			fmt.Fprintf(w, "Err: %v\n", err)
		}
	}
	return fn
}

func increment(server string) (int64, error) {
	req, err := http.NewRequest(http.MethodGet, server, nil)
	if err != nil {
		return 0, err
	}

	res, err := http.DefaultClient.Do(req)
	if err != nil {
		return 0, err
	}

	resBody, err := io.ReadAll(res.Body)
	if err != nil {
		return 0, err
	}

	fmt.Printf(">> %s <<", string(resBody))
	i, err := strconv.Atoi(string(resBody))
	if err != nil {
		return 0, err
	}

	return int64(i), nil
}
