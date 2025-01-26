package patientcommands

type PatientProfileCmdDTO struct {
	FullName      string `json:"full-name"`
	Gender        bool   `json:"gender"`
	Dob           string `json:"dob"`
	PhoneNumber   string `json:"phone-number"`
	Address       string `json:"address"`
	Ward          string `json:"ward"`
	District      string `json:"district"`
	City          string `json:"city"`
	DescPathology string `json:"desc-pathology"`
	NoteForNurse  string `json:"note-for-nurse"`
}
