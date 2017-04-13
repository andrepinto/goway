package service

import (
	"encoding/json"
	"golang.org/x/net/context"
	"github.com/andrepinto/goway/domain"
)

func (s *gowayImpl) GetProducts(context.Context, *domain.ProductsRequest) (*domain.ProductsResponse, error) {
	var products_v1 []*domain.ProductV1

	products, err := s.api.GetProducts(); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(products); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &products_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ProductsResponse{
		Result: products_v1,
	}, nil
}

func (s *gowayImpl) GetProduct(context context.Context, productRequest *domain.ProductRequest) (*domain.ProductResponse, error) {
	var products_v1 *domain.ProductV1

	product, err := s.api.GetProduct(productRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(product); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &products_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ProductResponse{
		Result: products_v1,
	}, nil
}
func (s *gowayImpl) GetProductByCode(context context.Context, productRequest *domain.ProductRequest) (*domain.ProductResponse, error) {
	var products_v1 *domain.ProductV1

	product, err := s.api.GetProductByCode(productRequest.Value); if err != nil {
		return nil, err
	}

	p, err := json.Marshal(product); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &products_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ProductResponse{
		Result: products_v1,
	}, nil
}
func (s *gowayImpl) UpdateProduct(context context.Context, productUpdateRequest *domain.ProductUpdateRequest) (*domain.ProductResponse, error) {
	var productModel *domain.Product

	productUpdate, err := json.Marshal(productUpdateRequest.Product); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(productUpdate, &productModel); if errUpdate != nil {
		return nil, err
	}

	product, err := s.api.UpdateProduct(productUpdateRequest.Id, productModel); if err != nil {
		return nil, err
	}

	var products_v1 *domain.ProductV1
	p, err := json.Marshal(product); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &products_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ProductResponse{
		Result: products_v1,
	}, nil
}
func (s *gowayImpl) DeleteProduct(context context.Context, productRequest *domain.ProductRequest) (*domain.DeleteResponse, error) {
	value, err := s.api.DeleteProduct(productRequest.Value); if err != nil {
		return nil, err
	}

	return &domain.DeleteResponse{
		value,
	}, nil
}
func (s *gowayImpl) CreateProduct(context context.Context, productCreateRequest *domain.ProductCreateRequest) (*domain.ProductResponse, error) {
	var productModel *domain.Product

	productCreate, err := json.Marshal(productCreateRequest.Product); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(productCreate, &productModel); if errUpdate != nil {
		return nil, err
	}

	product, err := s.api.CreateProduct(productModel); if err != nil {
		return nil, err
	}

	var products_v1 *domain.ProductV1
	p, err := json.Marshal(product); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &products_v1); if err2 != nil {
		return nil, err
	}

	return &domain.ProductResponse{
		Result: products_v1,
	}, nil
}
func (s *gowayImpl) CreateProductRoute(context context.Context, productCreateRouteRequest *domain.ProductCreateRouteRequest) (*domain.RouteResponse, error) {
	var routeModel *domain.Route

	routeCreate, err := json.Marshal(productCreateRouteRequest.Route); if err != nil {
		return nil, err
	}
	errUpdate := json.Unmarshal(routeCreate, &routeModel); if errUpdate != nil {
		return nil, err
	}

	product, err := s.api.CreateProductRoute(productCreateRouteRequest.Id, routeModel); if err != nil {
		return nil, err
	}

	var routes_v1 *domain.RoutesV1
	p, err := json.Marshal(product); if err != nil {
		return nil, err
	}
	err2 := json.Unmarshal(p, &routes_v1); if err2 != nil {
		return nil, err
	}

	return &domain.RouteResponse{
		Result: routes_v1,
	}, nil
}
