/**
 ******************************************************************************
 * @file    model.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package ModelService

import (
	"RobotChain/framework/utils"
	"github.com/gin-gonic/gin"
)

type requestList struct {
}

func List(c *gin.Context) {

	returnData := requestList{}

	Utils.Success(c, 0, "", returnData)
	return
}

type requestGet struct {
}

func Get(c *gin.Context) {

	returnData := requestGet{}

	Utils.Success(c, 0, "", returnData)
	return
}
