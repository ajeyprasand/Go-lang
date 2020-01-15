package main

import "github.com/gin-gonic/gin"

func main() {
      router := gin.Default()
      router.GET("/hello", func(c *gin.Context) {
c.JSON(200, gin.H{
"message":"Hello World!",
})
})
      router.GET("/greet/:name", func(c *gin.Context) {
	name := c.Param("name")
	c.JSON(200, gin.H{ 
"greet-message": "Hello " + name + "!",
})
})

router.Run(":8080")
}




