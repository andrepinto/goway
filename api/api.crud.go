package api

import (
	"github.com/imdario/mergo"
	"github.com/andrepinto/goway/action"
)

func (api *ApiResource) GetAll(model interface{}) (interface{}, error) {
	return api.Repository.GetAll(model)
}

func (api *ApiResource) Get(id string, model interface{}, preload []string) (interface{}, error) {
	return api.Repository.GetById(id, model, preload)
}

func (api *ApiResource) GetByCode(code string, model interface{}, preload []string) (interface{}, error) {
	return api.Repository.GetByCode(code, model, preload)
}

func (api *ApiResource) Update(id string, value interface{}, model interface{}) (interface{}, error) {
	if product, err := api.Repository.GetById(id, model, nil); err != nil {
		return nil, err
	} else {
		mergo.Merge(value, product)
		if err := api.Repository.Update(id, value); err != nil {
			return nil, err
		} else {
			if result, err := api.Repository.GetById(id, model, nil); err != nil {
				return nil, err
			} else {
				if api.ActionEvent != nil {
					api.ActionEvent.AddEvent(action.UPDATE_MODEL, result)
				}
				return result, nil
			}
		}
	}
}

func (api *ApiResource) Delete(id string, model interface{}) (bool, error) {
	if err := api.Repository.Delete(id, model); err != nil {
		return false, err
	} else {
		if api.ActionEvent != nil {
			api.ActionEvent.AddEvent(action.DELETE_MODEL, id)
		}
		return true, nil
	}
}
