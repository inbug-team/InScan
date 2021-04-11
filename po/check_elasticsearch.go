/*
负责人员：InBug Team
创建时间：2021/3/4
程序用途：es检测模块
*/
package po

import (
	"context"
	"fmt"
	"github.com/olivere/elastic/v7"
)

func CheckElasticSearch(ip, user, pwd string, port int) bool {
	flag := false

	client, err := elastic.NewClient(elastic.SetURL(fmt.Sprintf("http://%v:%v", ip, port)),
		elastic.SetBasicAuth(user, pwd),
	)
	if err == nil {
		ctx := context.Background()
		_, _, err = client.Ping(fmt.Sprintf("http://%v:%v", ip, port)).Do(ctx)
		if err == nil {
			flag = true
		}
	}
	return flag
}
