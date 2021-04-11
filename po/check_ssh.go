/*
负责人员：张金龙
创建时间：2021/3/1
程序用途：ssh检测模块
*/
package po

import (
	"fmt"
	"golang.org/x/crypto/ssh"
	"net"
	"time"
)

func CheckSSH(ip, user, pwd string, port int) bool {
	result := false
	config := &ssh.ClientConfig{
		User: user,
		Auth: []ssh.AuthMethod{ssh.Password(pwd)},
		HostKeyCallback: func(hostname string, remote net.Addr, key ssh.PublicKey) error {
			return nil
		},
		Timeout: 6 * time.Second,
	}

	client, err := ssh.Dial("tcp", fmt.Sprintf(`%s:%d`, ip, port), config)
	if err == nil {
		defer client.Close()
		session, err := client.NewSession()
		if err == nil {
			errEcho := session.Run("echo inBug")
			if errEcho == nil {
				defer session.Close()
				result = true
			}
		}

	}
	return result
}
