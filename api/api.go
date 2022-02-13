package api

import (
	"fmt"

	"github.com/Jeffail/gabs/v2"
	"github.com/gin-gonic/gin"
	"github.com/mochi-co/mqtt/server"
)

type Configuration struct {
	Secret     string
	Port       int
	MqttServer *server.Server
}

func publishMqtt(mqttServer *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		go func() {
			data, _ := c.Value("data").(*gabs.Container)

			organization := data.Path("organization_id").Data().(string)
			action := data.Path("action").Data().(string)

			mqttServer.Publish(string(organization)+"/"+action, data.Bytes(), false)
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

		if webhookSecret := data.Path("secret").Data(); webhookSecret != secret {
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

func CreateServer(config *Configuration) *gin.Engine {
	server := gin.Default()

	server.POST("/webhooks", assignBody(), requireSecret(config.Secret), publishMqtt(config.MqttServer), webhooksHandler)
	return server
}
