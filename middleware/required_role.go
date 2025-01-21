package middleware

import (
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/gin-gonic/gin"
)

func RequireRole(allowedRoles ...string) gin.HandlerFunc {
	return func(ctx *gin.Context) {
		req, ok := ctx.Request.Context().Value(common.KeyRequester).(common.Requester)
		if !ok {
			common.ResponseUnauthorizedError(ctx, "cannot found requester info")
			ctx.Abort()
			return
		}

		for _, role := range allowedRoles {
			if req.Role() == role {
				fmt.Println("role check role: ", req.Role())
				ctx.Next()
				return
			}
		}

		common.ResponseFobiddenError(ctx, "your role cannot use this api")
		ctx.Abort()
	}
}
