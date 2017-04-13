package api

import (
	"errors"
	"github.com/andrepinto/goway/domain"
	"github.com/andrepinto/goway/action"
)

func (api *ApiResource) GetRoutes() (interface{}, error) {
	var model domain.Route
	return api.GetAll(model)
}

func (api *ApiResource) GetRoute(id string) (interface{}, error) {
	var model domain.Route
	return api.Get(id, model, []string{"InjectData" })
}

func (api *ApiResource) GetRouteByCode(code string) (interface{}, error) {
	var model domain.Route
	return api.GetByCode(code, model, []string{"InjectData" })
}

func (api *ApiResource) UpdateRoute(id string, value interface{}) (interface{}, error) {
	var model domain.Route
	return api.Update(id, value, model)
}

func (api *ApiResource) DeleteRoute(id string) (bool, error) {
	var model domain.Route
	return api.Delete(id, model)
}

func (api *ApiResource) CreateProductRoute(productId string, value *domain.Route) (interface{}, error) {
	var model domain.Route
	var productModel domain.Product

	if client, errGet := api.Repository.GetById(productId, productModel, nil); errGet != nil {
		return nil, errors.New("productNotExists")
	} else {
		if exists, _ := api.Repository.GetByCode(value.Code, model, nil); exists.(*domain.Route).Code != "" {
			return nil, errors.New("conflict")
		} else {
			value.ReferrerID = client.(*domain.Product).ID
			if err := api.Repository.Create(&value); err != nil {
				return nil, err
			} else {
				if api.ActionEvent != nil {
					api.ActionEvent.AddEvent(action.ADD_MODEL, value)
				}
				return value, nil
			}
		}
	}
}

func (api *ApiResource) CreateClientRoute(clientId string, value *domain.Route) (interface{}, error) {
	var model domain.Route
	var clientModel domain.Client

	if client, errGet := api.Repository.GetById(clientId, clientModel, nil); errGet != nil {
		return nil, errors.New("clientNotExists")
	} else {
		if exists, _ := api.Repository.GetByCode(value.Code, model, nil); exists.(*domain.Route).Code != "" {
			return nil, errors.New("conflict")
		} else {
			value.ReferrerID = client.(*domain.Client).ID
			if err := api.Repository.Create(&value); err != nil {
				return nil, err
			} else {
				if api.ActionEvent != nil {
					api.ActionEvent.AddEvent(action.ADD_MODEL, value)
				}
				return value, nil
			}
		}
	}
}
