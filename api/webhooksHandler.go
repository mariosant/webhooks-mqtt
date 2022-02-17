package api

import (
	"fmt"
	"strconv"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/mochi-co/mqtt/server"
)

func publishMqtt(mqttServer *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			data, _ := c.Value("data").(*gabs.Container)

			organization := strconv.FormatFloat(data.Path("license_id").Data().(float64), 'f', 0, 64)
			action := data.Path("action").Data().(string)

			fmt.Println(organization, action)

			mqttServer.Publish(organization+"/"+action, data.Bytes(), false)
		}()

		c.Next()
	}
}

func assignBody() gin.HandlerFunc {
	return func(c *gin.Context) {
		data, err := gabs.ParseJSONBuffer(c.Request.Body)

		if err != nil {
			c.String(400, "Could not parse body")
			c.Abort()

			fmt.Println(err.Error())

			return
		}

		c.Set("data", data)

		c.Next()
	}
}

func requireSecret(secret string) gin.HandlerFunc {
	return func(c *gin.Context) {
		data, _ := c.Value("data").(*gabs.Container)

		if webhookSecret := data.Path("secret_key").Data(); webhookSecret != secret {
			c.String(400, "Could not parse body")
			c.Abort()

			return
		}

		c.Next()
	}
}

func webhooksHandler(c *gin.Context) {
	c.String(200, "ok")
}
