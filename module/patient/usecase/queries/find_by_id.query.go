package patientqueries

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	"github.com/google/uuid"
)

type findByIdHandler struct {
	queryRepo PatientQueryRepo
}

func NewFindByIdHandler(queryRepo PatientQueryRepo) *findByIdHandler {
	return &findByIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *findByIdHandler) Handle(ctx context.Context, patientId *uuid.UUID) (*PatientDTO, error) {
	entity, err := h.queryRepo.FindByID(ctx, *patientId)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get patient profile of id: " + patientId.String()).
			WithInner(err.Error())
	}

	return toDTO(entity), nil
}
