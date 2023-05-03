/**
 ******************************************************************************
 * @file    service.go
 * @author  GEEKROS site:www.geekros.com github:geekros.github.io
 ******************************************************************************
 */

package Service

import (
	"RobotChain/framework/service/model"
	"RobotChain/framework/service/ping"
	"github.com/gin-gonic/gin"
	"net/http"
)

func Router() http.Handler {

	router := gin.New()

	gin.SetMode(gin.ReleaseMode)

	router.Use(gin.Recovery())

	router.GET("/", PingService.Ping)

	model := router.Group("model")
	{
		model.GET("/list", ModelService.List)

		model.GET("/get", ModelService.Get)
	}

	return router
}
