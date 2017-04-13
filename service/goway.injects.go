package service

import (
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/andrepinto/goway/domain"
)

func (s *gowayImpl) GetInjects(context.Context, *domain.InjectsRequest) (*domain.InjectsResponse, error) {
	var injects_v1 []*domain.InjectDataV1

	injects, err := s.api.GetInjects(); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(injects); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &injects_v1); if err2 != nil {
		return nil, err
	}

	return &domain.InjectsResponse{
		Result: injects_v1,
	}, nil
}

func (s *gowayImpl) GetInject(context context.Context, injectRequest *domain.InjectRequest) (*domain.InjectResponse, error) {
	var injects_v1 *domain.InjectDataV1

	inject, err := s.api.GetInject(injectRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(inject); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &injects_v1); if err2 != nil {
		return nil, err
	}

	return &domain.InjectResponse{
		Result: injects_v1,
	}, nil
}
func (s *gowayImpl) GetInjectByCode(context context.Context, injectRequest *domain.InjectRequest) (*domain.InjectResponse, error) {
	var injects_v1 *domain.InjectDataV1

	inject, err := s.api.GetInjectByCode(injectRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(inject); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &injects_v1); if err2 != nil {
		return nil, err
	}

	return &domain.InjectResponse{
		Result: injects_v1,
	}, nil
}
func (s *gowayImpl) UpdateInject(context context.Context, injectUpdateRequest *domain.InjectUpdateRequest) (*domain.InjectResponse, error) {
	var injectModel *domain.Inject

	injectUpdate, err := json.Marshal(injectUpdateRequest.Inject); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(injectUpdate, &injectModel); if errUpdate != nil {
		return nil, err
	}

	inject, err := s.api.UpdateInject(injectUpdateRequest.Id, injectModel); if err != nil {
		return nil, err
	}

	var injects_v1 *domain.InjectDataV1
	p, err := json.Marshal(inject); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &injects_v1); if err2 != nil {
		return nil, err
	}

	return &domain.InjectResponse{
		Result: injects_v1,
	}, nil
}
func (s *gowayImpl) DeleteInject(context context.Context, injectRequest *domain.InjectRequest) (*domain.DeleteResponse, error) {
	value, err := s.api.DeleteInject(injectRequest.Value); if err != nil {
		return nil, err
	}

	return &domain.DeleteResponse{
		value,
	}, nil
}
