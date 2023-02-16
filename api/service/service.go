package service

import (
	"net/url"

	"github.com/fignocius/logic-api/api/model"
	uuid "github.com/satori/go.uuid"
)

// ProductService is a interface of service
type LogicService interface {
	List() ([]model.Logic, error)
	UpInsert(model.Logic) (model.Logic, error)
	Apply(uuid.UUID, url.Values) (bool, error)
	Delete(uuid.UUID) error
}
