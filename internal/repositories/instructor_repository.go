package repositories

import (
	"dance-instructor/internal/db"
	"dance-instructor/internal/models"
)

type InstructorRepository interface {
	GetInstructorByID(id int) (*models.Instructor, error)
	GetAllInstructors() ([]*models.Instructor, error)
	CreateInstructor(instructor *models.Instructor) error
	UpdateInstructor(instructor *models.Instructor) error
	DeleteInstructor(id int) error
}
type instructorRepository struct {
	db *db.DB[*models.Instructor]
}

func NewInstructorRepository() InstructorRepository {
	db := db.NewDB[*models.Instructor]()
	// Adding fake data
	db.Create(1, &models.Instructor{ID: 1, Name: "John Doe", Bio: "John Doe is a professional dancer with 10 years of experience in Salsa", Specialty: "Hip-hop"})
	db.Create(2, &models.Instructor{ID: 2, Name: "Jane Smith", Bio: "Jane Smith is a professional dancer with 5 years of experience in Salsa", Specialty: "Hip-hop"})
	db.Create(3, &models.Instructor{ID: 3, Name: "Mike Johnson", Bio: "Mike Johnson is a professional dancer with 15 years of experience in Salsa", Specialty: "Hip-hop"})
	return &instructorRepository{
		db: db,
	}
}

func (i *instructorRepository) GetInstructorByID(id int) (*models.Instructor, error) {
	return i.db.GetByID(id)
}

func (i *instructorRepository) GetAllInstructors() ([]*models.Instructor, error) {
	return i.db.GetAll()
}

func (i *instructorRepository) CreateInstructor(instructor *models.Instructor) error {
	return i.db.Create(instructor.ID, instructor)
}

func (i *instructorRepository) UpdateInstructor(instructor *models.Instructor) error {
	return i.db.Update(instructor.ID, instructor)
}

func (i *instructorRepository) DeleteInstructor(id int) error {
	return i.db.Delete(id)
}
