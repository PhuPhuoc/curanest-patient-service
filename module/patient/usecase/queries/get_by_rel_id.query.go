package patientqueries

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

type getPatientByRelIdHandler struct {
	queryRepo PatientQueryRepo
}

func NewGetPatientByRelIdHandler(queryRepo PatientQueryRepo) *getPatientByRelIdHandler {
	return &getPatientByRelIdHandler{
		queryRepo: queryRepo,
	}
}

type PatientDTO struct {
	Id            uuid.UUID `json:"id"`
	FullName      string    `json:"full-name"`
	Dob           string    `json:"dob"`
	PhoneNumber   string    `json:"phone-number"`
	Address       string    `json:"address"`
	Ward          string    `json:"ward"`
	District      string    `json:"district"`
	City          string    `json:"city"`
	DescPathology string    `json:"desc-pathology"`
	NoteForNurse  string    `json:"note-for-nurse"`
}

func toDTO(data *patientdomain.Patient) PatientDTO {
	dto := PatientDTO{
		Id:            data.GetID(),
		FullName:      data.GetFullName(),
		Dob:           data.GetDOB(),
		PhoneNumber:   data.GetPhoneNumber(),
		Address:       data.GetAddress(),
		Ward:          data.GetWard(),
		District:      data.GetDistrict(),
		City:          data.GetCity(),
		DescPathology: data.GetDescPathology(),
		NoteForNurse:  data.GetNoteForNurse(),
	}
	return dto
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
		list_dto[i] = toDTO(&entities[i])
	}
	return list_dto, nil
}
