/*
负责人员：InBug Team
创建时间：2021/3/4
程序用途：smb检测模块
*/
package po

import (
	"github.com/stacktitan/smb/smb"
)

func CheckSMB(ip, user, pwd string, port int) bool {
	flag := false
	options := smb.Options{
		Host:        ip,
		Port:        port,
		User:        user,
		Password:    pwd,
		Domain:      "",
		Workstation: "",
	}

	session, err := smb.NewSession(options, false)
	if err == nil {
		session.Close()
		if session.IsAuthenticated {
			flag = true
		}
	}
	return flag
}
