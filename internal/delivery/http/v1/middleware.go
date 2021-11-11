package v1

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
	"strings"
)

var (
	authHeaderName = "Authorization"
	bearerName = "Bearer"
	userCtx    = "userId"
)

func (h *Handler) AuthMiddleware(ctx *gin.Context) {
	authHeader := ctx.GetHeader(authHeaderName)
	if authHeader == "" {
		NewErrorResponse(ctx, http.StatusUnauthorized, "empty auth header")
		return
	}

	headerParts := strings.Split(authHeader, " ")
	if len(headerParts) != 2 || headerParts[0] != bearerName {
		NewErrorResponse(ctx, http.StatusUnauthorized, "not valid auth header")
		return
	}

	id, err := h.services.Auth.CheckToken(ctx, headerParts[1])
	if err != nil {
		NewErrorResponse(ctx, http.StatusUnauthorized, err.Error())
		return
	}

	ctx.Set(userCtx, id)
}

func GetUserIdFromCtx(ctx *gin.Context) (string, error) {
	idFromCtx, exists := ctx.Get(userCtx)
	if !exists {
		return "", fmt.Errorf("not exists userId in context")
	}

	return idFromCtx.(string), nil
}