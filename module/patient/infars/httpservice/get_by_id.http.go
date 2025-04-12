package patienthttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

// @Summary		get patient's profile by patient-id
// @Description	get patient's profile by patient-id
// @Tags			patient
// @Accept			json
// @Produce		json
// @Param			patient-id	path		string					true	"patient ID (UUID)"
// @Success		200			{object}	map[string]interface{}	"data"
// @Failure		400			{object}	error					"Bad request error"
// @Router			/api/v1/patients/{patient-id} [get]
// @Security		ApiKeyAuth
func (s *patientHttpService) handleGetPatientProfile() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		patientId := ctx.Param("patient-id")
		if patientId == "" {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("missing patient-id"))
			return
		}
		patientUUID, err := uuid.Parse(patientId)
		if err != nil {
			common.ResponseError(ctx, common.NewBadRequestError().WithReason("patient invalid (not a uuid)"))
			return
		}

		dto, err := s.query.FindById.Handle(ctx.Request.Context(), &patientUUID)
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, dto)
	}
}
