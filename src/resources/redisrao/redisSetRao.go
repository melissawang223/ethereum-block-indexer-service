package redisrao

import (
	"github.com/gomodule/redigo/redis"
)

// SADD add member to a set called key.  If key does not exist, a new set is created.
// param 'member' support array of string, int, int32, int64, float32, float64, interface{}
func (r *RedisRao) SADD(key string, member interface{}) {
	var err error
	switch member.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("SADD", redis.Args{}.Add(key).AddFlat(member)...)
	default:
		_, err = r.conn.Do("SADD", key, member)
	}

	if err != nil {
		panic(err)
	}
}

// SREM removes the specified members from the set stored at key.
// param 'member' support array of string, int, int32, int64, float32, float64, interface{}
func (r *RedisRao) SREM(key string, member interface{}) {
	var err error
	switch member.(type) {
	case []string, []int, []int32, []int64, []float32, []float64, []interface{}:
		_, err = r.conn.Do("SREM", redis.Args{}.Add(key).AddFlat(member)...)
	default:
		_, err = r.conn.Do("SREM", key, member)
	}

	if err != nil {
		panic(err)
	}
}

// SISMEMBER Returns if member is a member of the set stored at key.
func (r *RedisRao) SISMEMBER(key string, member interface{}) bool {
	exist, err := r.conn.Do("SISMEMBER", key, member)
	if err != nil {
		panic(err)
	}

	if exist.(int64) == 1 {
		return true
	}
	return false
}

// SMEMBERS get all members from the set called key
func (r *RedisRao) SMEMBERS(key string) *[]string {
	var result []string

	value, err := r.conn.Do("SMEMBERS", key)
	//fmt.Println("ERR", err)
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
