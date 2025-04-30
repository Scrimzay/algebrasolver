package main
	
import (
	"log"

	"github.com/gin-gonic/gin"
)
	
func main() {
	r := gin.Default()
	r.LoadHTMLGlob("**/*.html")

	r.GET("/", indexHandler)

	err := r.Run(":4000")
	if err != nil {
		log.Fatal(err)
	}
}
	
func indexHandler(c *gin.Context) {
	c.HTML(200, "index.html", nil)
}
