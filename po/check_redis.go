/*
负责人员：InBug Team
创建时间：2021/3/1
程序用途：redis检测模块
*/
package po

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"time"
)

func CheckRedis(ip, user, pwd string, port int) bool {
	client := redis.NewClient(&redis.Options{
		Addr:        fmt.Sprintf(`%s:%d`, ip, port),
		Username:    user,
		Password:    pwd,
		DB:          0,
		DialTimeout: 6 * time.Second,
	})

	defer client.Close()

	ctx := context.Background()
	status, err := client.Ping(ctx).Result()
	if err != nil {
		return false
	}

	if status == "PONG" {
		return true
	}

	return false
}
