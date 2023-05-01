/**
 ******************************************************************************
 * @file    main.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package main

import (
	"RobotChain/framework/command"
	"RobotChain/framework/config"
)

func main() {

	// 初始化核心配置
	Config.Init()

	// 初始化命令行工具
	Command.Init()
}
