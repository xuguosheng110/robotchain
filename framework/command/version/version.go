/**
 ******************************************************************************
 * @file    version.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package VersionCommand

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func Start(name string, version string) *cobra.Command {
	command := &cobra.Command{
		Use:     "version",
		Short:   "Get version number",
		Long:    "Command to get version number",
		Example: "geek version",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[version]：" + color.Green.Text(name+" "+version+" "))
		},
	}
	return command
}
