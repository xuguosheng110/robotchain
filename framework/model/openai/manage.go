/**
 ******************************************************************************
 * @file    manage.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package OpenAI

type RequestGetModel struct {
}

type ResponseGetModels struct {
	Error ErrorMessage `json:"error,omitempty"`
	Data  []struct {
		Id         string        `json:"id"`
		Object     string        `json:"object"`
		OwnedBy    string        `json:"owned_by"`
		Permission []interface{} `json:"permission"`
	} `json:"data,omitempty"`
	Object string `json:"object,omitempty"`
}

type ResponseGetModel struct {
	Error      ErrorMessage  `json:"error,omitempty"`
	Id         string        `json:"id"`
	Object     string        `json:"object"`
	OwnedBy    string        `json:"owned_by"`
	Permission []interface{} `json:"permission"`
}

func GetModels(baseParameter BaseParameter) OpenAI[RequestGetModel, ResponseGetModels] {
	return OpenAI[RequestGetModel, ResponseGetModels]{
		BaseParameter: baseParameter,
		Path:          "/v1/models",
		Method:        "GET",
	}
}

func GetModel(baseParameter BaseParameter, modelID string) OpenAI[RequestGetModel, ResponseGetModel] {
	return OpenAI[RequestGetModel, ResponseGetModel]{
		BaseParameter: baseParameter,
		Path:          "/v1/models/" + modelID,
		Method:        "GET",
	}
}
