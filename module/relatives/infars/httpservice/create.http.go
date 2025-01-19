package relativeshttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	relativescommands "github.com/PhuPhuoc/curanest-patient-service/module/relatives/usecase/commands"
	"github.com/gin-gonic/gin"
)

//	@Summary		create relatives account
//	@Description	create relatives account
//	@Tags			relative
//	@Accept			json
//	@Produce		json
//	@Param			create	form		body					relativescommands.CreateRelativeAccountCmdDTO	true	"account creation data"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/relatives [post]
func (s *relativesHttpService) handleCreateRelativesAccount() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto relativescommands.CreateRelativeAccountCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if err := s.cmd.CreateRelativesAccount.Handle(ctx, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		common.ResponseCreated(ctx)
	}
}
