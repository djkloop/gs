package main

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"time"
)

func main() {
	rdb := coonRDB()

	err := rdb.Set(context.Background(), "session_id:admin", "session_id", 30*time.Second).Err()
	if err != nil {
		panic(err)
	}

	val, err := rdb.Get(context.Background(), "session_id:admin").Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(val)

}

func coonRDB() *redis.Client {
	rdb := redis.NewClient(&redis.Options{
		Addr:     "43.143.243.166:6379",
		Password: "bS9@xG2?", // no password set
		DB:       0,          // use default DB
	})

	_, err := rdb.Ping(context.Background()).Result()
	if err != nil {
		panic(err)
	}
	return rdb
}
