package patienthttpservice

import (
	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/gin-gonic/gin"
)

// @Summary		get all patients belong to relatives user
// @Description	get all patients belong to relatives user
// @Tags			patient
// @Accept			json
// @Produce		json
// @Success		200	{object}	map[string]interface{}	"data"
// @Failure		400	{object}	error					"Bad request error"
// @Router			/api/v1/patients/relatives [get]
// @Security		ApiKeyAuth
func (s *patientHttpService) handleGetMyPatient() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		result, err := s.query.GetPatientByRelativesId.Handle(ctx.Request.Context())
		if err != nil {
			common.ResponseError(ctx, err)
			return
		}
		common.ResponseSuccess(ctx, result)
	}
}
