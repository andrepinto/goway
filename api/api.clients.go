package api

import (
	"errors"

	"github.com/andrepinto/goway/domain"
	"github.com/andrepinto/goway/action"
)

func (api *ApiResource) GetClients() (interface{}, error)  {
	var model domain.Client
	return api.GetAll(model)
}

func (api *ApiResource) GetClient(id string) (interface{}, error)  {
	var model domain.Client
	return api.Get(id, model, []string{ "Routes", "InjectData" })
}

func (api *ApiResource) GetClientByCode(code string) (interface{}, error)  {
	var model domain.Client
	return api.GetByCode(code, model, []string{ "Routes", "InjectData" })
}

func (api *ApiResource) UpdateClient(id string, value interface{}) (interface{}, error)  {
	var model domain.Client
	return api.Update(id, value, model)
}

func (api *ApiResource) DeleteClient(id string) (bool, error) {
	var model domain.Client
	return api.Delete(id, model)
}

func (api *ApiResource) CreateClient(value *domain.Client) (interface{}, error)  {
	var model domain.Client

	if exists, _ := api.Repository.GetByCode(value.Code, model, nil); exists.(*domain.Client).Code != "" {
		return nil, errors.New("conflict")
	} else {
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
