package controller

import (
	"net/http"

	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery/middleware"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type ScheduleController struct {
	uc             usecase.ScheduleUseCase
	router         *gin.Engine
	authMiddleware middleware.AuthMiddleware
}

func (u *ScheduleController) createHandler(c *gin.Context) {
	var payload model.ScheduleModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.RegisterNewSchedule(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *ScheduleController) listHandler(c *gin.Context) {
	schedule, err := u.uc.FindAllSchedule()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": schedule})
}

func (u *ScheduleController) getHandler(c *gin.Context) {
	// var id *string = func() *string { t := c.Param("id"); return &t }()
	id := c.Param("id")
	schedule, err := u.uc.FindByScheduleId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": schedule})
}
func (u *ScheduleController) updateHandler(c *gin.Context) {
	var payload model.ScheduleModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.UpdateSchedule(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *ScheduleController) deleteHandler(c *gin.Context) {
	var id string = c.Param("id")
	err := u.uc.DeleteSchedule(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Schedule_id : " + id + " deleted!"})
}

func NewScheduleController(uc usecase.ScheduleUseCase, r *gin.Engine, am middleware.AuthMiddleware) *ScheduleController {
	controller := &ScheduleController{
		uc:             uc,
		router:         r,
		authMiddleware: am,
	}
	rg := r.Group("/api/v1")
	rg.POST("/schedules", am.RequireToken("admin"), controller.createHandler)
	rg.GET("/schedules", am.RequireToken("*"), controller.listHandler)
	rg.GET("/schedules/:id", am.RequireToken("admin"), controller.getHandler)
	rg.PUT("/schedules", am.RequireToken("admin"), controller.updateHandler)
	rg.DELETE("/schedules/:id", am.RequireToken("admin"), controller.deleteHandler)

	return controller
}
