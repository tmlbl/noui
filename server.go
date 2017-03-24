package noui

import (
	"github.com/tmlbl/gin"
	"gopkg.in/mgo.v2"
)

var db *mgo.Database

// Config is a container for the options available when starting the server.
type Config struct {
	Prefix   string
	HostName string
	DBName   string
}

type ErrorResponse struct {
	Message string `json:"message"`
}

func NewServer() *gin.Engine {
	app := gin.New()

	app.GET("/api/news/:namespace", handleGetNews)
	app.POST("/api/news", handlePostNews)

	return app
}

func dbconnect(c Config) {
	// Connect to mongodb.
	sesh, err := mgo.Dial(c.HostName)
	if err != nil {
		panic(err)
	}
	db = sesh.DB(c.DBName)
}

func serve(c Config) {
	dbconnect(c)
	app := NewServer()
	app.Run("0.0.0.0:7070")
}

// Serve starts the server with a minimal default configuration.
func Serve() {
	serve(Config{
		HostName: "localhost",
		DBName:   "noui",
	})
}
