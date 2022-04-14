package redisrao

import (
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/gomodule/redigo/redis"
)

// Close redis connection
func (r *RedisRao) Close() {
	r.conn.Close()
}

// KEYS return all value whose key is matching with pattern
func (r *RedisRao) KEYS(pattern string) *[]string {
	value, err := r.conn.Do("KEYS", pattern)
	if err != nil {
		panic(err)
	}

	keys, err := redis.Strings(value, err)
	if err != nil {
		return nil
	}
	return &keys
}

// EXISTS returns if key exists
func (r *RedisRao) EXISTS(key string) bool {
	value, err := r.conn.Do("EXISTS", key)
	if err != nil {
		panic(err)
	}

	result, err := redis.Bool(value, err)
	if err != nil {
		return false
	}
	return result
}

// EXPIRE set a timeout on key
func (r *RedisRao) EXPIRE(key string, ttl int) {
	_, err := r.conn.Do("EXPIRE", key, ttl)

	if err != nil {
		panic(err)
	}
}

// DEL removes the specified keys. A key is ignored if it does not exist.
// Delete multiple keys should use []string for parameter {key}
func (r *RedisRao) DEL(key interface{}) {
	var err error
	switch key.(type) {
	case []string:
		_, err = r.conn.Do("DEL", redis.Args{}.AddFlat(key)...)
	default:
		_, err = r.conn.Do("DEL", key)
	}

	if err != nil {
		panic(err)
	}
}

// MULTI start a transaction in redis, queue later command
func (r *RedisRao) MULTI() *error {
	log := logger.Logger()

	_, err := r.conn.Do("MULTI")
	if err != nil {
		log.Error(err)
		return &err
	}
	return nil
}

// EXEC all commands in queue and end the transaction
func (r *RedisRao) EXEC() *error {
	log := logger.Logger()

	_, err := r.conn.Do("EXEC")
	if err != nil {
		log.Error(err)
		return &err
	}
	return nil
}

// DISCARD all commands in queue and end the transaction
func (r *RedisRao) DISCARD() *error {
	log := logger.Logger()

	_, err := r.conn.Do("DISCARD")
	if err != nil {
		log.Error(err)
		return &err
	}
	return nil
}
