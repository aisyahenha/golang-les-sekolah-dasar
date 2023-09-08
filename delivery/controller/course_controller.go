package controller

import (
	"net/http"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type CourseController struct {
	uc     usecase.CourseUseCase
	router *gin.Engine
}

func (u *CourseController) createHandler(c *gin.Context) {
	var payload model.CourseModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.RegisterNewCourse(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *CourseController) listHandler(c *gin.Context) {
	corses, err := u.uc.FindAllCourse()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": corses})
}

func (u *CourseController) getHandler(c *gin.Context) {
	// var id *string = func() *string { t := c.Param("id"); return &t }()
	id := c.Param("id")
	course, err := u.uc.FindByCourseId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": course})
}
func (u *CourseController) updateHandler(c *gin.Context) {
	var payload model.CourseModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.UpdateCourse(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *CourseController) deleteHandler(c *gin.Context) {
	var id string = c.Param("id")
	err := u.uc.DeleteCourse(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "course_id : " + id + " deleted!"})
}

func NewCourseController(uc usecase.CourseUseCase, r *gin.Engine) *CourseController {
	controller := &CourseController{
		uc:     uc,
		router: r,
	}
	rg := r.Group("/api/v1")
	rg.POST("/courses", controller.createHandler)
	rg.GET("/courses", controller.listHandler)
	rg.GET("/courses/:id", controller.getHandler)
	rg.PUT("/courses", controller.updateHandler)
	rg.DELETE("/courses/:id", controller.deleteHandler)

	return controller
}
