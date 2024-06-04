package main

import (
	"log"
	"os"

	"github.com/redis/go-redis/v9"
)

// TODO Move to json config file
type ServiceConfig struct {
	GrpcServePort uint16 `json:"grpcServePort"`
	HttpServePort uint16 `json:"httpServePort"`
}

func main() {
	// ctx := context.Background()

	redisConnStr, exist := os.LookupEnv("REDIS_CONN_STR")
	if !exist {
		log.Fatalln("Redis connection string now found in environment variables")
	}

	opt, err := redis.ParseURL(redisConnStr)
	fatalErrCheck(err, "Redis url invalid")
	redisClient := redis.NewClient(opt)

	_ = redisClient
}

func fatalErrCheck(err error, msg string) {
	if err != nil {
		log.Fatalf("%s: %v", msg, err)
	}
}
