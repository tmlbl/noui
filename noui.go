package noui

import (
	"gopkg.in/mgo.v2"
)

var conn *mgo.Session

type Config struct {
	HostName string
}

func serve(c Config) {
	sesh, err := mgo.Dial(c.HostName)
	if err != nil {
		panic(err)
	}
	conn = sesh
}

func Serve() {
	serve(Config{"localhost"})
}
