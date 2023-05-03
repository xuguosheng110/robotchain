/**
 ******************************************************************************
 * @file    chat.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ChatCommand

import (
	"RobotChain/framework/model/openai"
	"RobotChain/framework/model/openai/chat"
	"RobotChain/framework/model/openai/model"
	"context"
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
					ctx := context.Background()
					openai := OpenAI.NewBaseOpenAI("", "")
					client := OpenAIChat.NewChat(openai, OpenAIChatModel.GPT3p5Turbo)
					response, err := client.ChatCompletion(ctx, &OpenAIChat.RequestParameter{
						Messages: []*OpenAIChat.ChatMessage{
							{Role: "user", Content: args[1]},
						},
						Temperature: 0.6,
					})
					if err != nil {
						fmt.Println("[chat]：" + color.Red.Text(err.Error()))
						return
					}
					for _, choice := range response.Choices {
						msg := choice.Message
						fmt.Println("[chat]：" + color.Info.Sprintf("role=%q, content=%q", msg.Role, msg.Content))
						return
					}
				} else {
					fmt.Println("[chat]：" + color.Yellow.Text("Currently only the OpenAI platform is supported"))
					return
				}
			}
		},
	}
	return command
}
