package main

import (
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	log.SetFlags(log.Lshortfile | log.LstdFlags)
	r := gin.Default()
	r.POST("/ping", TestingPostAPI)
	r.Run()
}

type PostData struct {
	Name  string `json:"name"`
	Class string `json:"class"`
}

func TestingPostAPI(c *gin.Context) {
	var data PostData
	if err := c.BindJSON(&data); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Please pass all the required",
		})
		return
	}
	if err := testingPostAPI(data); err != nil {
		c.JSON(400, gin.H{
			"error":   true,
			"message": "Please pass all the required",
		})
		return
	}
	c.JSON(200, gin.H{
		"error":   false,
		"message": "Data added successfully",
	})
	return

}

func testingPostAPI(data PostData) error {
	defer recover()

	return nil
}
