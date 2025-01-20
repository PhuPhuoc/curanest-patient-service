package middleware

import (
	"context"
	"errors"
	"strings"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/gin-gonic/gin"
)

type AuthClient interface {
	IntrospectToken(ctx context.Context, accessToken string) (common.Requester, error)
}

func RequireAuth(ac AuthClient) func(*gin.Context) {
	return func(ctx *gin.Context) {
		token, err := extractTokenFromHeaderString(ctx.GetHeader("Authorization"))
		if err != nil {
			common.ResponseError(ctx, err)
			ctx.Abort()
			return
		}

		requester, err := ac.IntrospectToken(ctx.Request.Context(), token)
		if err != nil {
			common.ResponseError(ctx, err)
			ctx.Abort()
			return
		}

		ctx.Set(common.KeyRequester, requester)
		ctx.Next()
	}
}

func extractTokenFromHeaderString(s string) (string, error) {
	parts := strings.Split(s, " ") //"Authorization" : "Bearer {token}"

	if parts[0] != "Bearer" || len(parts) < 2 || strings.TrimSpace(parts[1]) == "" {
		return "", errors.New("missing access token")
	}

	return parts[1], nil
}
