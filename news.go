package noui

import (
	"fmt"
	"time"

	"github.com/tmlbl/gin"
)

// News represents an event in time that can be displayed in an activity feed
// type interface.
type News struct {
	Time     time.Time
	Headline string
	Content  string
}

func handlePostNews(c *gin.Context) {
	n := News{}
	err := c.BindJSON(&n)
	if err != nil {
		c.JSON(200, ErrorResponse{"Hunta dunta"})
	}
	fmt.Println(n)
	conn.DB("noui").C("news").Insert(&n)
	c.Status(200)
}
