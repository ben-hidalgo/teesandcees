package app

import (
	"github.com/go-redis/redis"
	"os"
	"strconv"
)

type Server struct {
	Rc *redis.Client
}

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

	_, err = client.Ping().Result()
	if err != nil {
		return err
	}

	server.Rc = client

	return nil
}
