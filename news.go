package noui

import (
	"fmt"
	"time"

	"gopkg.in/mgo.v2/bson"

	"github.com/tmlbl/gin"
)

const (
	errBadJSON = "JSON body is malformed"
)

// News represents an event in time that can be displayed in an activity feed
// type interface.
type News struct {
	Namespace string    `json:"namespace"`
	Time      time.Time `json:"time"`
	Headline  string    `json:"headline"`
	Content   string    `json:"content"`
}

func handlePostNews(c *gin.Context) {
	n := News{}
	err := c.BindJSON(&n)
	if err != nil {
		c.JSON(400, ErrorResponse{errBadJSON})
		c.Abort()
		return
	}
	fmt.Println(n)
	// If the time is null, fill in the current time as a default.
	if n.Time.Equal(time.Time{}) {
		n.Time = time.Now()
	}
	err = db.C("news").Insert(&n)
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
	err := db.C("news").Find(bson.M{"namespace": namespace}).All(&n)
	if err != nil {
		c.JSON(400, ErrorResponse{err.Error()})
		c.Abort()
		return
	}
	c.JSON(200, &n)
}
