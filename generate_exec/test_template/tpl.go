/*
负责人员：InBug Team
创建时间：2021/3/31
程序用途：生成器模板
*/
package main

import (
	"fmt"
	"github.com/inbug-team/InScan/generate_exec"
	"os"
)

func main() {
	path := generate_exec.GetExecPath()
	//path := "./new"
	selfFile, err := os.OpenFile(path, os.O_RDONLY, 0666)
	generate_exec.CheckError(err)
	defer selfFile.Close()

	selfInfo, err := selfFile.Stat()
	generate_exec.CheckError(err)

	size := selfInfo.Size()
	selfBuf := make([]byte, size)
	selfFile.Read(selfBuf)

	fileLength := generate_exec.BytesToInt(selfBuf[size-4:])
	paramsByte := selfBuf[fileLength : size-4]

	fmt.Println(generate_exec.AesDecrypt(string(paramsByte), "1234567890123456"))
}
