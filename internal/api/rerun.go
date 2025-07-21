package api

import (
	"net/http"
	"strconv"

	"scheduler/internal/scheduler"

	"github.com/gin-gonic/gin"
)

func ReRunHandler(c *gin.Context) {
	idStr := c.Param("id")
	id, err := strconv.Atoi(idStr)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid job run ID"})
		return
	}
	if err := scheduler.ManualRerun(id); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.Status(http.StatusOK)
}
