package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
	db "./db"
)

func main() {
	router := gin.Default()
	
	api := router.Group("/api")
	{
	  api.GET("/", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H {
		  "message": "pong",
		})
	  })
	}

	api.GET("/court/display", DisplayCourtStatus)
	api.POST("/court/collection/spam",SpamCourt)
	api.POST("/court/collection/upload", UploadCourt)
	router.Run() // listen and serve on 0.0.0.0:8080 (for windows "localhost:8080")
}

func DisplayCourtStatus(c *gin.Context) {
	firstname := c.DefaultQuery("firstname", "Guest")
	lastname := c.Query("lastname") // shortcut for c.Request.URL.Query().Get("lastname")
	courtInfo, err := db.RetrieveCourtInfo() 
	if err == nil {
		c.JSON(http.StatusOK, gin.H {
			"lon": courtInfo.Lon,
			"lat": courtInfo.Lat,
			"firstname": firstname,
			"lastname": lastname,
		})		
	}
}

func SpamCourt(c *gin.Context) {

}

func UploadCourt(c *gin.Context) {

}