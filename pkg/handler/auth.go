package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/ovasquezbrito/tax-collection/pkg/handler/dtos"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
)

// CreateAccount godoc
// @Summary 			Create account
// @Description 	Register account in the system
// @Tags 					auth
// @Accept 				json
// @Produce 			json
// @Param 				input body dtos.RegisterUser true "account info"
// @Success 			200 {object} dtos.Response
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Router 				/auth/register [post]
func (h *Handler) register(c *gin.Context) {
	var input dtos.RegisterUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	dto := models.User{
		FirstLast:  input.FirstLast,
		Email:      input.Email,
		Password:   input.Password,
		AvatarUser: input.AvatarUser,
	}

	id, err := h.services.Authorization.CreateUser(c, dto)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	webResponse := dtos.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   id,
	}
	c.JSON(http.StatusOK, webResponse)
}

// Login godoc
// @Summary 			Login
// @Description 	Login account
// @Tags 					auth
// @Accept 				json
// @Produce 			json
// @Param 				input body dtos.LoginUser true "account info"
// @Success 			200 {object} dtos.Response
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Router 				/auth/login [post]
func (h *Handler) login(c *gin.Context) {
	var input dtos.LoginUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	data, err := h.services.Authorization.LoginUser(c, input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	webResponse := dtos.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data: dtos.UserResponse{
			ID:         data.UserLogin.Id,
			Email:      data.UserLogin.Email,
			Name:       data.UserLogin.FirstLast,
			AvatarUser: data.UserLogin.AvatarUser,
			Token:      data.AccessToken,
		},
	}

	c.JSON(http.StatusOK, webResponse)

}

func (h *Handler) verifyToken(c *gin.Context) {
	var input dtos.TokenUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	token, err := h.services.Authorization.VerifyToken(input.Token)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusUnauthorized, responseError(err))
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"token": token,
	})

}
