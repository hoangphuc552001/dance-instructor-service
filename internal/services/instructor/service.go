package instructor

import (
	"dance-instructor/internal/models"
	"dance-instructor/internal/repositories"
	"dance-instructor/internal/requests"
)

type IInstructorService interface {
	GetInstructorByID(id int) (*models.Instructor, error)
	GetAllInstructors() ([]*models.Instructor, error)
	CreateInstructor(req *requests.UpsertInstructorRequest) error
	UpdateInstructor(req *requests.UpsertInstructorRequest) error
	DeleteInstructor(id int) error
}

type instructorService struct {
	instructorRepo repositories.InstructorRepository
}

func NewInstructorService(instructorRepo repositories.InstructorRepository) IInstructorService {
	return &instructorService{
		instructorRepo: instructorRepo,
	}
}

func (i *instructorService) GetInstructorByID(id int) (*models.Instructor, error) {
	return i.instructorRepo.GetInstructorByID(id)
}

func (i *instructorService) GetAllInstructors() ([]*models.Instructor, error) {
	return i.instructorRepo.GetAllInstructors()
}

func (i *instructorService) CreateInstructor(req *requests.UpsertInstructorRequest) error {
	return i.instructorRepo.CreateInstructor(req.ToInstructor())
}

func (i *instructorService) UpdateInstructor(req *requests.UpsertInstructorRequest) error {
	instructor, err := i.GetInstructorByID(req.ID)
	if err != nil {
		return err
	}
	instructor.Name = req.Name
	instructor.Bio = req.Bio
	instructor.Specialty = req.Specialty
	instructor.Availability = req.Availability
	return i.instructorRepo.UpdateInstructor(instructor)
}

func (i *instructorService) DeleteInstructor(id int) error {
	return i.instructorRepo.DeleteInstructor(id)
}
