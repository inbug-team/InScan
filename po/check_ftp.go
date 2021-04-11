/*
负责人员：张金龙
创建时间：2021/3/4
程序用途：ftp检测模块
*/
package po

import (
	"fmt"
	"github.com/jlaffaye/ftp"
	"time"
)

func CheckFTP(ip, user, pwd string, port int) bool {
	client, err := ftp.Dial(fmt.Sprintf(`%s:%d`, ip, port), ftp.DialWithTimeout(6*time.Second))
	if err != nil {
		fmt.Println(err)
		return false
	}

	err = client.Login(user, pwd)
	if err != nil {
		fmt.Println(err)
		return false
	}

	defer client.Quit()

	return true
}
