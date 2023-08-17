package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"

	"github.com/ovasquezbrito/tax-collection/pkg/handler/dtos"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
)

// @Summary SignUp
// @Tags auth
// @Description create account
// @ID create-account
// @Accept  json
// @Produce  json
// @Param input body dona.User true "account info"
// @Success 200 {object} dtos.Response
// @Failure 400,404 {object} errorResponse
// @Failure 500 {object} errorResponse
// @Failure default {object} errorResponse
// @Router /auth/sign-up [post]

func (h *Handler) createUser(c *gin.Context) {
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

// CreateRole godoc
// @Summary 			Get all users
// @Description 		Return list of users
// @Security 			Bearer
// @Tags 				usuarios
// @Accept 				json
// @Produce 			json
// @Success 			200 {array} dtos.GetAllUsersResponse
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Param 				_page query int false "integer default" default(1)
// @Param 				_limit query int false "integer default" default(10)
// @Param 				name_like query string false "search"
// @Router 				/api/user/ [get]
func (h *Handler) getAllUsers(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		return
	}

	query, err := getPageLimitSearch(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, total, err := h.services.UsersServices.GetAll(c, query)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllUsersResponse{
		Data:  lists,
		Total: total,
	})
}

// CreateRole godoc
// @Summary 			Get a user
// @Description 		Return a user
// @Security 			Bearer
// @Tags 				usuarios
// @Accept 				json
// @Produce 			json
// @Success 			200 {object} dtos.Response
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Param 				id path int true "role id"
// @Router 				/api/user/{id}/show [get]
func (h *Handler) getUserById(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.services.Authorization.GetUserById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	webResponse := dtos.Response{
		Code:   http.StatusOK,
		Status: "Ok",
		Data:   item,
	}

	c.JSON(http.StatusOK, webResponse)
}

// CreateAccount godoc
// @Summary 			Asociate role to account
// @Description 		Asociate role to account
// @Security 			Bearer
// @Tags 					usuarios
// @Accept 				json
// @Produce 			json
// @Param 				input body models.AsociateRoleToUser true "asociate account"
// @Success 			200 {object} dtos.Response
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Router 				/api/user/associaterole [post]
func (h *Handler) associateRoleToUser(c *gin.Context) {
	var input models.AsociateRoleToUser

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	af, err := h.services.UsersServices.AddRoleToUser(c, input)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	messageDelete := fmt.Sprintf("%d RowsAffected", af)

	webResponse := dtos.Response{
		Code:   http.StatusNoContent,
		Status: "Ok",
		Data:   messageDelete,
	}

	c.JSON(http.StatusOK, webResponse)
}
