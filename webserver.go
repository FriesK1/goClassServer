package main

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"github.com/spf13/viper"
)

func StartWebServer() (err error) {
	var bindUrl string

	bindUrl = strings.Join([]string{viper.GetString("server.host"), viper.GetString("server.port")}, ":")
	gin.SetMode(gin.ReleaseMode)

	router := gin.Default()
	if err = router.SetTrustedProxies([]string{"10.0.0.0/8", "192.168.0.0/16", "172.16.0.0/12", "127.0.0.0/8"}); err != nil {
		logrus.Fatal(err)
	}

	router.GET("/users", GetUserList)
	router.GET("/user/:name", GetUser)
	router.GET("/health-check", HealthCheck)

	_ = router.Run(bindUrl)
	return
}

func GetUserList(c *gin.Context) {
	c.JSON(http.StatusOK, Staff)
}

func GetUser(c *gin.Context) {
	name := c.Param("name")

	if data, ok := Staff[strings.ToLower(name)]; ok {
		c.JSON(http.StatusOK, data)
	} else {
		c.JSON(http.StatusNotFound, gin.H{
			"status": "Error",
			"key":    name,
			"msg":    fmt.Sprintf("No employee record found for key %s", name),
		})
	}

}

func HealthCheck(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"status": "healthy",
	})
}
