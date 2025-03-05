package requests

import (
	"errors"
	"strings"

	"dance-instructor/internal/constants"
	"dance-instructor/internal/models"

	validation "github.com/go-ozzo/ozzo-validation/v4"
)

type UpsertInstructorRequest struct {
	ID           int    `json:"id" validate:"required"`
	Name         string `json:"name" validate:"required"`
	Bio          string `json:"bio" validate:"required"`
	Specialty    string `json:"specialty"`
	Availability string `json:"availability"`
}

func (i *UpsertInstructorRequest) Validate() error {
	isValidSpecialty := false
	for _, allowed := range constants.Specialties {
		if strings.EqualFold(string(allowed), i.Specialty) {
			isValidSpecialty = true
			break
		}
	}
	if !isValidSpecialty {
		return validation.Errors{"specialty": errors.New("invalid specialty")}
	}
	return validation.ValidateStruct(i,
		validation.Field(&i.ID, validation.Required),
		validation.Field(&i.Name, validation.Required),
		validation.Field(&i.Bio, validation.Required))
}

func (i *UpsertInstructorRequest) ToInstructor() *models.Instructor {
	return &models.Instructor{
		ID:           i.ID,
		Name:         i.Name,
		Bio:          i.Bio,
		Specialty:    i.Specialty,
		Availability: i.Availability,
	}
}
