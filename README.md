## Go语言实现的命令行解析程序

包名：cli

### 编译打包

使用`git clone`下载代码后，将当前路径，比如 `F:\<project_name> `添加到`GOPATH`环境变量中，然后在<project_name>目录中打开一个控制台，输入：
```
go build main
```
就可以在当前目录下看到 `main.exe`

测试：
```
> main -h
帮助信息
> main -v
版本信息 0.0.1
> main -f file1 file2
目标文件 [file1 file2]
```

### 使用说明：

#### 初始化
```
cmd := cli.Cli{Options:make([]cli.Option, 3), CmdMap:make(map[string]bool)}
```
`Option`是`cli`包中的结构体，一个`Option`就是一个命令，`3`指定的是命令项数目（我的示例中有3个命令）

```
type Option struct {
	Cmd string // 命令项-简写
	Command string // 命令项-全称
	IsHasArgs bool // 是否有参数
	Description string // 命令描述
}
```

#### 添加不带参数的命令

```
cmd.Option(cli.Option{"h","help",false, "帮助信息"})
```

`Option`方法添加命令，指定其四个参数，第一、二个参数是指定命令名称，第三个参数是是否带参数，第四个是命令描述

#### 添加带参数的命令
```
cmd.Option(cli.Option{"f","file",true, "目标文件"})  // 目标文件
```

#### 命令解析
```
var r = cmd.Parse(args)
```

调用Parse方法进行命令解析，返回一个`map[string]interface{}`对象

调用命令：
```
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
```





