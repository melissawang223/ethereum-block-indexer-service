package redisrao

import "github.com/gomodule/redigo/redis"

// HSET sets field in the hash stored at key to value
func (r *RedisRao) HSET(key string, field string, value interface{}) {
	_, err := r.conn.Do("HSET", key, field, value)
	if err != nil {
		panic(err)
	}
}

// HGET gets value of a specific field of key
func (r *RedisRao) HGET(key string, field string) *string {
	value, err := r.conn.Do("HGET", key, field)
	if err != nil {
		panic(err)
	}

	result, err := redis.String(value, err)
	if err != nil {
		return nil
	}

	return &result
}

// HDEL Removes the specified fields from the hash stored at key
func (r *RedisRao) HDEL(key string, fieldValue interface{}) {
	var err error
	switch fieldValue.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("HDEL", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	default:
		_, err = r.conn.Do("HDEL", key, fieldValue)
	}

	if err != nil {
		panic(err)
	}
}

// HMSET sets multiple fields for a specific key
func (r *RedisRao) HMSET(key string, fieldValue interface{}) {
	_, err := r.conn.Do("HMSET", redis.Args{}.Add(key).AddFlat(fieldValue)...)
	if err != nil {
		panic(err)
	}
}

// HMGET returns the values associated with the specified fields in the hash stored at key.
func (r *RedisRao) HMGET(key string, fields interface{}) *[]string {
	var result []string

	value, err := r.conn.Do("HMGET", redis.Args{}.Add(key).AddFlat(fields)...)
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

// HGETALL get all fields for a given key
func (r *RedisRao) HGETALL(key string) *map[string]string {
	var result map[string]string

	value, err := r.conn.Do("HGETALL", key)

	if err != nil {
		panic(err)
	}

	switch value := value.(type) {
	case []interface{}:
		result = make(map[string]string, len(value))
		for i := 0; i < len(value); i += 2 {
			result[string(value[i].([]uint8))] = string(value[i+1].([]uint8))
		}
	}

	return &result
}
