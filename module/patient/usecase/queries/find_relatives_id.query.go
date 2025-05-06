package patientqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/google/uuid"
)

type findRelativesIdHandler struct {
	queryRepo PatientQueryRepo
}

func NewFindRelatviesIdHandler(queryRepo PatientQueryRepo) *findRelativesIdHandler {
	return &findRelativesIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findRelativesIdHandler) Handle(ctx context.Context, patientId *uuid.UUID) (*ResponseRelativesIdOfPatient, error) {
	entity, err := h.queryRepo.FindByID(ctx, *patientId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get patient profile of id: " + patientId.String()).
			WithInner(err.Error())
	}

	resp := ResponseRelativesIdOfPatient{
		RelativesId: entity.GetRelativesID(),
	}
	return &resp, nil
}
