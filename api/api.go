package api

import (
	"github.com/gin-contrib/cors"
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

	corsConfig := cors.DefaultConfig()
	corsConfig.AllowAllOrigins = true
	corsConfig.AddAllowHeaders("Authorization")

	server.Use(cors.New(corsConfig))

	server.POST("/webhooks", assignBody(), requireSecret(config.Secret), publishMqtt(config.MqttServer), webhooksHandler)
	server.GET("/me", meHandler(config.Secret))

	return server
}
