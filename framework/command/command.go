/**
 ******************************************************************************
 * @file    command.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package Command

import (
	ModelCommand "RobotChain/framework/command/model"
	"RobotChain/framework/command/service"
	"RobotChain/framework/command/version"
	"RobotChain/framework/config"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"os"
)

func Init() {

	// 实例化命令行工具
	command := &cobra.Command{
		Use:   "geek",
		Short: "Welcome to " + Config.Get.Name,
		Long:  Config.Get.Name + " Version " + color.Green.Text(Config.Get.Version) + "\nDevelopment Team：GEEKROS https://www.geekros.com",
	}

	// 隐藏内置Completion命令行工具
	command.CompletionOptions.HiddenDefaultCmd = true

	command.AddCommand(VersionCommand.Start(Config.Get.Name, Config.Get.Version))

	command.AddCommand(ServiceCommand.Start())

	command.AddCommand(ModelCommand.Models())

	if err := command.Execute(); err != nil {
		os.Exit(1)
	}
}
