package service

import (
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/andrepinto/goway/domain"
)

func (s *gowayImpl) GetRoutes(context.Context, *domain.RoutesRequest) (*domain.RoutesResponse, error) {
	var routes_v1 []*domain.RoutesV1

	routes, err := s.api.GetRoutes(); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(routes); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &routes_v1); if err2 != nil {
		return nil, err
	}

	return &domain.RoutesResponse{
		Result: routes_v1,
	}, nil
}

func (s *gowayImpl) GetRoute(context context.Context, routeRequest *domain.RouteRequest) (*domain.RouteResponse, error) {
	var routes_v1 *domain.RoutesV1

	route, err := s.api.GetRoute(routeRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(route); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &routes_v1); if err2 != nil {
		return nil, err
	}

	return &domain.RouteResponse{
		Result: routes_v1,
	}, nil
}
func (s *gowayImpl) GetRouteByCode(context context.Context, routeRequest *domain.RouteRequest) (*domain.RouteResponse, error) {
	var routes_v1 *domain.RoutesV1

	route, err := s.api.GetRouteByCode(routeRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(route); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &routes_v1); if err2 != nil {
		return nil, err
	}

	return &domain.RouteResponse{
		Result: routes_v1,
	}, nil
}
func (s *gowayImpl) UpdateRoute(context context.Context, routeUpdateRequest *domain.RouteUpdateRequest) (*domain.RouteResponse, error) {
	var routeModel *domain.Route

	routeUpdate, err := json.Marshal(routeUpdateRequest.Route); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(routeUpdate, &routeModel); if errUpdate != nil {
		return nil, err
	}

	route, err := s.api.UpdateRoute(routeUpdateRequest.Id, routeModel); if err != nil {
		return nil, err
	}

	var routes_v1 *domain.RoutesV1
	p, err := json.Marshal(route); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &routes_v1); if err2 != nil {
		return nil, err
	}

	return &domain.RouteResponse{
		Result: routes_v1,
	}, nil
}
func (s *gowayImpl) DeleteRoute(context context.Context, routeRequest *domain.RouteRequest) (*domain.DeleteResponse, error) {
	value, err := s.api.DeleteRoute(routeRequest.Value); if err != nil {
		return nil, err
	}

	return &domain.DeleteResponse{
		value,
	}, nil
}

func (s *gowayImpl) CreateRouteInject(context context.Context, routeCreateInjectRequest *domain.RouteCreateInjectRequest) (*domain.InjectResponse, error) {
	var injectModel *domain.Inject

	injectCreate, err := json.Marshal(routeCreateInjectRequest.Inject); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(injectCreate, &injectModel); if errUpdate != nil {
		return nil, err
	}

	inject, err := s.api.CreateRouteInject(routeCreateInjectRequest.Id, injectModel); if err != nil {
		return nil, err
	}

	var injectData_v1 *domain.InjectDataV1
	p, err := json.Marshal(inject); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &injectData_v1); if err2 != nil {
		return nil, err
	}

	return &domain.InjectResponse{
		Result: injectData_v1,
	}, nil
}
