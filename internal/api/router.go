package api

import (
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	r := gin.Default()

	r.POST("/schedule", ScheduleHandler)
	r.POST("/run-now", RunNowHandler)
	r.POST("/rerun/:id", ReRunHandler)
	r.POST("/create-job", CreateJobHandler)

	return r
}
