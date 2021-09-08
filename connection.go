package hlredis

import (
	"context"
	"fmt"
	"log"
	"os"
	"time"

	"github.com/go-redis/redis/v8"
)

// TODO: this file should contain a common redis connecton setup routine
// that honors common environment variables and configuration
// like a sentinel or a clustered redis.

func NewCommonConfigRedisClient() (*redis.Client, error) {
  redisURL, present := os.LookupEnv("REDIS_URL")
	if !present {
    return nil, fmt.Errorf("REDIS_URL not specified in environment")
	}

  user, userPresent := os.LookupEnv("REDIS_USER")
  pass, passPresent := os.LookupEnv("REDIS_PASS")

  opts := &redis.Options{
		Addr:     redisURL,
    DB:       0,  // use default DB
	}

  if userPresent && passPresent {
    opts.Username = user
    opts.Password = pass
  }

  rdb := redis.NewClient(opts)
  ctx, cancel := context.WithTimeout(context.Background(), 100 * time.Millisecond)
  defer cancel()
	_, err := rdb.Ping(ctx).Result()
	if err != nil {
    return nil, fmt.Errorf("could not establish connection to redis: %v", err)
	}
	log.Printf("[INFO] Successfully connected to redis!")
  return rdb, nil
}
