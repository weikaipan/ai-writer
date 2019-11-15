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
	/*
	 *  Method: GET
	 *  Return a specific basketball court information.
	 */
	name := c.Query("name")
	courtInfo, err := db.RetrieveCourtInfo(name) 
	if err == nil {
		c.JSON(http.StatusOK, gin.H {
			"lon": courtInfo.Lon,
			"lat": courtInfo.Lat,
			"name": courtInfo.name,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "error",
		})
	}
}

func SpamCourt(c *gin.Context) {
	/*
	 *  Method: POST
	 *  User specifies wrong reported court
	 *  by its id.
	 */
	courtId := c.Query("id")
	message, err := db.SpamCourtById(0)
	if err == nil {
		c.JSON(http.StatusOK, gin.H{
			"message": message,
			"id": courtId,
		})
	} else {
		c.JSON(http.StatusInternalServerError, gin.H {
			"error": "error",
		})
	}
}

func UploadCourt(c *gin.Context) {
	/*
	 *	Method: POST
	 *  User uploads a basketball court's location, 
	 *  indoor or outdoor, and other information.
	 */
	 lon := c.Query("lon")
	 lat := c.Query("lat")
	 indoor := c.Query("indoor")

	 message, err = db.SaveNewCourt()
	 if err != nil {
		c.JSON(http.StatusOK, gin.H{
			"message": message,
		})
	 } else {
		c.JSON(http.StatusInternalServerError, gin.H{
			"error": "error",
		})
	 }
}
