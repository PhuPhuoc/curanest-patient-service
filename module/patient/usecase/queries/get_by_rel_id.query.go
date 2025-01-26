package patientqueries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
)

type getPatientByRelIdHandler struct {
	queryRepo PatientQueryRepo
}

func NewGetPatientByRelIdHandler(queryRepo PatientQueryRepo) *getPatientByRelIdHandler {
	return &getPatientByRelIdHandler{
		queryRepo: queryRepo,
	}
}

func (h *getPatientByRelIdHandler) Handle(ctx context.Context) ([]PatientDTO, error) {
	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	fmt.Println("requester: ", requester)
	if !ok {
		return nil, common.NewUnauthorizedError()
	}
	sub := requester.UserId()

	entities, err := h.queryRepo.GetPatientsByRelId(ctx, sub)
	if err != nil {
		return nil, common.NewInternalServerError().
			WithReason("cannot get list patient of relatives user").
			WithInner(err.Error())
	}

	list_dto := make([]PatientDTO, len(entities))
	for i := range entities {
		list_dto[i] = *toDTO(&entities[i])
	}
	return list_dto, nil
}
