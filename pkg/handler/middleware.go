package handler

import (
	"errors"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/gin-gonic/gin"

	"github.com/ovasquezbrito/tax-collection/pkg/models"
	"github.com/ovasquezbrito/tax-collection/token"
)

const (
	authorizationHeaderKey  = "authorization"
	authorizationTypeBearer = "bearer"
	authorizationPayloadKey = "authorization_payload"
)

// AuthMiddleware creates a gin middleware for authorization
func (h *Handler) authMiddleware(tokenMaker token.Maker) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		authorizationHeader := ctx.GetHeader(authorizationHeaderKey)

		if len(authorizationHeader) == 0 {
			err := errors.New("authorization header is not provided")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseError(err))
			return
		}

		fields := strings.Fields(authorizationHeader)
		if len(fields) < 2 {
			err := errors.New("invalid authorization header format")
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseError(err))
			return
		}

		authorizationType := strings.ToLower(fields[0])
		if authorizationType != authorizationTypeBearer {
			err := fmt.Errorf("unsupported authorization type %s", authorizationType)
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseError(err))
			return
		}

		accessToken := fields[1]
		payload, err := tokenMaker.VerifyToken(accessToken)
		if err != nil {
			ctx.AbortWithStatusJSON(http.StatusUnauthorized, responseError(err))
			return
		}

		ctx.Set(authorizationPayloadKey, payload)
		ctx.Next()
	}
}

func getUserId(c *gin.Context) (string, error) {
	authPayload, ok := c.MustGet(authorizationPayloadKey).(*token.Payload)

	if !ok {
		newErrorResponse(c, http.StatusInternalServerError, "user id not found")
		return "", errors.New("user id not found")
	}

	return authPayload.Email, nil
}

func getPageLimitSearch(c *gin.Context) (models.QueryParameter, error) {
	var query models.QueryParameter

	limit := c.Query("_limit")
	page := c.Query("_page")
	search := c.Query("search_like")

	if sv, err := strconv.Atoi(limit); err == nil {
		query.Limit = sv
	} else {
		newErrorResponse(c, http.StatusInternalServerError, "el parámetro limit no es un número real")
		return query, errors.New("el parámetro limit no es un número real")
	}
	if sv, err := strconv.Atoi(page); err == nil {
		query.Page = sv
	} else {
		newErrorResponse(c, http.StatusInternalServerError, "el parámetro page no es un número real")
		return query, errors.New("el parámetro page no es un número real")
	}

	query.Search = search

	return query, nil
}
