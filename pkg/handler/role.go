package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	stopapp "github.com/ovasquezbrito/base-app"
)

type getAllRolesResponse struct {
	Data  []stopapp.Role `json:"data"`
	Total int            `json:"totalCount"`
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

	lists, total, err := h.services.RoleService.GetAll(query)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, getAllRolesResponse{
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

	item, err := h.services.RoleService.GetById(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": item,
	})
}
