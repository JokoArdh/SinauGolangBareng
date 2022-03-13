package repository

import (
	"unit-test-golang/entity"

	"github.com/stretchr/testify/mock"
)

type RepositoryMock struct {
	Mock mock.Mock
}

func (repository *RepositoryMock) FindById(Id string) *entity.Category {

	arguments := repository.Mock.Called(Id)

	if arguments.Get(0) == nil {
		return nil
	} else {
		category := arguments.Get(0).(entity.Category)
		return &category
	}

}
