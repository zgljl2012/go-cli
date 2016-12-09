package cli

/*
 * Go语言命令行解析： 
 * 解析 -<command简写> 与 --<command全称> 命令
 */

import (
    "strings"
)

/*
 * 单个命令项
 */
type Option struct {
	Cmd string // 命令项-简写
	Command string // 命令项-全称
	IsHasArgs bool // 是否有参数
	Description string // 命令描述
}

/*
 * 命令行解析类
 */
type Cli struct {
	Options []Option // 最多支持100个参数
	Pos int // 参数添加的当前位置
	CmdMap map[string]bool
}

/*
 * 增加命令项
 */
func (c *Cli) Option(o Option) {
	c.Options[c.Pos] = o
	c.Pos++
	c.CmdMap[o.Cmd] = o.IsHasArgs
	c.CmdMap[o.Command] = o.IsHasArgs
}

/*
 * 解析命令
 * 传入的参数是去掉了程序名称的，如命令行调用为： main -h hello
 * 传入的参数为：-h hello
 */
func (c *Cli) Parse(args []string) map[string]interface{} {
	r:= make(map[string]interface{})
	var isValue bool = false
	var curKey string
	for _, s := range args {
		if len(s) > 0 {
			if strings.HasPrefix(s, "-") || strings.HasPrefix(s, "--") {
				pos := 1
				if strings.HasPrefix(s, "--") {
					pos++
				}
				s=s[pos:]
				if c.CmdMap[s] { // 带参数
					isValue = true
					curKey = s
					r[curKey] = make([]string,0)  // 带参数返回参数数组
				} else { // 不带参数
					isValue = false
					r[s] = true // 不带参数返回一个true
				}
			} else {
				if isValue {
					r[curKey] = append(r[curKey].([]string),s)
				}
			}
		}
	}
	// 简写和全称同步
	for _,o := range c.Options {
		cmd,ok1 := r[o.Cmd]
		command,ok2 := r[o.Command]
		if ok1 && !ok2 {
			r[o.Command] = cmd
		}
		if !ok1 && ok2 {
			r[o.Cmd] = command
		}
	}
	return r
}
