package controller

import (
	"net/http"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type TeacherController struct {
	uc     usecase.TeacherUseCase
	router *gin.Engine
}

func (u *TeacherController) createHandler(c *gin.Context) {
	var payload model.TeacherModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.RegisterNewTeacher(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *TeacherController) listHandler(c *gin.Context) {
	teacher, err := u.uc.FindAllTeacher()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teacher})
}

func (u *TeacherController) getHandler(c *gin.Context) {
	// var id *string = func() *string { t := c.Param("id"); return &t }()
	id := c.Param("id")
	teacher, err := u.uc.FindByTeacherId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": teacher})
}
func (u *TeacherController) updateHandler(c *gin.Context) {
	var payload model.TeacherModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.UpdateTeacher(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *TeacherController) deleteHandler(c *gin.Context) {
	var id string = c.Param("id")
	err := u.uc.DeleteTeacher(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Teacher_id : " + id + " deleted!"})
}

func NewTeacherController(uc usecase.TeacherUseCase, r *gin.Engine) *TeacherController {
	controller := &TeacherController{
		uc:     uc,
		router: r,
	}
	rg := r.Group("/api/v1")
	rg.POST("/teachers", controller.createHandler)
	rg.GET("/teachers", controller.listHandler)
	rg.GET("/teachers/:id", controller.getHandler)
	rg.PUT("/teachers", controller.updateHandler)
	rg.DELETE("/teachers/:id", controller.deleteHandler)

	return controller
}
