package main

import (
	"fmt"
	"cli"
	"os"
)

func main() {
	args := os.Args
	cmd := cli.Cli{Options:make([]cli.Option, 3), CmdMap:make(map[string]bool)}
	cmd.Option(cli.Option{"h","help",false, "帮助信息"}) // 帮助信息
	cmd.Option(cli.Option{"f","file",true, "目标文件"})  // 目标文件
	cmd.Option(cli.Option{"v","version",false, "版本信息"})
	args = args[1:]
	// 命令解析
	var r = cmd.Parse(args)
	_,ok1 := r["help"]
	file,ok2 := r["file"]
	_,ok3 := r["version"]

	if ok1 {
		fmt.Println("帮助信息")
	}

	if ok2 {
		fmt.Println("目标文件", file)
	}

	if ok3 {
		fmt.Println("版本：0.0.1")
	}
}
