package tests

import (
	"net/http"
	"testing"

	"dance-instructor/internal/requests"
	"dance-instructor/internal/server"
	"dance-instructor/internal/server/handlers"
	"dance-instructor/tests/helpers"

	"github.com/labstack/echo/v4"
	"github.com/stretchr/testify/assert"
)

func TestCreateInstructor(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPost,
		Url:    "/api/v1/instructors",
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers.NewInstructorHandler(s).PostInstructor(c)
	}

	cases := []helpers.TestCase{
		{
			"Post to create new instructor",
			request,
			requests.UpsertInstructorRequest{
				ID:           100,
				Name:         "Felix Le",
				Bio:          "Ballet dancer",
				Specialty:    "Ballet",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 201,
				BodyPart:   `{"message":"Instructor created successfully","data":null}`,
			},
		},
		{
			"Post to create new instructor with empty name",
			request,
			requests.UpsertInstructorRequest{
				ID:           100,
				Name:         "",
				Bio:          "Ballet dancer",
				Specialty:    "Ballet",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   `{"code":400,"error":"name: cannot be blank."}`,
			},
		},
		{
			"Post to create new instructor with wrong specialty",
			request,
			requests.UpsertInstructorRequest{
				ID:           100,
				Name:         "Felix",
				Bio:          "Ballet dancer",
				Specialty:    "Ballesst",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   `{"code":400,"error":"specialty: invalid specialty."}`,
			},
		},
		{
			"Post to create new instructor with empty bio",
			request,
			requests.UpsertInstructorRequest{
				ID:           100,
				Name:         "Felix",
				Bio:          "",
				Specialty:    "Ballet",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   `{"code":400,"error":"bio: cannot be blank."}`,
			},
		},
	}

	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			s := helpers.NewServer(nil)

			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			}
		})
	}
}

func TestUpdateInstructor(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodPut,
		Url:    "/api/v1/instructors/:id",
		PathParam: &helpers.PathParam{
			Name:  "id",
			Value: "1",
		},
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers.NewInstructorHandler(s).PutInstructor(c)
	}

	cases := []helpers.TestCase{
		{
			"Put to update instructor",
			request,
			requests.UpsertInstructorRequest{
				ID:           1,
				Name:         "Felix Le",
				Bio:          "Ballet dancer",
				Specialty:    "Ballet",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   `{"message":"Instructor updated successfully","data":null}`,
			},
		},
		{
			"Put to update instructor with empty name",
			request,
			requests.UpsertInstructorRequest{
				ID:           1,
				Name:         "",
				Bio:          "Ballet dancer",
				Specialty:    "Ballet",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   `{"code":400,"error":"name: cannot be blank."}`,
			},
		},
		{
			"Put to update instructor with wrong specialty",
			request,
			requests.UpsertInstructorRequest{
				ID:           1,
				Name:         "Felix",
				Bio:          "Ballet dancer",
				Specialty:    "Ballesst",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   `{"code":400,"error":"specialty: invalid specialty."}`,
			},
		},
		{
			"Put to update instructor with empty bio",
			request,
			requests.UpsertInstructorRequest{
				ID:           1,
				Name:         "Felix",
				Bio:          "",
				Specialty:    "Ballet",
				Availability: "Monday",
			},
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 400,
				BodyPart:   `{"code":400,"error":"bio: cannot be blank."}`,
			},
		},
	}
	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			s := helpers.NewServer(nil)

			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			}
		})
	}
}

func TestGetInstructor(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/api/v1/instructors/:id",
		PathParam: &helpers.PathParam{
			Name:  "id",
			Value: "1",
		},
	}

	requestNotFound := helpers.Request{
		Method: http.MethodGet,
		Url:    "/api/v1/instructors/:id",
		PathParam: &helpers.PathParam{
			Name:  "id",
			Value: "100",
		},
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers.NewInstructorHandler(s).GetInstructorByID(c)
	}

	cases := []helpers.TestCase{
		{
			"Get instructor",
			request,
			nil,
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   `{"message":"Instructor found","data":{"id":1,"name":"John Doe","bio":"John Doe is a professional dancer with 10 years of experience in Salsa","specialty":"Hip-hop","availability":""}}`,
			},
		},
		{
			"Get instructor with wrong id",
			requestNotFound,
			nil,
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 500,
				BodyPart:   `{"code":500,"error":"item not found"}`,
			},
		},
	}
	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			s := helpers.NewServer(nil)

			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			}
		})
	}
}

func TestGetAllInstructors(t *testing.T) {
	request := helpers.Request{
		Method: http.MethodGet,
		Url:    "/api/v1/instructors",
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers.NewInstructorHandler(s).GetAllInstructors(c)
	}

	cases := []helpers.TestCase{
		{
			"Get all instructors",
			request,
			nil,
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   "{\"total\":3",
			},
		},
	}
	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			s := helpers.NewServer(nil)

			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			}
		})
	}
}

func TestDeleteInstructor(t *testing.T) {
	request1 := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/api/v1/instructors/:id",
		PathParam: &helpers.PathParam{
			Name:  "id",
			Value: "1",
		},
	}
	request2 := helpers.Request{
		Method: http.MethodDelete,
		Url:    "/api/v1/instructors/:id",
		PathParam: &helpers.PathParam{
			Name:  "id",
			Value: "100",
		},
	}
	handlerFunc := func(s *server.Server, c echo.Context) error {
		return handlers.NewInstructorHandler(s).DeleteInstructor(c)
	}

	cases := []helpers.TestCase{
		{
			"Delete instructor",
			request1,
			nil,
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 200,
				BodyPart:   `{"message":"Instructor deleted successfully","data":null}`,
			},
		},
		{
			"Delete instructor with wrong id",
			request2,
			nil,
			handlerFunc,
			nil,
			helpers.ExpectedResponse{
				StatusCode: 500,
				BodyPart:   `{"code":500,"error":"item not found"}`,
			},
		},
	}
	for _, test := range cases {
		t.Run(test.TestName, func(t *testing.T) {
			s := helpers.NewServer(nil)

			c, recorder := helpers.PrepareContextFromTestCase(s, test)

			if assert.NoError(t, test.HandlerFunc(s, c)) {
				assert.Equal(t, test.Expected.StatusCode, recorder.Code)
				assert.Contains(t, recorder.Body.String(), test.Expected.BodyPart)
			}
		})
	}
}
