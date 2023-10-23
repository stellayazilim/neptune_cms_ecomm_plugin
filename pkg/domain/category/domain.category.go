package domain_category

import "github.com/stellayazilim/neptune_cms_ecomm_plugin/pkg/aggregate"

type ICategoryRepository interface {
	Create(aggregate.Category) error
	GetByID(uint16) aggregate.Category
	GetByName(string) aggregate.Category
	GetAll() []aggregate.Category
	UpdateByID(uint16, aggregate.Category) error
	DeleteByID(uint16) error
}

type CategoryGetByIdRequest struct {
	Params struct {
		ID uint16 `param:"id"`
	}
}

type CategoryResponseDataObject struct {
}

type CategoryResponse struct {
	Body struct {
		Data CategoryResponseDataObject `json:"data"`
	}
}
