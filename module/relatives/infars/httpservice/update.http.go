package relativeshttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		create relatives account
// @Description	create relatives account
// @Tags			relative
// @Accept			json
// @Produce		json
// @Param			relatives-id	path		string					true											"Account ID (UUID)"
// @Param			create			form		body					relativescommands.UpdateRelativeAccountCmdDTO	true	"account creation data"
// @Success		200				{object}	map[string]interface{}	"data"
// @Failure		400				{object}	error					"Bad request error"
// @Router			/api/v1/relatives/{relatives-id} [put]
// @Security		ApiKeyAuth
func (s *relativesHttpService) handleUpdateRelativesAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto relativescommands.UpdateRelativeAccountCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		relativesId := ctx.Param("relatives-id")
		if relativesId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing patient-id"))
			return
		}

		relativesUUID, err := uuid.Parse(relativesId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("id invalid (not a uuid)"))
			return
		}

		if err := s.cmd.UpdateRelativesAccount.Handle(ctx.Request.Context(), &relativesUUID, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseUpdated(ctx)
	}
}
