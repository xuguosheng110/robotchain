/**
 ******************************************************************************
 * @file    model.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ModelService

import (
	"RobotChain/framework/model/openai"
	"RobotChain/framework/utils"
	"github.com/gin-gonic/gin"
)

type requestList struct {
	OpenAI.ResponseGetModels
}

func List(c *gin.Context) {

	returnData := requestList{}

	Utils.Success(c, 0, "", returnData)
	return
}

type requestGet struct {
	OpenAI.ResponseGetModel
}

func Get(c *gin.Context) {

	returnData := requestGet{}

	Utils.Success(c, 0, "", returnData)
	return
}
