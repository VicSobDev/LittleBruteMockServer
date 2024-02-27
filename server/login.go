package server

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *Api) login(c *gin.Context) {

	requestBody := loginRequest{}

	if err := c.BindJSON(&requestBody); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": "Invalid Request Body",
		})
		return
	}

	//sleep to simulate a real login operation
	timeSlept := randomSleep(c.Request.Context(), MAX_SLEEP)

	c.JSON(http.StatusOK, gin.H{
		"error":      "",
		"time_slept": timeSlept,
	})
}
