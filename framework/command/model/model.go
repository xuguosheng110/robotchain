/**
 ******************************************************************************
 * @file    model.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ModelCommand

import (
	"RobotChain/framework/model/openai"
	"encoding/json"
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
			fmt.Println("[model]：" + color.Gray.Text("Searching for the relevant model data for you..."))
			if len(args) == 1 {
				if args[0] == "openai" {
					base := OpenAI.NewBaseOpenAI("")
				} else {
					fmt.Println("[model]：" + color.Yellow.Text("Currently only the OpenAI platform is supported"))
					return
				}
			}
			if len(args) == 2 {
				if args[0] == "openai" {
					base, err := OpenAI.NewBaseOpenAI()
					if err.Error.Message != "" {
						fmt.Println("[model]：" + color.Red.Text(err.Error.Message))
						return
					}
					request := OpenAI.GetModel(base, args[1]).OnRequest()
					if request.Error.Message != "" {
						fmt.Println("[model]：" + color.Red.Text(request.Error.Message))
						return
					}
					requestString, _ := json.Marshal(request)
					fmt.Println("[model]：" + color.Info.Sprintf(string(requestString)))
					return
				} else {
					fmt.Println("[model]：" + color.Yellow.Text("Currently only the OpenAI platform is supported"))
					return
				}
			}
		},
	}
	return command
}
