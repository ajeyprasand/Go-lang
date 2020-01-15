package main

import (
"github.com/gin-gonic/gin"
"net/http"
)

const name string = "ajey"

func main() {
      router := gin.Default()
      router.GET("/hello", func(c *gin.Context) {
c.JSON(http.StatusOK, gin.H{
"message":"Hello World!",
})
})
      router.GET("/greet/:name", func(c *gin.Context) {
	name := c.Param("name")
	c.JSON(http.StatusOK, gin.H{ 
"message": "Hello " + name + "!",
})
})

router.Run(":8080")
}




