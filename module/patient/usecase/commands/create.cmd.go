package patientcommands

import (
	"context"
)

type CreatePatientProfileCmdDTO struct {
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Ward        string `json:"ward"`
	District    string `json:"district"`
	City        string `json:"city"`
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
	return nil
}
