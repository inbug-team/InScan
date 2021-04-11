/*
负责人员：张金龙
创建时间：2021/3/31
程序用途：生成器测试
*/
package main

import "github.com/inbug-team/InScan/generate_exec"

func main() {
	generate_exec.AppendExec(
		map[string]interface{}{"name": "张三", "age": 18, "sex": "男"},
		"./tpl",
		"./new",
	)
}
