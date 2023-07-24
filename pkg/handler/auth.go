package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	baseapp "github.com/ovasquezbrito/base-app"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body dona.User true "account info"
// @Success 200 {integer} integer 1
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]

type signInInputUser struct {
	Email    string `json:"email" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type UserResponse struct {
	ID     int    `json:"id"`
	Email  string `json:"email"`
	Name   string `json:"name"`
	UriImg string `json:"uri_img"`
	Token  string `json:"token"`
}

type tokenUser struct {
	Token string `json:"token" binding:"required"`
}

func (h *Handler) register(c *gin.Context) {
	var input baseapp.User

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	existUserName, err := h.services.Authorization.GetUserByUserName(input.Email)

	if existUserName > 0 {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	id, err := h.services.Authorization.CreateUser(input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id": id,
	})
}

func (h *Handler) login(c *gin.Context) {
	var input signInInputUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}
	data, err := h.services.Authorization.LoginUser(input.Email, input.Password)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, UserResponse{
		ID:     data.UserLogin.Id,
		Email:  data.UserLogin.Email,
		Name:   data.UserLogin.FirstLast,
		UriImg: data.UserLogin.UriImg,
		Token:  data.AccessToken,
	})

}

func (h *Handler) verifyToken(c *gin.Context) {
	var input tokenUser

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
