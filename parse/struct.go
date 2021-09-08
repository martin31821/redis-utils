package parse

import (
	"errors"

	"github.com/go-redis/redis/v8"
	"github.com/mitchellh/mapstructure"
)

// ExtractStructFromRedisMessage tries to fill a given struct with the given redis message
// automagically performing all the necessary type conversions.
// The mapping is done via struct tags (specifically `rediskey`).
func ExtractStructFromRedisMessage(msg *redis.XMessage, structOut interface{}) error {
  decoder, err := mapstructure.NewDecoder(&mapstructure.DecoderConfig{
    WeaklyTypedInput: true,
    TagName: "rediskey",
    ErrorUnused: false,
    ZeroFields: false,
    Result: structOut,
  });
  if err != nil {
    return err;
  }
  return decoder.Decode(msg.Values)
}

// ExtractStructFromRedisSliceCmd is a shorthand to parse the last message from a messageSlice into an arbitrary struct
func ExtractStructFromRedisSliceCmd(msg *redis.XMessageSliceCmd, structOut interface{}) error {
  msgs := msg.Val()
  if len(msgs) == 0 {
    return errors.New("no message to parse")
  }
  return ExtractStructFromRedisMessage(&msgs[len(msgs) - 1], structOut)
}
