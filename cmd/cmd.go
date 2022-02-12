package cmd

import (
	"fmt"
	"os"
	"webhooks/api"
	"webhooks/mqtt"

	"github.com/urfave/cli/v2"
)

func App() error {
	app := &cli.App{
		Name: "Livechat webhooks",
		Flags: []cli.Flag{
			&cli.StringFlag{
				Name:    "Secret",
				EnvVars: []string{"SECRET"},
			},
		},
		Action: func(c *cli.Context) error {
			fmt.Print("Skata")
			// secret := c.String("Secret")
			mqttServer := mqtt.CreateMqtt()

			apiConfig := api.Configuration{
				MqttServer: mqttServer,
				Port:       3000,
				Secret:     "secret",
			}

			apiServer := api.CreateServer(&apiConfig)

			go func() { mqttServer.Serve() }()
			apiServer.Run()

			return nil
		},
	}

	return app.Run(os.Args)
}
