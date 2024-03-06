package redis

import (
	"fmt"
	"github.com/redis/go-redis/v9"
)

type Config struct {
	Host     string `json:"host"`
	Port     int    `json:"port"`
	Password string `json:"password"`
	DB       int    `json:"db"`
}
type Adapter struct {
	client *redis.Client
}

func New(c Config) Adapter {
	rdb := redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", c.Host, c.Port),
		Password: c.Password, // no password set
		DB:       c.DB,       // use default DB

	})
	return Adapter{
		client: rdb,
	}
}

func (a Adapter) Client() *redis.Client {
	return a.client
}
