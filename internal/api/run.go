package api

import (
	"net/http"

	"scheduler/internal/scheduler"

	"github.com/gin-gonic/gin"
)

func RunNowHandler(c *gin.Context) {
	if err := scheduler.RunScheduledJobs(); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
