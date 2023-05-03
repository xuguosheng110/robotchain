/**
 ******************************************************************************
 * @file    chat.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ChatCommand

import (
	"RobotChain/framework/model/openai"
	"fmt"
	"github.com/gookit/color"
	"github.com/spf13/cobra"
)

func Chat() *cobra.Command {
	command := &cobra.Command{
		Use:     "chat",
		Short:   "Sending session message to the specified platform",
		Long:    "Command to sending session message to the specified platform",
		Example: "geek chat [openai] [message]",
		Run: func(cmd *cobra.Command, args []string) {
			if len(args) == 0 {
				fmt.Println("[chat]：" + color.Yellow.Text("I'm not sure which platform you want to use for the conversation."))
				return
			}
			fmt.Println("[chat]：" + color.Gray.Text("Message sent successfully, waiting for platform feedback..."))
			if len(args) == 2 {
				if args[0] == "openai" {
					base, err := OpenAI.NewBaseOpenAI()
					if err.Error.Message != "" {
						fmt.Println("[chat]：" + color.Red.Text(err.Error.Message))
						return
					}

					request := OpenAI.NewChat(base).OnStream(OpenAI.RequestChat{
						Model: OpenAI.GPT3p5Turbo.ModelName(),
						Messages: []OpenAI.RequestChatMessage{{
							Role:    "user",
							Content: args[1],
						}},
						Temperature: 0,
						Stream:      true,
					})
					data := request.Recv()
					fmt.Println(request, data)
				} else {
					fmt.Println("[chat]：" + color.Yellow.Text("Currently only the OpenAI platform is supported"))
					return
				}
			}
		},
	}
	return command
}
