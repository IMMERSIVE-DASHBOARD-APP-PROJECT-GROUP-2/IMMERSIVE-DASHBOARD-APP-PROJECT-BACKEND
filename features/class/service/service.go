package service

import (
	"errors"

	"github.com/DASHBOARDAPP/features/class"
	"github.com/go-playground/validator/v10"
)

type classService struct {
	classData class.ClassDataInterface
	validate  *validator.Validate
}

// CreateClass implements class.ClassServiceInterface.
func (service *classService) CreateClass(classInput class.Core) error {
	errValidate := service.validate.Struct(classInput)
	if errValidate != nil {
		return errValidate
	}

	// Validasi nama tidak boleh kosong
	if classInput.Name == "" {
		return errors.New("Nama kelas tidak boleh kosong")
	}

	errCreateClass := service.classData.CreateClass(classInput)
	return errCreateClass
}

func New(repo class.ClassDataInterface) class.ClassServiceInterface {
	return &classService{
		classData: repo,
		validate:  validator.New(),
	}
}
