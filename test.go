package main

import (
	"encoding/json"
	"log"

	"github.com/gin-gonic/gin"
)

// DingResponse response to a /ding request
type DingResponse struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}

func main() {
	log.Println("Testing build/execution of go program in docker container")
	r := gin.Default()

	// ping
	r.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	// ding
	r.GET("/ding", func(c *gin.Context) {
		jsonStr := `
		{
			"name": "harry",
			"age": 26
		}`
		dingResponse := DingResponse{}
		if err := json.Unmarshal([]byte(jsonStr), &dingResponse); err == nil {
			enc := json.NewEncoder(c.Writer)
			enc.SetIndent("", "  ")
			enc.Encode(dingResponse)
		} else {
			log.Panic(err)
		}
	})

	r.Run()
}
