package noui

import (
	"github.com/tmlbl/gin"
	"gopkg.in/mgo.v2"
)

var conn *mgo.Session

// Config is a container for the options available when starting the server.
type Config struct {
	Prefix   string
	HostName string
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewServer() *gin.Engine {
	app := gin.New()

	app.POST("/api/news", handlePostNews)

	return app
}

func serve(c Config) {
	// Connect to mongodb.
	sesh, err := mgo.Dial(c.HostName)
	if err != nil {
		panic(err)
	}
	conn = sesh
	app := NewServer()
	app.Run("0.0.0.0:7070")
}

// Serve starts the server with a minimal default configuration.
func Serve() {
	serve(Config{
		HostName: "localhost",
	})
}