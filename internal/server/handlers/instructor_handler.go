package handlers

import (
	"net/http"
	"strconv"

	"dance-instructor/internal/repositories"
	"dance-instructor/internal/requests"
	"dance-instructor/internal/responses"
	"dance-instructor/internal/server"
	"dance-instructor/internal/services/instructor"

	"github.com/labstack/echo/v4"
)

type InstructorHandlers struct {
	server *server.Server

	instructorService instructor.IInstructorService
}

func NewInstructorHandler(server *server.Server) *InstructorHandlers {
	return &InstructorHandlers{
		server:            server,
		instructorService: instructor.NewInstructorService(repositories.NewInstructorRepository()),
	}
}

// GetAllInstructors @Summary Get all instructors
// @Description Get all instructors
// @Tags instructors
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Success 200 {object} responses.ResponseDataForGetList
// @Failure 500 {object} responses.ResponseData
// @Failure 400 {object} responses.ResponseData
// @Router /api/v1/instructors [get]
func (userHandlers *InstructorHandlers) GetAllInstructors(c echo.Context) error {
	instructors, err := userHandlers.instructorService.GetAllInstructors()
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	instructorsResp := responses.NewInstructorResponse(instructors)

	return responses.Response(c, http.StatusOK, &responses.ResponseDataForGetList{
		List:  instructorsResp,
		Total: int64(len(instructorsResp)),
	})
}

// GetInstructorByID @Summary Get instructor by ID
// @Description Get instructor by ID
// @Tags instructors
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "Instructor ID"
// @Success 200 {object} responses.ResponseData
// @Failure 500 {object} responses.ResponseData
// @Failure 400 {object} responses.ResponseData
// @Router /api/v1/instructors/{id} [get]
func (userHandlers *InstructorHandlers) GetInstructorByID(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "ID is required")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}
	instructor, err := userHandlers.instructorService.GetInstructorByID(idInt)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	instructorResp := responses.NewInstructorResponseSingle(instructor)

	return responses.Response(c, http.StatusOK, &responses.ResponseData{
		Data:    instructorResp,
		Message: "Instructor found",
	})
}

// PostInstructor @Summary Create instructor
// @Description Create instructor
// @Tags instructors
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param body body requests.UpsertInstructorRequest true "Instructor data"
// @Success 201 {object} responses.ResponseData
// @Failure 500 {object} responses.ResponseData
// @Failure 400 {object} responses.ResponseData
// @Router /api/v1/instructors [post]
func (userHandlers *InstructorHandlers) PostInstructor(c echo.Context) error {
	instructorRequest := new(requests.UpsertInstructorRequest)
	if err := c.Bind(instructorRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if errIsValidate := instructorRequest.Validate(); errIsValidate != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, errIsValidate.Error())
	}

	err := userHandlers.instructorService.CreateInstructor(instructorRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusCreated, &responses.ResponseData{
		Message: "Instructor created successfully",
	})
}

// PutInstructor @Summary Update instructor
// @Description Update instructor
// @Tags instructors
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "Instructor ID"
// @Param body body requests.UpsertInstructorRequest true "Instructor data"
// @Success 200 {object} responses.ResponseData
// @Failure 500 {object} responses.ResponseData
// @Failure 400 {object} responses.ResponseData
// @Router /api/v1/instructors/{id} [put]
func (userHandlers *InstructorHandlers) PutInstructor(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "ID is required")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	instructorRequest := new(requests.UpsertInstructorRequest)
	if err := c.Bind(instructorRequest); err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, err.Error())
	} else if errIsValidate := instructorRequest.Validate(); errIsValidate != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, errIsValidate.Error())
	}
	instructorRequest.ID = idInt

	err = userHandlers.instructorService.UpdateInstructor(instructorRequest)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, &responses.ResponseData{
		Message: "Instructor updated successfully",
	})
}

// DeleteInstructor @Summary Delete instructor
// @Description Delete instructor
// @Tags instructors
// @Accept json
// @Produce json
// @Param Authorization header string true "token"
// @Param id path int true "Instructor ID"
// @Success 200 {object} responses.ResponseData
// @Failure 500 {object} responses.ResponseData
// @Failure 400 {object} responses.ResponseData
// @Router /api/v1/instructors/{id} [delete]
func (userHandlers *InstructorHandlers) DeleteInstructor(c echo.Context) error {
	id := c.Param("id")
	if id == "" {
		return responses.ErrorResponse(c, http.StatusBadRequest, "ID is required")
	}
	idInt, err := strconv.Atoi(id)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusBadRequest, "ID must be a number")
	}

	err = userHandlers.instructorService.DeleteInstructor(idInt)
	if err != nil {
		return responses.ErrorResponse(c, http.StatusInternalServerError, err.Error())
	}

	return responses.Response(c, http.StatusOK, &responses.ResponseData{
		Message: "Instructor deleted successfully",
	})
}
