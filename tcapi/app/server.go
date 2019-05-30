package app

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
)

type Server struct{}

func ExampleNewClient() {
	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: "", // no password set
		DB:       0,  // use default DB
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
}
