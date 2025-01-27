package patienthttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientcommands "github.com/PhuPhuoc/curanest-patient-service/module/patient/usecase/commands"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		update patient profile
// @Description	update patient profile
// @Tags			patient
// @Accept			json
// @Produce		json
// @Param			patient-id	path		string					true									"Patient ID (UUID)"
// @Param			update		form		body					patientcommands.PatientProfileCmdDTO	true	"account creation data"
// @Success		201			{object}	map[string]interface{}	"data"
// @Failure		401			{object}	error					"Bad request error"
// @Router			/api/v1/patients/{patient-id} [put]
// @Security		ApiKeyAuth
func (s *patientHttpService) handleUpdatePatientProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var dto patientcommands.PatientProfileCmdDTO
		if err := ctx.BindJSON(&dto); err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("invalid request body").WithInner(err.Error()))
			return
		}

		patientId := ctx.Param("patient-id")
		if patientId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing patient-id"))
			return
		}

		patientUUID, err := uuid.Parse(patientId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("id invalid (not a uuid)"))
			return
		}

		requester, ok := ctx.Request.Context().Value(common.KeyRequester).(common.Requester)
		if !ok {
			common.ResponseError(ctx, common.NewUnauthorizedError().WithReason("cannot found requester"))
			return
		}
		sub := requester.UserId()

		patient, err := s.query.FindById.Handle(ctx, &patientUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}

		if patient.RelativesId != sub {
			common.ResponseError(ctx, common.NewForbiddenError().WithReason("this patient is not your relative"))
			return
		}

		if err := s.cmd.UpdatePatientProfile.Handle(ctx.Request.Context(), &patientUUID, &dto); err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseUpdated(ctx)
	}
}
