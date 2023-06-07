package service

import (
	"github.com/DASHBOARDAPP/features/mentee"

	"github.com/go-playground/validator/v10"
)

type menteeService struct {
	menteeData mentee.MenteeDataInterface
	validate   *validator.Validate
}

// CreateMentee implements mentee.MenteeServiceInterface.
func (service *menteeService) CreateMentee(menteeInput mentee.Core) error {
	errValidate := service.validate.Struct(menteeInput)
	if errValidate != nil {
		return errValidate
	}

	errCreateMentee := service.menteeData.CreateMentee(menteeInput)
	return errCreateMentee
}

// GetAllMentee implements mentee.MenteeServiceInterface.
func (service *menteeService) GetAllMentee() ([]mentee.Core, error) {
	panic("unimplemented")
}

func New(repo mentee.MenteeDataInterface) mentee.MenteeServiceInterface {
	return &menteeService{
		menteeData: repo,
		validate:   validator.New(),
	}
}
