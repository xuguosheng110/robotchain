/**
 ******************************************************************************
 * @file    model.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package model

var (
	Gpt4                 = ModelName{Name: "gpt-4"}
	Gpt40314             = ModelName{Name: "gpt-4-0314"}
	Gpt432k              = ModelName{Name: "gpt-4-32k"}
	Gpt432k0314          = ModelName{Name: "gpt-4-32k-0314"}
	GPT3p5Turbo          = ModelName{Name: "gpt-3.5-turbo"}
	GPT3p5Turbo0301      = ModelName{Name: "gpt-3.5-turbo-0301"}
	TextDavinci003       = ModelName{Name: "text-davinci-003"}
	TextDavinci002       = ModelName{Name: "text-davinci-002"}
	CodeDavinci002       = ModelName{Name: "code-davinci-002"}
	TextCurie001         = ModelName{Name: "text-curie-001"}
	TextBabbage001       = ModelName{Name: "text-babbage-001"}
	TextAda001           = ModelName{Name: "text-ada-001"}
	Davinci              = ModelName{Name: "davinci"}
	Curie                = ModelName{Name: "curie"}
	Babbage              = ModelName{Name: "babbage"}
	Ada                  = ModelName{Name: "ada"}
	CodeDavinci001       = ModelName{Name: "code-davinci-001"}
	CodeCushman002       = ModelName{Name: "code-cushman-002"}
	CodeCushman001       = ModelName{Name: "code-cushman-001"}
	Whisper              = ModelName{Name: "whisper-1"}
	TextEmbeddingAda002  = ModelName{Name: "text-embedding-ada-002"}
	TextModerationLatest = ModelName{Name: "text-moderation-latest"}
	TextModerationStable = ModelName{Name: "text-moderation-stable"}
)

type ModelName struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (model ModelName) ModelName() string {
	return model.Name
}
