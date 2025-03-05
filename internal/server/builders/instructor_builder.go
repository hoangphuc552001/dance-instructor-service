package builders

import (
	"dance-instructor/internal/models"
)

type InstructorBuilder struct {
	ID           int
	Name         string
	Bio          string
	Specialty    string
	Availability string
}

func NewInstructorBuilder() *InstructorBuilder {
	return &InstructorBuilder{}
}

func (instructorBuilder *InstructorBuilder) SetID(id int) (u *InstructorBuilder) {
	instructorBuilder.ID = id
	return instructorBuilder
}

func (instructorBuilder *InstructorBuilder) SetName(name string) (u *InstructorBuilder) {
	instructorBuilder.Name = name
	return instructorBuilder
}

func (instructorBuilder *InstructorBuilder) SetBio(bio string) (u *InstructorBuilder) {
	instructorBuilder.Bio = bio
	return instructorBuilder
}

func (instructorBuilder *InstructorBuilder) SetSpecialty(specialty string) (u *InstructorBuilder) {
	instructorBuilder.Specialty = specialty
	return instructorBuilder
}

func (instructorBuilder *InstructorBuilder) SetAvailability(availability string) (u *InstructorBuilder) {
	instructorBuilder.Availability = availability
	return instructorBuilder
}

func (instructorBuilder *InstructorBuilder) Build() models.Instructor {
	instructor := models.Instructor{
		ID:           instructorBuilder.ID,
		Name:         instructorBuilder.Name,
		Bio:          instructorBuilder.Bio,
		Specialty:    instructorBuilder.Specialty,
		Availability: instructorBuilder.Availability,
	}

	return instructor
}
