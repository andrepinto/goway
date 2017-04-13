package service

import (
	"golang.org/x/net/context"

	"github.com/andrepinto/goway/api"
	"github.com/andrepinto/goway/domain"
)

type gowayImpl struct {
	api *api.ApiResource
}

func (s *gowayImpl) Version(context.Context, *domain.VersionRequest) (*domain.VersionResponse, error) {
	return &domain.VersionResponse{
		Value: "v1.0",
	}, nil
}

func NewGowayImpl(api *api.ApiResource) *gowayImpl {
	return &gowayImpl{
		api,
	}
}
