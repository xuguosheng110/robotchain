/**
 ******************************************************************************
 * @file    model.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package OpenAIChatModel

var (
	Gpt4                 = ModelConfig{Name: "gpt-4"}
	Gpt40314             = ModelConfig{Name: "gpt-4-0314"}
	Gpt432k              = ModelConfig{Name: "gpt-4-32k"}
	Gpt432k0314          = ModelConfig{Name: "gpt-4-32k-0314"}
	GPT3p5Turbo          = ModelConfig{Name: "gpt-3.5-turbo", Path: "/v1/chat/completions", Method: "POST"}
	GPT3p5Turbo0301      = ModelConfig{Name: "gpt-3.5-turbo-0301"}
	TextDavinci003       = ModelConfig{Name: "text-davinci-003"}
	TextDavinci002       = ModelConfig{Name: "text-davinci-002"}
	CodeDavinci002       = ModelConfig{Name: "code-davinci-002"}
	TextCurie001         = ModelConfig{Name: "text-curie-001"}
	TextBabbage001       = ModelConfig{Name: "text-babbage-001"}
	TextAda001           = ModelConfig{Name: "text-ada-001"}
	Davinci              = ModelConfig{Name: "davinci"}
	Curie                = ModelConfig{Name: "curie"}
	Babbage              = ModelConfig{Name: "babbage"}
	Ada                  = ModelConfig{Name: "ada"}
	CodeDavinci001       = ModelConfig{Name: "code-davinci-001"}
	CodeCushman002       = ModelConfig{Name: "code-cushman-002"}
	CodeCushman001       = ModelConfig{Name: "code-cushman-001"}
	Whisper              = ModelConfig{Name: "whisper-1"}
	TextEmbeddingAda002  = ModelConfig{Name: "text-embedding-ada-002"}
	TextModerationLatest = ModelConfig{Name: "text-moderation-latest"}
	TextModerationStable = ModelConfig{Name: "text-moderation-stable"}
)

type ModelConfig struct {
	Name   string `json:"name"`
	Path   string `json:"path"`
	Method string `json:"method"`
}

func (model ModelConfig) GetModel() ModelConfig {
	return model
}
