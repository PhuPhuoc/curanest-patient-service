package patientcommands

import (
	"context"

	"github.com/PhuPhuoc/curanest-patient-service/common"
	patientdomain "github.com/PhuPhuoc/curanest-patient-service/module/patient/domain"
	"github.com/google/uuid"
)

type updatePatientHandler struct {
	cmdRepo PatientCommandRepo
}

func NewUpdatePatientHandler(cmdRepo PatientCommandRepo) *updatePatientHandler {
	return &updatePatientHandler{
		cmdRepo: cmdRepo,
	}
}

func (h *updatePatientHandler) Handle(ctx context.Context, patientId *uuid.UUID, dto *PatientProfileCmdDTO) error {
	requester, ok := ctx.Value(common.KeyRequester).(common.Requester)
	if !ok {
		return common.NewUnauthorizedError()
	}
	sub := requester.UserId()

	if patientId == nil {
		return common.NewInternalServerError().
			WithReason("cannot get patient id to update patient profile")
	}

	entity, _ := patientdomain.NewPatient(
		*patientId,
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
		nil,
	)
	if err := h.cmdRepo.Update(ctx, entity); err != nil {
		return common.NewInternalServerError().
			WithReason("cannot update patient profile").
			WithInner(err.Error())
	}

	return nil
}
