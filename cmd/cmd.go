package cmd

import (
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
				Name:    "secret",
				EnvVars: []string{"SECRET"},
				Required: true,
			},
		},
		Action: func(c *cli.Context) error {
			mqttServer := mqtt.CreateMqtt(c.String("secret"))

			apiConfig := api.Configuration{
				MqttServer: mqttServer,
				Port:       3000,
				Secret:     c.String("secret"),
			}

			apiServer := api.CreateServer(&apiConfig)

			go func() { mqttServer.Serve() }()
			apiServer.Run()

			return nil
		},
	}

	return app.Run(os.Args)
}
