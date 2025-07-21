package api

import (
	"net/http"

	"scheduler/internal/scheduler"

	"github.com/gin-gonic/gin"
)

func ScheduleHandler(c *gin.Context) {
	var req struct {
		JobName string `json:"job_name"`
		RunAt   string `json:"run_at"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := scheduler.ScheduleJob(req.JobName, req.RunAt); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusCreated)
}
