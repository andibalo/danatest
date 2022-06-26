package request

import (
	"fmt"
	"streaming/internal/util"

	"github.com/go-playground/validator/v10"
)

type CreateChatRequest struct {
	Message string `json:"message" validate:"required"`
}

func (r *CreateChatRequest) Validate() error {
	validate := util.GetNewValidator()

	if err := validate.Validate(r); err != nil {
		validationErrorMessage := validate.GetErrorMessage(err.(validator.ValidationErrors))

		for _, err := range err.(validator.ValidationErrors) {
			switch err.Tag() {
			case "alphanumeric":
				validationErrorMessage = fmt.Sprintf("%s has special characters",
					err.Field())

			case "required":
				validationErrorMessage = fmt.Sprintf("%s is required",
					err.Field())

			}
		}

		return fmt.Errorf("%s", validationErrorMessage)
	}

	return nil
}
