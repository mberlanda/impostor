package controller

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/mberlanda/impostor/api/models"
	"github.com/mberlanda/impostor/api/service"
)

type UserController struct {
	service *service.UserService
}

func NewUserController(s *service.UserService) UserController {
	return UserController{service: s}
}

func (uc *UserController) Create(c *gin.Context) {
	// TODO: Replace with some request specific models
	user := models.User{}
	err := c.BindJSON(&user)
	if err != nil {
		ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	created, err := uc.service.Create(&user)
	if err != nil {
		ErrorJSON(c, http.StatusBadRequest, err.Error())
		return
	}
	SuccessJSON(c, http.StatusAccepted, created)
}

func (uc *UserController) Index(c *gin.Context) {
	users := uc.service.GetAllUsers()
	SuccessJSON(c, http.StatusOK, users)
}

func (uc *UserController) Show(c *gin.Context) {
	id := c.Param("id")
	if id == "" {
		ErrorJSON(c, http.StatusBadRequest, "invalid user id")
		return
	}
	user, err := uc.service.GetUser(id)
	if err != nil {
		ErrorJSON(c, http.StatusNotFound, err.Error())
		return
	}
	SuccessJSON(c, http.StatusOK, user)
}

func handleError(c *gin.Context, err error) {
	ErrorJSON(c, http.StatusBadRequest, err.Error())
	return
}
