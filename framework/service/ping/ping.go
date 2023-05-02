/**
 ******************************************************************************
 * @file    ping.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package PingService

import (
	"RobotChain/framework/config"
	"RobotChain/framework/utils"
	"github.com/gin-gonic/gin"
)

type requestPing struct {
	Name    string `json:"name"`
	Version string `json:"version"`
}

func Ping(c *gin.Context) {

	returnData := requestPing{}

	returnData.Name = Config.Get.Name
	returnData.Version = Config.Get.Version

	Utils.Success(c, 0, "", returnData)
}
