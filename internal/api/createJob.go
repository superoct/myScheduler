package api

import (
	"net/http"

	"scheduler/internal/jobs"

	"github.com/gin-gonic/gin"
)

func CreateJobHandler(c *gin.Context) {
	var req struct {
		Name              string   `json:"name"`
		Command           string   `json:"command"`
		Dependencies      []string `json:"dependencies"`
		ScheduleType      string   `json:"schedule_type"`
		ScheduleStartTime string   `json:"schedule_start_time"`
		RepeatEvery       int      `json:"repeat_every"`
		RepeatAmount      int      `json:"repeat_amount"`
		ConnectionID      int      `json:"connection_id"`
		ParentID          int      `json:"parent_id"`
		VariableIDs       []int    `json:"variable_ids"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}
	if err := jobs.CreateJob(
		req.Name, req.Command, req.Dependencies, req.ScheduleType,
		req.ScheduleStartTime, req.RepeatEvery, req.RepeatAmount,
		req.ConnectionID, req.ParentID, req.VariableIDs,
	); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	c.Status(http.StatusCreated)
}
