package base

import "github.com/undefined7887/telepuz-backend/base/endpoint"

type Service interface {
	GetMethod(name string) endpoint.Method
}
