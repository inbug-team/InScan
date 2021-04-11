/*
负责人员：张金龙
创建时间：2021/3/1
程序用途：mongodb检测模块
*/
package po

import (
	"context"
	"fmt"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
)

func CheckMongoDB(ip, user, pwd string, port int) bool {
	// 设置客户端连接配置
	clientOptions := options.Client().ApplyURI(
		fmt.Sprintf(`mongodb://%s:%s@%s:%d`, user, pwd, ip, port),
	)

	// 连接到MongoDB
	client, err := mongo.Connect(context.TODO(), clientOptions)
	if err != nil {
		return false
	}

	// 检查连接
	err = client.Ping(context.TODO(), nil)
	if err != nil {
		return false
	}

	return true
}
