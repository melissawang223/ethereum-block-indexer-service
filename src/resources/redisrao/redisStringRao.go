package redisrao

import (
	"github.com/gomodule/redigo/redis"
)

// SETEX sets a key-value pair with expire timeutil
func (r *RedisRao) SETEX(key string, ttl int, value interface{}) {
	_, err := r.conn.Do("SETEX", key, ttl, value)
	if err != nil {
		panic(err)
	}
}

// GET gets value of given key
func (r *RedisRao) GET(key string) *string {
	value, err := r.conn.Do("GET", key)
	if err != nil {
		panic(err)
	}

	result, err := redis.String(value, err)
	if err != nil {
		return nil
	}

	return &result
}

// MGET Returns the values of all specified keys. For every key that
// does not hold a string value or does not exist, the special value nil is returned.
// Because of this, the operation never fails.
func (r *RedisRao) MGET(key []string) *[]string {
	var result []string

	value, err := r.conn.Do("MGET", redis.Args{}.AddFlat(key)...)
	if err != nil {
		panic(err)
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			switch v := value[i].(type) {
			case []uint8:
				result[i] = string(v)
			}
		}
	default:
		return nil
	}

	return &result
}
