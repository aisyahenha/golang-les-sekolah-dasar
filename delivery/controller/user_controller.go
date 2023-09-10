package controller

import (
	"net/http"

	"github.com/aisyahenha/golang-les-sekolah-dasar/delivery/middleware"
	"github.com/aisyahenha/golang-les-sekolah-dasar/model"
	"github.com/aisyahenha/golang-les-sekolah-dasar/usecase"
	"github.com/gin-gonic/gin"
	// "github.com/google/uuid"
)

type UserController struct {
	uc             usecase.UserUseCase
	router         *gin.Engine
	authMiddleware middleware.AuthMiddleware
}

func (u *UserController) createHandler(c *gin.Context) {
	var payload model.User
	// var respon model.UserRespon
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.RegisterNewUser(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *UserController) listHandler(c *gin.Context) {
	users, err := u.uc.FindAllUser()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	// var respon model.UserRespon
	c.JSON(http.StatusOK, gin.H{"data": users})
}

func (u *UserController) getHandler(c *gin.Context) {
	// var id *string = func() *string { t := c.Param("id"); return &t }()
	id := c.Param("id")
	user, err := u.uc.FindByUserId(id)

	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"data": user})
}
func (u *UserController) updateHandler(c *gin.Context) {
	var payload model.User
	if err := c.ShouldBindJSON(&payload); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	if err := u.uc.UpdateUser(&payload); err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusCreated, payload)
}
func (u *UserController) deleteHandler(c *gin.Context) {
	var id string = c.Param("id")
	err := u.uc.DeleteUser(id)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "user_id : " + id + " deleted!"})
}

func NewUserController(uc usecase.UserUseCase, r *gin.Engine, am middleware.AuthMiddleware) *UserController {
	controller := &UserController{
		uc:             uc,
		router:         r,
		authMiddleware: am,
	}
	rg := r.Group("/api/v1")
	rg.POST("/users", am.RequireToken("admin"), controller.createHandler)
	rg.GET("/users", am.RequireToken("admin"), controller.listHandler)
	rg.GET("/users/:id", am.RequireToken("admin"), controller.getHandler)
	rg.PUT("/users", am.RequireToken("admin"), controller.updateHandler)
	rg.DELETE("/users/:id", am.RequireToken("admin"), controller.deleteHandler)

	// rg.POST("/users", controller.createHandler)
	// rg.GET("/users", controller.listHandler)
	// rg.GET("/users/:id", controller.getHandler)
	// rg.PUT("/users", controller.updateHandler)
	// rg.DELETE("/users/:id", controller.deleteHandler)

	return controller
}
