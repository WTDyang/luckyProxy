package svc

import (
	"luckyProxy/client"
)

type ServiceContext struct {
	Config client.Config
}

func NewServiceContext(c client.Config) *ServiceContext {
	return &ServiceContext{
		Config: c,
	}
}
