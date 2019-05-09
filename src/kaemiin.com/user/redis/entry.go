package main

import (
	"errors"
	"fmt"
	"os"
	"strconv"
	"github.com/joho/godotenv"
	"github.com/go-redis/redis"
)

const REDIS_PORT = ":6379"

func init() {
	fmt.Println("redis/entry.go ==> init()")
	err := godotenv.Load()
	if err != nil {
		panic(errors.New("Error loading .env file"))
	}
}

func getKeys(redisdb *redis.Client) int {
	var cursor uint64
	var n int
	for {
		var keys []string
		var err error
		keys, cursor, err = redisdb.Scan(cursor, "key*", 10).Result()
		if err != nil {
			panic(err)
		}
		n += len(keys)
		if cursor == 0 {
			break
		}
	}
	return n
}

func main() {
	fmt.Println("redis/entry.go ==> main()")
	var redisdb *redis.Client
	REDIS_URL := "localhost" + REDIS_PORT
	REDIS_HOST := os.Getenv("REDIS_HOST")
	if REDIS_HOST != "" {
		REDIS_URL = REDIS_HOST + REDIS_PORT
	}
	REDIS_DB, err := strconv.Atoi(os.Getenv("REDIS_DB"))
	if err != nil {
		panic(err)
	}
	redisdb = redis.NewClient(&redis.Options{
		Addr:     REDIS_URL ,
		Password: "",
		DB:       REDIS_DB,
		PoolSize: 1,
	})
	defer redisdb.Close()
	pong, err := redisdb.Ping().Result()
	if err != nil {
		panic(err)
	}
	fmt.Println(pong)
	counts := getKeys(redisdb)
	fmt.Println(counts)
}
