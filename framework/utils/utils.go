/**
 ******************************************************************************
 * @file    utils.go
 * @author  GEEKROS,  site:www.geekros.com
 ******************************************************************************
 */

package Utils

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

type EmptyData struct {
}

func Success(c *gin.Context, code int, msg string, data interface{}) {

	if msg == "" {
		msg = "success"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})

	c.Set("code", code)
}

func Error(c *gin.Context, code int, msg string, data interface{}) {

	if msg == "" {
		msg = "error"
	}

	c.JSON(http.StatusOK, gin.H{
		"code": code,
		"msg":  msg,
		"data": data,
	})
}
