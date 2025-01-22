package patientcommands

import (
	"context"
	"fmt"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
)

type CreatePatientProfileCmdDTO struct {
	FullName      string `json:"full-name"`
	PhoneNumber   string `json:"phone-number"`
	Email         string `json:"email"`
	Dob           string `json:"dob"`
	Address       string `json:"address"`
	Ward          string `json:"ward"`
	District      string `json:"district"`
	City          string `json:"city"`
	DescPathology string `json:"desc-pathology"`
	NoteForNurse  string `json:"note-for-nurse"`
}

type createPatientHandler struct {
	cmdRepo          PatientCommandRepo
	relativesFetcher ReletiveFetcher
}

func NewCreatePatientHandler(cmdRepo PatientCommandRepo, reFetcher ReletiveFetcher) *createPatientHandler {
	return &createPatientHandler{
		cmdRepo:          cmdRepo,
		relativesFetcher: reFetcher,
	}
}

func (h *createPatientHandler) Handle(ctx context.Context, dto *CreatePatientProfileCmdDTO) error {
	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	fmt.Println("requester: ", requester)
	if !ok {
		return common.NewUnauthorizedError()
	}
	sub := requester.UserId()

	patient_id := common.GenUUID()
	entity, _ := patientdomain.NewPatient(
		patient_id,
		sub,
		dto.FullName,
		dto.Dob,
		dto.PhoneNumber,
		dto.Address,
		dto.Ward,
		dto.District,
		dto.City,
		dto.DescPathology,
		dto.NoteForNurse,
		nil,
		nil,
	)
	if err := h.cmdRepo.Create(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create patient profile").
			WithInner(err.Error())
	}

	return nil
}
