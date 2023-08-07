package handler

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/ovasquezbrito/tax-collection/pkg/handler/dtos"
	"github.com/ovasquezbrito/tax-collection/pkg/models"
)

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

	fmt.Println(r)
	id, err := h.services.RoleService.CreateRole(c, r)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"id_uuid": id,
	})

}

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

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": item,
	})
}
