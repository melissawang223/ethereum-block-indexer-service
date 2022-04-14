package redisrao

import (
	"fmt"
	"time"

	"github.com/ethereum-block-indexer-service/src/helpers/config"
	"github.com/ethereum-block-indexer-service/src/helpers/logger"
	"github.com/gomodule/redigo/redis"
)

var pool *redis.Pool
var r *RedisRao

func init() {
	connectPool()
	Redis()
}

// RedisRao is an alias of redis.Conn, for which imports redisrao to use
type RedisRao struct {
	conn redis.Conn
}

// Redis retuen a redis connect from pool
func Redis() *RedisRao {
	//	if r == nil {
	r = &RedisRao{
		conn: pool.Get(),
	}
	//}
	return r
}

// Close pool instance
func Close() {
	log := logger.Logger()

	if pool != nil {
		err := pool.Close()
		if err != nil {
			log.Error(err.Error())
		}
	}
}

func connectPool() {
	log := logger.Logger()

	p := redis.Pool{
		// Maximum number of idle connections in the pool
		MaxIdle: config.Get("redis.max_idle").(int),

		// Maximum number of connections allocated by the pool at a given timeutil.
		MaxActive: config.Get("redis.max_active").(int),

		// Close connections after remaining idle for this duration.
		IdleTimeout: time.Duration(config.Get("redis.idle_timeout").(int)) * time.Millisecond,

		// Dial is an application supplied function for creating and configuring a connection.
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial(
				"tcp",
				fmt.Sprintf("%s:%d", config.Get("redis.host"), config.Get("redis.port")),
				redis.DialDatabase(config.Get("redis.database").(int)),
				redis.DialPassword(config.Get("redis.auth").(string)),
			)

			if err != nil {
				log.Error(err.Error())
				return nil, err
			}
			return c, nil
		},

		// PING PONG test
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			if time.Since(t) < time.Minute {
				return nil
			}
			_, err := c.Do("PING")
			return err
		},
	}

	pool = &p
}
