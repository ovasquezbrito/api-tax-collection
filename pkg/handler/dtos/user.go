package dtos

import "github.com/ovasquezbrito/tax-collection/pkg/models"

type GetAllUsersResponse struct {
	Data  []models.UserResponse `json:"data"`
	Total int                   `json:"totalCount"`
}
