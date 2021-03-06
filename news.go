package noui

import (
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/tmlbl/gin"
)

const (
	newsCollection = "news"
	errBadJSON     = "JSON body is malformed"
)

type Model struct {
	Namespace string   `json:"namespace"`
	Roles     []string `json:"roles"`
}

// News represents an event in time that can be displayed in an activity feed
// type interface.
type News struct {
	Model
	Time     time.Time `json:"time"`
	Headline string    `json:"headline"`
	Content  string    `json:"content"`
}

func handlePostNews(c *gin.Context) {
	n := News{}
	err := c.BindJSON(&n)
	if err != nil {
		c.JSON(400, ErrorResponse{errBadJSON})
		c.Abort()
		return
	}
	// If the time is null, fill in the current time as a default.
	if n.Time.Equal(time.Time{}) {
		n.Time = time.Now()
	}
	err = db.C(newsCollection).Insert(&n)
	if err != nil {
		c.JSON(400, ErrorResponse{err.Error()})
		c.Abort()
		return
	}
	c.Status(200)
}

func handleGetNews(c *gin.Context) {
	namespace := c.Param("namespace")
	n := []News{}
	err := db.C(newsCollection).Find(bson.M{
		"model.namespace": namespace,
	}).Sort("time -1").All(&n)
	if err != nil {
		c.JSON(400, ErrorResponse{err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, &n)
}
