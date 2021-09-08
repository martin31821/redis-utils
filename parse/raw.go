package parse

import (
	"fmt"
	"math"
	"strconv"
)

/**
This file contains low-level methods to interact with redis values
*/

//GetStringFromRedisMap reads a string with the given key from the redis map
func GetStringFromRedisMap(redisMap map[string]interface{}, keyName string) (string, error) {
	entryVal, ok := redisMap[keyName]
	if !ok {
		return "", fmt.Errorf("entry %v is missing in kv map", keyName)
	}
	valStr, ok := entryVal.(string)
	if !ok {
		return "", fmt.Errorf("entry %v is not a string", keyName)
	}

	return valStr, nil
}

//GetFloatFromRedisMap tries to find the entry with name keyName in the redisMap
//Assumes that the entry is a string and then tries to parse it as a float
//This uses the same semantics as strconv.ParseFloat, i.e. always returns a float64, but if you specify
//bitSize 32, it is convertible to 32bit float without loss of precision
func GetFloatFromRedisMap(redisMap map[string]interface{}, keyName string, bitSize int) (float64, error) {
	valStr, err := GetStringFromRedisMap(redisMap, keyName)
	if err != nil {
		return math.NaN(), err
	}

	valFloat, err := strconv.ParseFloat(valStr, bitSize)
	if err != nil {
		return math.NaN(), fmt.Errorf("could not parse entry %v (%v) as float", keyName, valStr)
	}

	return valFloat, nil
}

//GetBoolFromRedisMap tries to find the entry with name keyName in the redisMap and parse it as bool
//ONLY accepts "true" and "false"
func GetBoolFromRedisMap(redisMap map[string]interface{}, keyName string) (bool, error) {
	valStr, err := GetStringFromRedisMap(redisMap, keyName)
	if err != nil {
		return false, err
	}

  return strconv.ParseBool(valStr)
}
