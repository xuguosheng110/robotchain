/**
 ******************************************************************************
 * @file    openai.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package OpenAI

import (
	"encoding/json"
	"errors"
	"github.com/go-resty/resty/v2"
	"os"
)

const BaseURL = "https://chat.chatdesktop.com"

type ErrorReturn struct {
	Error ErrorMessage `json:"error"`
}

type ErrorMessage struct {
	Message string `json:"message"`
}

type BaseResty struct {
	client *resty.Client
}

type BaseParameter struct {
	ApiKey string `json:"api_key"`
	BaseResty
}

type BaseOption func(*BaseResty)

type OpenAI[Request any, Response any] struct {
	BaseParameter
	Method string `json:"method"`
	Path   string `json:"path"`
}

func NewBaseOpenAI(apiKey string, options ...BaseOption) (BaseParameter, ErrorReturn) {
	if apiKey == "" {
		apiKey = os.Getenv("OPENAI_API_KEY")
	}
	if apiKey == "" {
		return BaseParameter{}, ErrorReturn{Error: ErrorMessage{Message: "OPENAI_API_KEY Not Set"}}
	}

	base := BaseParameter{
		ApiKey: apiKey,
	}

	for _, option := range options {
		option(&base.BaseResty)
	}

	return base, ErrorReturn{}
}

func (o OpenAI[Request, Response]) Request(parameter ...Request) (Response, ErrorReturn) {
	var returnResponse Response
	requestBytes, err := json.Marshal(parameter)
	if err != nil {
		return returnResponse, ErrorReturn{Error: ErrorMessage{Message: "response error"}}
	}

	if o.BaseResty.client == nil {
		o.BaseResty.client = resty.New()
		o.BaseResty.client.SetBaseURL(BaseURL)
	}

	query := o.BaseResty.client.R()
	query.SetHeader("Authorization", "Bearer "+o.ApiKey)
	query.SetHeader("Content-Type", "application/json")
	query.SetBody(string(requestBytes))

	var response *resty.Response

	switch o.Method {
	case "GET":
		response, err = query.Get(o.Path)
	default:
		err = errors.New("response error")
	}

	if err != nil {
		return returnResponse, ErrorReturn{Error: ErrorMessage{Message: "response error"}}
	}

	if response.StatusCode() != 200 {
		var errorReturn ErrorReturn
		_ = json.Unmarshal(response.Body(), &errorReturn)
		if err != nil {
			return returnResponse, errorReturn
		}
	}

	err = json.Unmarshal(response.Body(), &returnResponse)
	if err != nil {
		return returnResponse, ErrorReturn{Error: ErrorMessage{Message: "response error"}}
	}

	return returnResponse, ErrorReturn{}
}
