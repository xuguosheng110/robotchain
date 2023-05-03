/**
 ******************************************************************************
 * @file    stream.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package OpenAIChat

import (
	"RobotChain/framework/model/openai"
	"RobotChain/framework/model/openai/model"
	"context"
)

type StreamClient struct {
	openai *OpenAI.BaseOpenAI
	model  OpenAIChatModel.ModelConfig
}

type StreamResponseParameter struct {
	ID        string             `json:"id,omitempty"`
	Object    string             `json:"object,omitempty"`
	CreatedAt int64              `json:"created_at,omitempty"`
	Choices   []*StreamingChoice `json:"choices,omitempty"`
}

type StreamingChoice struct {
	Delta        *ChatMessage `json:"delta,omitempty"`
	Index        int          `json:"index,omitempty"`
	LogProbs     int          `json:"logprobs,omitempty"`
	FinishReason string       `json:"finish_reason,omitempty"`
}

func NewStream(openai *OpenAI.BaseOpenAI, model OpenAIChatModel.ModelConfig) *StreamClient {
	return &StreamClient{
		openai: openai,
		model:  model,
	}
}

func (client *Client) StreamCompletion(ctx context.Context, request *RequestParameter, function func(response *StreamResponseParameter)) error {

	request.Model = client.model.Name
	request.Stream = true

	var response StreamResponseParameter
	return client.openai.OnStream(ctx, client.model.Path, client.model.Method, &request, &response, func(response any) {
		function(response.(*StreamResponseParameter))
	})
}
