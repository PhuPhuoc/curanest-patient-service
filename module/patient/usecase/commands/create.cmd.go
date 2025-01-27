package patientcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
)

type createPatientHandler struct {
	cmdRepo PatientCommandRepo
}

func NewCreatePatientHandler(cmdRepo PatientCommandRepo) *createPatientHandler {
	return &createPatientHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *createPatientHandler) Handle(ctx context.Context, dto *PatientProfileCmdDTO) error {
	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	if !ok {
		return common.NewUnauthorizedError()
	}
	sub := requester.UserId()

	patient_id := common.GenUUID()
	entity, _ := patientdomain.NewPatient(
		patient_id,
		sub,
		dto.Gender,
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
	)
	if err := h.cmdRepo.Create(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot create patient profile").
			WithInner(err.Error())
	}

	return nil
}
