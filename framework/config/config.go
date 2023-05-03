/**
 ******************************************************************************
 * @file    config.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package Config

var Get = &config{}

type config struct {
	Name      string       `json:"name"`
	Version   string       `json:"version"`
	Workspace string       `json:"workspace"`
	OpenAI    configOpenAi `json:"open_ai"`
}

type configOpenAi struct {
	ApiKey string `json:"api_key"`
}

func Init() {
	// 框架基础信息配置
	Get.Name = "RobotChain"
	Get.Version = "1.0.0"
	Get.Workspace = "/opt/workspace"

	// OpenAI配置
	Get.OpenAI.ApiKey = "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"
}
