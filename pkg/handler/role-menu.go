package handler

import (
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
	baseapp "github.com/ovasquezbrito/base-app"
)

type getAllRolesMenuResponse struct {
	Data []baseapp.RoleUser `json:"data"`
}

func (h *Handler) getAllMenuRoleUserById(c *gin.Context) {
	_, err := getUserId(c)

	if err != nil {
		return
	}

	id, err := strconv.Atoi(c.Param("id"))
	if err != nil {
		newErrorResponse(c, http.StatusBadRequest, "invalid id param")
		return
	}

	user, err := h.services.Authorization.GetMenuOptionAll(id)
	if err != nil {
		newErrorResponse(c, http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, map[string]interface{}{
		"data": user,
	})
}
