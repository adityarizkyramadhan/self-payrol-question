package request

import validation "github.com/go-ozzo/ozzo-validation"

type (
	CompanyRequest struct {
		Name    string `json:"name"`
		Balance int    `json:"balance"`
		Address string `json:"address"`
	}

	TopupCompanyBalance struct {
		Balance int `json:"balance" validate:"required"`
	}
)

// TODO: tuliskan validasi untuk CompanyRequest dengan rule semua field required ✔️

func (req TopupCompanyBalance) Validate() error {
	return validation.ValidateStruct(
		&req,
		validation.Field(&req.Balance, validation.Required),
	)
}

func (com CompanyRequest) Validate() error {
	return validation.ValidateStruct(
		&com,
		validation.Field(&com.Balance, validation.Required),
		validation.Field(&com.Address, validation.Required),
		validation.Field(&com.Name, validation.Required),
	)
}
