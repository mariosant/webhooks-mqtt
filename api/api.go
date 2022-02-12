package api

import (
	"encoding/json"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/mochi-co/mqtt/server"
)

type Configuration struct {
	Secret     string
	Port       int
	MqttServer *server.Server
}

func addMqtt(mqttServer *server.Server) gin.HandlerFunc {
	return func(c *gin.Context) {
		c.Set("mqttServer", mqttServer)
		c.Next()
	}
}

func webhooksHandler(c *gin.Context) {
	var data interface{}

	if err := json.NewDecoder(c.Request.Body).Decode(&data); err != nil {
		c.String(400, "Could not parse body")
		fmt.Errorf(err.Error())

		return
	}

	mqttServer, _ := c.Value("mqttServer").(*server.Server)

	go func() {
		marshalledData, _ := json.Marshal(data)
		mqttServer.Publish("test", marshalledData, false)
	}()

	c.String(200, "ok")
}

func CreateServer(config *Configuration) *gin.Engine {
	server := gin.Default()

	server.POST("/webhooks", addMqtt(config.MqttServer), webhooksHandler)
	return server
}
