package relativescommands

type CreateRelativeAccountCmdDTO struct {
	FullName    string `json:"full-name"`
	Gender      bool   `json:"gender"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
	Dob         string `json:"dob"`
	Address     string `json:"address"`
	Ward        string `json:"ward"`
	District    string `json:"district"`
	City        string `json:"city"`
}

type AccountInfoDTO struct {
	RoleName    string `json:"role-name"`
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Password    string `json:"password"`
}

type UpdateRelativeAccountCmdDTO struct {
	UpdateAccountInfoDTO
	UpdateRelativeInfoDTO
}

type UpdateAccountInfoDTO struct {
	FullName    string `json:"full-name"`
	PhoneNumber string `json:"phone-number"`
	Email       string `json:"email"`
	Avatar      string `json:"avatar"`
}

type UpdateRelativeInfoDTO struct {
	Gender   bool   `json:"gender"`
	Dob      string `json:"dob"`
	Address  string `json:"address"`
	Ward     string `json:"ward"`
	District string `json:"district"`
	City     string `json:"city"`
}
