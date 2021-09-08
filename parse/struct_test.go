package parse

import (
	"testing"

	"github.com/go-redis/redis/v8"
)

type temp struct {
  Foobar float64 `rediskey:"barfoo"`
}

func TestStructParse(t *testing.T) {
  testMsg := &redis.XMessage{
    ID: "42",
    Values: map[string]interface{}{
      "barfoo": "42.1",
    },
  };
  value := &temp{}
  if err := ExtractStructFromRedisMessage(testMsg, value); err != nil {
    t.FailNow()
  }
  t.Log(value)
}
