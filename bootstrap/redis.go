package bootstrap

import (
	"fmt"
	"gohub/pkg/config"
	"gohub/pkg/redis"
)

// SetupRedis 初始化 Redis
func SetupRedis() {

	// 建立 Redis 连接
	redis.ConnectRedis(
		fmt.Sprintf("%v:%v", config.GetStirng("redis.host"), config.GetStirng("redis.port")),
		config.GetStirng("redis.username"),
		config.GetStirng("redis.password"),
		config.GetInt("redis.database"),
	)
}