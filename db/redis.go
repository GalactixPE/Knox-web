package db

import (
	"github.com/go-redis/redis/v8"
	"os"
)

var rdb = redis.NewClient(&redis.Options{
	Addr:     os.Getenv("HOST") + ":" + os.Getenv("PORT"),
	Password: os.Getenv("PASSWORD"), // no password set
	DB:       0,                     // use default DB
})

func Client() *redis.Client {
	return rdb
}
