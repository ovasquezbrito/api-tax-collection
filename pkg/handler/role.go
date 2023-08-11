package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ovasquezbrito/tax-collection/pkg/handler/dtos"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
)

// CreateRole godoc
// @Summary 			Create role
// @Description 	Register un role for user
// @Tags 					roles
// @Accept 				json
// @Produce 			json
// @Param 				input body dtos.Role true "role info"
// @Success 			200 {object} dtos.Response
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Router 				api/roles/new [post]
func (h *Handler) createRole(c *gin.Context) {
	var input dtos.Role

	if err := c.BindJSON(&input); err != nil {
		newErrorResponse(c, http.StatusBadRequest, err.Error())
		return
	}

	r := models.Role{
		RoleName:  input.RoleName,
		RoleNivel: input.RoleNivel,
	}

	id, err := h.services.RoleService.CreateRole(c, r)
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
// @Summary 			Get all role
// @Description 	Return list of roles
// @Tags 					roles
// @Accept 				json
// @Produce 			json
// @Param 				input body dtos.Role true "role info"
// @Success 			200 {object} dtos.Response
// @Failure 			400,404 {object} errorResponse
// @Failure 			500 {object} errorResponse
// @Router 				api/roles/new [post]
func (h *Handler) getAllRoles(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		return
	}

	query, err := getPageLimitSearch(c)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	lists, total, err := h.services.RoleService.GetAll(c, query)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, dtos.GetAllRolesResponse{
		Data:  lists,
		Total: total,
	})
}

func (h *Handler) getRoleById(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.services.RoleService.GetById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": item,
	})
}

func (h *Handler) deleteRoleById(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	item, err := h.services.RoleService.DeleteById(c, id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}
	messageDelete := fmt.Sprintf("%d RowsAffected", item)

	c.JSON(http.StatusOK, map[string]interface{}{
		"code":    204,
		"message": messageDelete,
	})
}
