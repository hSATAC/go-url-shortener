package shortener

import (
	"github.com/garyburd/redigo/redis"
	"time"
)

// Ability to generate a uniqID for base62
type UniqIDGenerator interface {
	GetUniqID() int64
}

type Backend interface {
	UniqIDGenerator
}

const uniqIDStartsAt = 4000

type RedisBackend struct {
	Backend
	config RedisConfig
	pool   *redis.Pool
}

func (backend *RedisBackend) GetUniqID() int64 {
	return 4000
}

func NewRedisBackend(config RedisConfig) *RedisBackend {
	var pool = &redis.Pool{
		MaxIdle:     config.maxIdleConn,
		IdleTimeout: time.Duration(config.maxIdleTimeoutSec) * time.Second,
		Dial: func() (redis.Conn, error) {
			c, err := redis.Dial("tcp", config.address)
			if err != nil {
				return nil, err
			}
			return c, err
		},
		TestOnBorrow: func(c redis.Conn, t time.Time) error {
			_, err := c.Do("PING")
			return err
		},
	}

	return &RedisBackend{
		config: config,
		pool:   pool,
	}
}
