package patienthttpservice

import (
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	"github.com/gin-gonic/gin"
)

// @Summary		create patient profile
// @Description	create patient profile
// @Tags			patient
// @Accept			json
// @Produce		json
// @Param			create	form		body					patientcommands.CreatePatientProfileCmdDTO	true	"account creation data"
// @Success		200		{object}	map[string]interface{}	"data"
// @Failure		400		{object}	error					"Bad request error"
// @Router			/api/v1/patients [post]
// @Security		ApiKeyAuth
func (s *patientHttpService) handleCreatePatientProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto patientcommands.CreatePatientProfileCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		req, _ := ctx.Request.Context().Value(common.KeyRequester).(common.Requester)
		fmt.Println("requester in handler: ", req.Role())
		if err := s.cmd.CreatePatientProfile.Handle(ctx.Request.Context(), &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseCreated(ctx)
	}
}
