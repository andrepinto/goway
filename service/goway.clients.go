package service

import (
	"golang.org/x/net/context"
	"encoding/json"
	"github.com/andrepinto/goway/domain"
)


func (s *gowayImpl) GetClients(context.Context, *domain.ClientsRequest) (*domain.ClientsResponse, error) {
	var clients_v1 []*domain.ClientV1

	clients, err := s.api.GetClients(); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(clients); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &clients_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ClientsResponse{
		Result: clients_v1,
	}, nil
}

func (s *gowayImpl) GetClient(context context.Context, clientRequest *domain.ClientRequest) (*domain.ClientResponse, error) {
	var clients_v1 *domain.ClientV1

	client, err := s.api.GetClient(clientRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(client); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &clients_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ClientResponse{
		Result: clients_v1,
	}, nil
}
func (s *gowayImpl) GetClientByCode(context context.Context, clientRequest *domain.ClientRequest) (*domain.ClientResponse, error) {
	var clients_v1 *domain.ClientV1

	client, err := s.api.GetClientByCode(clientRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(client); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &clients_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ClientResponse{
		Result: clients_v1,
	}, nil
}
func (s *gowayImpl) UpdateClient(context context.Context, clientUpdateRequest *domain.ClientUpdateRequest) (*domain.ClientResponse, error) {
	var clientModel *domain.Client

	clientUpdate, err := json.Marshal(clientUpdateRequest.Client); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(clientUpdate, &clientModel); if errUpdate != nil {
		return nil, err
	}

	client, err := s.api.UpdateClient(clientUpdateRequest.Id, clientModel); if err != nil {
		return nil, err
	}

	var clients_v1 *domain.ClientV1
	p, err := json.Marshal(client); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &clients_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ClientResponse{
		Result: clients_v1,
	}, nil
}
func (s *gowayImpl) DeleteClient(context context.Context, clientRequest *domain.ClientRequest) (*domain.DeleteResponse, error) {
	value, err := s.api.DeleteClient(clientRequest.Value); if err != nil {
		return nil, err
	}

	return &domain.DeleteResponse{
		value,
	}, nil
}
func (s *gowayImpl) CreateClient(context context.Context, clientCreateRequest *domain.ClientCreateRequest) (*domain.ClientResponse, error) {
	var clientModel *domain.Client

	clientCreate, err := json.Marshal(clientCreateRequest.Client); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(clientCreate, &clientModel); if errUpdate != nil {
		return nil, err
	}

	client, err := s.api.CreateClient(clientModel); if err != nil {
		return nil, err
	}

	var clients_v1 *domain.ClientV1
	p, err := json.Marshal(client); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &clients_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ClientResponse{
		Result: clients_v1,
	}, nil
}
func (s *gowayImpl) CreateClientRoute(context context.Context, clientCreateRouteRequest *domain.ClientCreateRouteRequest) (*domain.RouteResponse, error) {
	var routeModel *domain.Route

	routeCreate, err := json.Marshal(clientCreateRouteRequest.Route); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(routeCreate, &routeModel); if errUpdate != nil {
		return nil, err
	}

	client, err := s.api.CreateClientRoute(clientCreateRouteRequest.Id, routeModel); if err != nil {
		return nil, err
	}

	var routes_v1 *domain.RoutesV1
	p, err := json.Marshal(client); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &routes_v1); if err2 != nil {
		return nil, err
	}

	return &domain.RouteResponse{
		Result: routes_v1,
	}, nil
}

func (s *gowayImpl) CreateClientInject(context context.Context, clientCreateInjectRequest *domain.ClientCreateInjectRequest) (*domain.InjectResponse, error) {
	var injectModel *domain.Inject

	injectCreate, err := json.Marshal(clientCreateInjectRequest.Inject); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(injectCreate, &injectModel); if errUpdate != nil {
		return nil, err
	}

	inject, err := s.api.CreateClientInject(clientCreateInjectRequest.Id, injectModel); if err != nil {
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