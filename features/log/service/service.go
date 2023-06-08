package service

import (
	"errors"

	"github.com/DASHBOARDAPP/features/log"
	"github.com/DASHBOARDAPP/features/mentee"
	"github.com/DASHBOARDAPP/features/user"
	"github.com/go-playground/validator/v10"
)

type logService struct {
	logData    log.LogDataInterface
	menteeData mentee.MenteeDataInterface
	userData   user.UserDataInterface
	validate   *validator.Validate
}

// GetLogsByID implements log.LogServiceInterface.
func (*logService) GetLogsByID(logID uint) ([]log.Core, error) {
	panic("unimplemented")
}

func (service *logService) GetLogsByMenteeID(menteeID uint) ([]log.Core, error) {
	_, err := service.menteeData.GetMenteeByID(menteeID)
	if err != nil {
		// Handle error
		return nil, err
	}

	logs, err := service.logData.GetLogsByMenteeID(menteeID)
	if err != nil {
		// Handle error
		return nil, err
	}

	return logs, nil
}

// CreateLog implements log.LogServiceInterface.
func (service *logService) Insert(logInput log.Core) error {
	errValidate := service.validate.Struct(logInput)
	if errValidate != nil {
		return errValidate
	}
	// Validasi nama tidak boleh kosong
	if logInput.Description == "" {
		return errors.New("description tidak boleh kosong")
	}
	// Validasi nama tidak boleh kosong
	if logInput.MenteeID == 0 {
		return errors.New("id tidak boleh kosong")
	}
	if logInput.UserID == 0 {
		return errors.New("id tidak boleh kosong")
	}

	err := service.logData.Create(logInput)
	if err != nil {
		// Handle error
		return err
	}

	return nil
}

func New(repo log.LogDataInterface) log.LogServiceInterface {
	return &logService{
		logData:  repo,
		validate: validator.New(),
	}
}
