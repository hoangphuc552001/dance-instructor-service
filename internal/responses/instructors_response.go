package responses

import (
	"dance-instructor/internal/models"
)

type InstructorResponse struct {
	ID           int    `json:"id"`
	Name         string `json:"name"`
	Bio          string `json:"bio"`
	Specialty    string `json:"specialty"`
	Availability string `json:"availability"`
}

func NewInstructorResponse(instructors []*models.Instructor) []*InstructorResponse {
	instructorResponse := make([]*InstructorResponse, 0)
	for _, instructor := range instructors {
		instructorResponse = append(instructorResponse, &InstructorResponse{
			ID:           instructor.ID,
			Name:         instructor.Name,
			Bio:          instructor.Bio,
			Specialty:    instructor.Specialty,
			Availability: instructor.Availability,
		})
	}
	return instructorResponse
}

func NewInstructorResponseSingle(instructor *models.Instructor) *InstructorResponse {
	return &InstructorResponse{
		ID:           instructor.ID,
		Name:         instructor.Name,
		Bio:          instructor.Bio,
		Specialty:    instructor.Specialty,
		Availability: instructor.Availability,
	}
}
