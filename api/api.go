package api

import (
	"github.com/gin-gonic/gin"
	"github.com/mochi-co/mqtt/server"
)

type Configuration struct {
	Secret     string
	Port       int
	MqttServer *server.Server
}

func CreateServer(config *Configuration) *gin.Engine {
	server := gin.Default()

	server.POST("/webhooks", assignBody(), requireSecret(config.Secret), publishMqtt(config.MqttServer), webhooksHandler)
	server.POST("/me", meHandler(config.Secret))

	return server
}
