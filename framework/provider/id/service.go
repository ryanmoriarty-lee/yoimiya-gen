package id

import (
	"github.com/rs/xid"
)

type YoimiyaIDService struct {
}

func NewYoimiyaIDService(params ...interface{}) (interface{}, error) {
	return &YoimiyaIDService{}, nil
}

func (s *YoimiyaIDService) NewID() string {
	return xid.New().String()
}
