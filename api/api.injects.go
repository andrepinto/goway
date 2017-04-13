package api

import (
	"errors"
	"github.com/andrepinto/goway/domain"
	"github.com/andrepinto/goway/action"
)

func (api *ApiResource) GetInjects() (interface{}, error) {
	var model domain.Inject
	return api.GetAll(model)
}

func (api *ApiResource) GetInject(id string) (interface{}, error) {
	var model domain.Inject
	return api.Get(id, model, []string{"InjectData" })
}

func (api *ApiResource) GetInjectByCode(code string) (interface{}, error) {
	var model domain.Inject
	return api.GetByCode(code, model, []string{"InjectData" })
}

func (api *ApiResource) UpdateInject(id string, value interface{}) (interface{}, error) {
	var model domain.Inject
	return api.Update(id, value, model)
}

func (api *ApiResource) DeleteInject(id string) (bool, error) {
	var model domain.Inject
	return api.Delete(id, model)
}

func (api *ApiResource) CreateClientInject(clientId string, value *domain.Inject) (interface{}, error) {
	var model domain.Inject
	var clientModel domain.Client

	if client, errGet := api.Repository.GetById(clientId, clientModel, nil); errGet != nil {
		return nil, errors.New("clientNotExists")
	} else {
		if exists, _ := api.Repository.GetByCode(value.Code, model, nil); exists.(*domain.Inject).Code != "" {
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

func (api *ApiResource) CreateRouteInject(routeId string, value *domain.Inject) (interface{}, error) {
	var model domain.Inject
	var routeModel domain.Route

	if route, errGet := api.Repository.GetById(routeId, routeModel, nil); errGet != nil {
		return nil, errors.New("routeNotExists")
	} else {
		if exists, _ := api.Repository.GetByCode(value.Code, model, nil); exists.(*domain.Inject).Code != "" {
			return nil, errors.New("conflict")
		} else {
			value.ReferrerID = route.(*domain.Route).ID
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