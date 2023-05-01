/**
 ******************************************************************************
 * @file    service.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ServiceCommand

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
)

var (
	ServiceGroup errgroup.Group
)

func Start() *cobra.Command {
	command := &cobra.Command{
		Use:     "service",
		Short:   "Start ore service",
		Long:    "Command to start core service",
		Example: "geek service",
		Run: func(cmd *cobra.Command, args []string) {
			fmt.Println("[service]：" + color.Gray.Text("Core service is starting..."))
		},
	}
	return command
}
