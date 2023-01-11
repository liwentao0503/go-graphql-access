package redis

import (
	"fmt"

	"github.com/go-redis/redis"
)

var client *redis.Client

func NewClient() *redis.Client {
	return client
}

func InitDB() error {
	// 创建 redis 客户端
	client = redis.NewClient(&redis.Options{
		Addr:     "127.0.0.1:6379",
		Password: "", // no password set
		DB:       0,  // use default DB
		PoolSize: 10, // 连接池的大小为 10
	})
	//defer client.Close()
	// 通过 cient.Ping() 来检查是否成功连接到了 redis 服务器
	pong, err := client.Ping().Result()
	fmt.Println("Redis Client: " + pong)
	return err
}
