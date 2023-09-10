package controller

import (
	"net/http"

	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery/middleware"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type StudentController struct {
	uc             usecase.StudentUseCase
	router         *gin.Engine
	authMiddleware middleware.AuthMiddleware
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

func NewStudentController(uc usecase.StudentUseCase, r *gin.Engine, am middleware.AuthMiddleware) *StudentController {
	controller := &StudentController{
		uc:             uc,
		router:         r,
		authMiddleware: am,
	}
	rg := r.Group("/api/v1")
	rg.POST("/students", am.RequireToken("admin"), controller.createHandler)
	rg.GET("/students", am.RequireToken("admin"), controller.listHandler)
	rg.GET("/students/:id", am.RequireToken("admin"), controller.getHandler)
	rg.PUT("/students", am.RequireToken("admin"), controller.updateHandler)
	rg.DELETE("/students/:id", am.RequireToken("admin"), controller.deleteHandler)

	return controller
}
