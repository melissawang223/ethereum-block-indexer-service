package redisrao

import "github.com/gomodule/redigo/redis"

// ZADD add value in a sorted-set(key) with score
func (r *RedisRao) ZADD(key string, score string, value string) {
	_, err := r.conn.Do("ZADD", key, score, value)

	if err != nil {
		panic(err)
	}
}

// ZRANGE return members of sorted-set(key) between start and stop.
// All members: start = 0, end = -1
func (r *RedisRao) ZRANGE(key string, start, stop int) *[]string {
	var result []string
	value, err := r.conn.Do("ZRANGE", key, start, stop)

	if err != nil {
		panic(err)
	}

	switch value := value.(type) {
	case []interface{}:
		result = make([]string, len(value))
		for i := 0; i < len(value); i++ {
			result[i] = string(value[i].([]uint8))
		}
	default:
		return nil
	}

	return &result
}

// ZCARD return the num of member within key
func (r *RedisRao) ZCARD(key string) int {
	value, err := r.conn.Do("ZCARD", key)
	if err != nil {
		panic(err)
	}

	result, err := redis.Int(value, err)
	if err != nil {
		panic(err)
	}

	return result
}

// ZREM remove members within key
func (r *RedisRao) ZREM(key string, members interface{}) {
	_, err := r.conn.Do("ZREM", redis.Args{}.Add(key).AddFlat(members)...)
	if err != nil {
		panic(err)
	}
}
