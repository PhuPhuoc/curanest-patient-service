package patienthttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	"github.com/gin-gonic/gin"
)

//	@Summary		create patient profile
//	@Description	create patient profile
//	@Tags			patient
//	@Accept			json
//	@Produce		json
//	@Param			create	form		body					patientcommands.CreatePatientProfileCmdDTO	true	"account creation data"
//	@Success		200		{object}	map[string]interface{}	"data"
//	@Failure		400		{object}	error					"Bad request error"
//	@Router			/api/v1/patients [post]
func (s *patientHttpService) handleCreatePatientProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto patientcommands.CreatePatientProfileCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if err := s.cmd.CreatePatientProfile.Handle(ctx, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseCreated(ctx)
	}
}
