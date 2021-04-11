/*
负责人员：InBug Team
创建时间：2021/3/4
程序用途：snmp检测模块
*/
package po

import (
	"github.com/gosnmp/gosnmp"
	"time"
)

func CheckSNMP(ip, user, pwd string, port int) bool {
	flag := false

	gosnmp.Default.Target = ip
	gosnmp.Default.Port = uint16(port)
	gosnmp.Default.Community = pwd
	gosnmp.Default.Timeout = 4 * time.Second

	err := gosnmp.Default.Connect()
	if err == nil {
		oidList := []string{"1.3.6.1.2.1.1.4.0", "1.3.6.1.2.1.1.7.0"}
		_, err := gosnmp.Default.Get(oidList)
		if err == nil {
			flag = true
		}
	}

	return flag
}
