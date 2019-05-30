package app

import (
	"fmt"
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

type Server struct{}

func InitRedis(server *Server) error {

	db, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		return err
	}

	client := redis.NewClient(&redis.Options{
		Addr:     os.Getenv("REDIS_ADDRESS"),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})

	pong, err := client.Ping().Result()
	fmt.Println(pong, err)
	// Output: PONG <nil>
	return nil
}
