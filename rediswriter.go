package klog

import (
	_ "fmt"
	"github.com/vmihailenco/redis/v2"
	"strings"
	"time"
)

type RedisBackend struct {
	Addr     string
	Password string
	DB       int64
	client   *redis.Client
}

func GenResidBackend(addr string, passwd string, db int64) (b *RedisBackend, err error) {
	b = &RedisBackend{
		Addr:     addr,
		Password: passwd,
		DB:       db,
	}

	b.client = redis.NewTCPClient(&redis.Options{
		Addr:     b.Addr,
		Password: b.Password,
		DB:       b.DB,
	})

	ping := b.client.Ping()
	if ping.Err() != nil {
		return b, ping.Err()
	}
	return

}

func (r *RedisBackend) Write(p []byte) (n int, err error) {
	write_str := string(p)
	write_item := strings.Split(write_str, "\t")
	level, body := write_item[0], write_item[1]
	now := time.Now().Unix()
	zadd := r.client.ZAdd(level, redis.Z{float64(now), body})
	return n, zadd.Err()

}
