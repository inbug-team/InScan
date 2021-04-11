/*
负责人员：InBug Team
创建时间：2021/3/4
程序用途：mysql、pgsql、mssql检测模块
*/
package po

import (
	"fmt"
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mssql"
	_ "github.com/jinzhu/gorm/dialects/mysql"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func CheckSQL(hostType, ip, user, pwd string, port int) bool {

	connectStr := ""

	switch hostType {
	case "mysql":
		connectStr = fmt.Sprintf(
			"%s:%s@tcp(%s:%d)/%s?timeout=%ds",
			user, pwd, ip, port, "", 6,
		)
	case "postgres":
		connectStr = fmt.Sprintf(
			"host=%s port=%d user=%s dbname=%s sslmode=disable password=%s timeout=%ds",
			ip, port, user, "", pwd, 6,
		)
	case "mssql":
		connectStr = fmt.Sprintf(
			"server=%s;user id=%s;password=%s;port=%d;database=%s;timeout=%ds",
			ip, user, pwd, port, "", 6,
		)
	}

	db, err := gorm.Open(hostType, connectStr)
	if err != nil {
		return false
	}

	db.Close()
	return true
}
