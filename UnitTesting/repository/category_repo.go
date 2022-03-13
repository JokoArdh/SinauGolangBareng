package repository

import "unit-test-golang/entity"

type CateogoryRepo interface {
	FindById(Id string) *entity.Category
}
