package delivery

import (
	"aulia-majoo-test/internal/domain"
	"aulia-majoo-test/internal/dto"
	"errors"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler struct {
	ucase domain.UserUsecase
}

func NewUserHandler(ucase domain.UserUsecase) domain.UserHandler {
	return &UserHandler{
		ucase: ucase,
	}
}

// Login  godoc
// @Summary Login
// @Description Login
// @Tags Users
// @Accept json
// @Produce json
// @Param User body dto.LoginRequestDTO true "Login"
// @Success 200 {object} dto.LoginResponseDTO
// @Router /auth/login [post]
func (pm UserHandler) Login(c *gin.Context) {

	// get request body and bind json
	var req dto.LoginRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false})
		return
	}

	res, err := pm.ucase.UserLogin(req)
	if err != nil {
		fmt.Println(err.Error())
		c.JSON(http.StatusUnauthorized, errors.New(err.Error()))
	}

	c.JSON(http.StatusOK, res)
}

// Register  godoc
// @Summary Register
// @Description Register
// @Tags Users
// @Accept json
// @Produce json
// @Param User body dto.RegisterRequestDTO true "Register"
// @Success 200 {object} dto.RegisterResponseDTO
// @Router /auth/register [post]
func (pm UserHandler) Register(c *gin.Context) {

	// get request body and bind json
	var req dto.RegisterRequestDTO

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"message": err.Error(),
			"status":  false})
		return
	}

	// business logic handle on this service
	res, err := pm.ucase.UserRegister(req)
	if err != nil {
		c.JSON(http.StatusUnauthorized, err)
	}

	c.JSON(http.StatusOK, res)
}
