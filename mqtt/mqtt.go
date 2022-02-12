package mqtt

import (
	"fmt"
	"log"

	m "github.com/mochi-co/mqtt/server"
	"github.com/mochi-co/mqtt/server/listeners"
)

type Auth struct{}

func (a *Auth) Authenticate(user, password []byte) bool {
	fmt.Println(user, password)
	return true
}

func (a *Auth) ACL(user []byte, topic string, write bool) bool {
	u := string(user)
	fmt.Println(u, topic, write)

	return true
}

func CreateMqtt() *m.Server {
	server := m.New()

	tcp := listeners.NewTCP("t1", ":1883")

	err := server.AddListener(tcp, &listeners.Config{
		Auth: new(Auth),
	})

	if err != nil {
		log.Fatal(err)
	}

	return server
}
