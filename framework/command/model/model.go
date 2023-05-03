/**
 ******************************************************************************
 * @file    model.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ModelCommand

import (
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func Models() *cobra.Command {
	command := &cobra.Command{
		Use:     "model",
		Short:   "Retrieving the specified platform's model data",
		Long:    "Command to retrieving the specified platform's model data",
		Example: "geek model [openai]\ngeek model [openai] [model_id]",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[model]：" + color.Yellow.Text("I'm not sure which platform's model data you want to retrieve."))
				return
			}
		},
	}
	return command
}
