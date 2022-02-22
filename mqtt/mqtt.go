package mqtt

import (
	"fmt"
	"log"
	"strings"

	"webhooks/lib"

	m "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

type Auth struct {
	secret string
}

func (a *Auth) Authenticate(user, password []byte) bool {
	expectedPassword := lib.GeneratePassword(string(user), a.secret)

	fmt.Println("Authentication request: " + string(user))

	return expectedPassword == string(password)
}

func (a *Auth) ACL(user []byte, topic string, write bool) bool {
	isLegitTopic := strings.HasPrefix(topic, string(user))
	isReadOnly := !write

	fmt.Println("Authorization request: " + string(user))

	return isLegitTopic && isReadOnly
}

func CreateMqtt(secret string) *m.Server {
	server := m.New()

	listener := listeners.NewWebsocket("ws1", ":1883")

	auth := &Auth{
		secret: secret,
	}

	err := server.AddListener(listener, &listeners.Config{
		Auth: auth,
	})

	if err != nil {
		log.Fatal(err)
	}

	return server
}
