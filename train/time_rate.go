package train

import (
	"fmt"
	"github.com/go-redis/redis"
	"golang.org/x/time/rate"
	"time"
)

func TimeRate() {
	// 创建一个每秒限制qps次的速率限制器
	fmt.Println(rate.Every(time.Second), rate.Limit(2))
	limiter := rate.NewLimiter(rate.Limit(5), 5)

	// 使用Redis做分布式限流
	//redisClient := redis.NewClient(&redis.Options{
	//	Addr:     "localhost:6379",
	//	Password: "", // 没有密码时设置为空字符串
	//	DB:       0,  // 默认数据库
	//})

	// 定义一个使用Redis做分布式速率限制的函数
	distributedLimiter := NewDistributedLimiter(limiter)

	for i := 0; i < 10; i++ {
		go func(i int) {
			// 尝试执行被限制的操作
			if distributedLimiter.Allow() {
				fmt.Println("操作被允许执行: ", i)
			} else {
				fmt.Println("操作被限流了: ", i)
			}
		}(i)
	}

	time.Sleep(1 * time.Second) // 模拟操作耗时

	for i := 10; i < 20; i++ {
		go func(i int) {
			// 尝试执行被限制的操作
			if distributedLimiter.Allow() {
				fmt.Println("操作被允许执行: ", i)
			} else {
				fmt.Println("操作被限流了: ", i)
			}
		}(i)
	}

	time.Sleep(time.Second) // 模拟操作耗时
}

// DistributedLimiter 使用Redis的速率限制器
type DistributedLimiter struct {
	redisClient *redis.Client
	limiter     *rate.Limiter
}

// NewDistributedLimiter 创建一个新的分布式速率限制器
func NewDistributedLimiter(limiter *rate.Limiter) *DistributedLimiter {
	return &DistributedLimiter{
		limiter: limiter,
	}
}

// Allow 检查是否允许执行操作
func (l *DistributedLimiter) Allow() bool {
	if l.limiter.Allow() {
		return true
	}

	// 如果本地限制器不允许，尝试从Redis获取令牌
	//_, err := l.redisClient.Ping().Result()
	//if err != nil {
	//	fmt.Println("Redis连接失败:", err)
	//	return false
	//}
	//
	//// 尝试重新获取令牌
	//if l.redisClient.SetNX("lock", "1", 1*time.Second).Val() {
	//	// 获取到锁，重置令牌
	//	l.redisClient.Expire("lock", 1*time.Second)
	//	return true
	//}

	return false
}
