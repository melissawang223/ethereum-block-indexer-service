package redisrao

import "github.com/gomodule/redigo/redis"

// RPUSH push a value to the end of the list
func (r *RedisRao) RPUSH(key string, value string) {
	_, err := r.conn.Do("RPUSH", key, value)
	if err != nil {
		panic(err)
	}
}

// LLEN return len of key
func (r *RedisRao) LLEN(key string) int {
	value, err := r.conn.Do("LLEN", key)
	if err != nil {
		panic(err)
	}

	result, err := redis.Int(value, err)
	if err != nil {
		return 0
	}

	return result
}

// LRANGE return values from start to end of key
func (r *RedisRao) LRANGE(key string, start, end int) *[]string {
	value, err := r.conn.Do("LRANGE", key, start, end)
	if err != nil {
		panic(err)
	}

	result, err := redis.Strings(value, err)
	if err != nil {
		return nil
	}

	return &result
}
