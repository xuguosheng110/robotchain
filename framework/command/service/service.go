/**
 ******************************************************************************
 * @file    service.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ServiceCommand

import (
	"RobotChain/framework/service"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
	"golang.org/x/sync/errgroup"
	"net/http"
	"time"
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
			start := &http.Server{
				Addr:           fmt.Sprintf(":%d", 10081),
				Handler:        Service.Router(),
				ReadTimeout:    60 * time.Second,
				WriteTimeout:   60 * time.Second,
				MaxHeaderBytes: 1 << 20,
			}

			ServiceGroup.Go(func() error {
				return start.ListenAndServe()
			})

			fmt.Println("[service]：" + color.Info.Sprintf("Core service started successfully"))

			if err := ServiceGroup.Wait(); err != nil {
			}
		},
	}
	return command
}
