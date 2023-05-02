/**
 ******************************************************************************
 * @file    utils.go
 * @author  GEEKROS,  site:www.geekros.com
 ******************************************************************************
 */

package Utils

import (
	"github.com/gin-gonic/gin"
	"net"
	"net/http"
	"os"
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

func GetIp() string {
	Ip := "0.0.0.0"
	addRs, err := net.InterfaceAddrs()
	if err != nil {
		os.Exit(1)
	}
	for _, address := range addRs {
		if ipNet, ok := address.(*net.IPNet); ok && !ipNet.IP.IsLoopback() {
			if ipNet.IP.To4() != nil {
				Ip = ipNet.IP.String()
			}
		}
	}
	return Ip
}
