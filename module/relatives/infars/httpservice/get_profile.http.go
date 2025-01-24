package relativeshttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/gin-gonic/gin"
)

// @Summary		get profile of relatives account
// @Description	get profile of relatives account
// @Tags			relative
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/api/v1/relatives/me [get]
// @Security		ApiKeyAuth
func (s *relativesHttpService) handleGetProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		dto, err := s.query.GetMyProfile.Handle(ctx.Request.Context())
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseSuccess(ctx, dto)
	}
}
