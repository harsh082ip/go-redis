package main

import (
	"context"
	"fmt"
	"net/http"

	"github.com/redis/go-redis/v9"
)

var client *redis.Client
var ctx = context.Background()

const WEBPORT = ":8001"

func main() {

	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "",
		DB:       0,
	})
	defer client.Close()

	client.Set(ctx, "visits", 0, 0)

	http.HandleFunc("/", getVisits)

	fmt.Println("Listening on Port ", WEBPORT)
	http.ListenAndServe(WEBPORT, nil)

}
