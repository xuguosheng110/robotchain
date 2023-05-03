/**
 ******************************************************************************
 * @file    chat.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package OpenAIChat

import (
	"RobotChain/framework/model/openai"
	"RobotChain/framework/model/openai/model"
	"context"
	"errors"
)

type Client struct {
	openai *OpenAI.BaseOpenAI
	model  OpenAIChatModel.ModelConfig
}

type RequestParameter struct {
	Model            string         `json:"model,omitempty"`
	Messages         []*ChatMessage `json:"messages,omitempty"`
	Stop             []string       `json:"stop,omitempty"`
	Stream           bool           `json:"stream,omitempty"`
	N                int            `json:"n,omitempty"`
	TopP             float64        `json:"top_p,omitempty"`
	Temperature      float64        `json:"temperature,omitempty"`
	MaxTokens        int            `json:"max_tokens,omitempty"`
	PresencePenalty  float64        `json:"presence_penalty,omitempty"`
	FrequencyPenalty float64        `json:"frequency_penalty,omitempty"`
	User             string         `json:"user,omitempty"`
}

type ResponseParameter struct {
	ID        string                     `json:"id,omitempty"`
	Object    string                     `json:"object,omitempty"`
	CreatedAt int64                      `json:"created_at,omitempty"`
	Choices   []*responseParameterChoice `json:"choices,omitempty"`
	Usage     *responseParameterUsage    `json:"usage,omitempty"`
}

type responseParameterChoice struct {
	Message      *ChatMessage `json:"message,omitempty"`
	Index        int          `json:"index,omitempty"`
	LogProbs     int          `json:"logprobs,omitempty"`
	FinishReason string       `json:"finish_reason,omitempty"`
}

type responseParameterUsage struct {
	PromptTokens     int `json:"prompt_tokens,omitempty"`
	CompletionTokens int `json:"completion_tokens,omitempty"`
	TotalTokens      int `json:"total_tokens,omitempty"`
}

type ChatMessage struct {
	Role    string `json:"role,omitempty"`
	Content string `json:"content,omitempty"`
	Name    string `json:"name,omitempty"`
}

func NewChat(openai *OpenAI.BaseOpenAI, model OpenAIChatModel.ModelConfig) *Client {
	return &Client{
		openai: openai,
		model:  model,
	}
}

func (client *Client) ChatCompletion(ctx context.Context, request *RequestParameter) (*ResponseParameter, error) {

	request.Model = client.model.Name

	if request.Stream {
		return nil, errors.New("use NewStream instead")
	}

	var response ResponseParameter
	if err := client.openai.OnRequest(ctx, client.model.Path, client.model.Method, &request, &response); err != nil {
		return nil, err
	}
	return &response, nil
}
