/**
 ******************************************************************************
 * @file    openai.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package OpenAI

import (
	"RobotChain/framework/config"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"time"
)

const BaseURL = "https://chat.chatdesktop.com"

type BaseOpenAI struct {
	apiKey     string
	baseURL    string
	HTTPClient *http.Client
	userAgent  string
}

type BaseOpenAIError struct {
	StatusCode int
	Payload    []byte
}

func (e *BaseOpenAIError) Error() string {
	return fmt.Sprintf("status_code=%d, payload=%s", e.StatusCode, e.Payload)
}

func NewBaseOpenAI(apiKey string, baseURL string) *BaseOpenAI {
	return &BaseOpenAI{
		apiKey:  apiKey,
		baseURL: baseURL,
		HTTPClient: &http.Client{
			Timeout: 30 * time.Second,
		},
		userAgent: "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/112.0.0.0 Safari/537.36 RobotChain/" + Config.Get.Version,
	}
}

func (openai *BaseOpenAI) OnRequest(ctx context.Context, path string, parameter any, response any) error {
	requestBody, err := json.Marshal(parameter)
	if err != nil {
		return err
	}

	if openai.baseURL == "" {
		openai.baseURL = BaseURL
	}

	request, err := http.NewRequestWithContext(ctx, http.MethodPost, openai.baseURL+path, bytes.NewReader(requestBody))
	if err != nil {
		return err
	}

	responseBody, err := openai.doRequest(request, "application/json")
	if err != nil {
		return err
	}
	defer responseBody.Close()

	return json.NewDecoder(responseBody).Decode(response)
}

func (openai *BaseOpenAI) doRequest(request *http.Request, contentType string) (io.ReadCloser, error) {
	if openai.apiKey == "" {
		openai.apiKey = Config.Get.OpenAI.ApiKey
	}
	if openai.apiKey != "" {
		request.Header.Set("Authorization", "Bearer "+openai.apiKey)
	}
	request.Header.Set("Content-Type", contentType)
	request.Header.Add("User-Agent", openai.userAgent)

	response, err := openai.HTTPClient.Do(request)
	if err != nil {
		return nil, fmt.Errorf("error request: %w", err)
	}

	if response.StatusCode < 200 || response.StatusCode >= 400 {
		respBody, err := io.ReadAll(response.Body)
		if err != nil {
			return nil, err
		}
		return nil, &BaseOpenAIError{
			StatusCode: response.StatusCode,
			Payload:    respBody,
		}
	}

	return response.Body, nil
}
