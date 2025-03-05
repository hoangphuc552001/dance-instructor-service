package constants

const (
	AppName = "Dance Instructor"
)

type Specialty string

const (
	Ballet Specialty = "Ballet"
	Jazz   Specialty = "Jazz"
	HipHop Specialty = "Hip-Hop"
)

var Specialties = []Specialty{Ballet, Jazz, HipHop}
