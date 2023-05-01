/**
 ******************************************************************************
 * @file    config.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package Config

var Get = &config{}

type config struct {
	Name      string `json:"name"`
	Version   string `json:"version"`
	Workspace string `json:"workspace"`
}

func Init() {
	// 框架基础信息配置
	Get.Name = "RobotChain"
	Get.Version = "1.0.0"
	Get.Workspace = "/opt/workspace"
}
