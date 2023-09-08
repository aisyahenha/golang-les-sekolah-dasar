package controller

import (
	"net/http"

	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type StudentController struct {
	uc     usecase.StudentUseCase
	router *gin.Engine
}

func (u *StudentController) createHandler(c *gin.Context) {
	var payload model.StudentModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.RegisterNewStudent(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *StudentController) listHandler(c *gin.Context) {
	student, err := u.uc.FindAllStudent()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}

func (u *StudentController) getHandler(c *gin.Context) {
	// var id *string = func() *string { t := c.Param("id"); return &t }()
	id := c.Param("id")
	student, err := u.uc.FindByStudentId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": student})
}
func (u *StudentController) updateHandler(c *gin.Context) {
	var payload model.StudentModel
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.UpdateStudent(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *StudentController) deleteHandler(c *gin.Context) {
	var id string = c.Param("id")
	err := u.uc.DeleteStudent(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Student_id : " + id + " deleted!"})
}

func NewStudentController(uc usecase.StudentUseCase, r *gin.Engine) *StudentController {
	controller := &StudentController{
		uc:     uc,
		router: r,
	}
	rg := r.Group("/api/v1")
	rg.POST("/students", controller.createHandler)
	rg.GET("/students", controller.listHandler)
	rg.GET("/students/:id", controller.getHandler)
	rg.PUT("/students", controller.updateHandler)
	rg.DELETE("/students/:id", controller.deleteHandler)

	return controller
}
