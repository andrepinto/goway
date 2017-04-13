package api

import (
	"errors"
	"github.com/andrepinto/goway/domain"
	"github.com/andrepinto/goway/action"
)

func (api *ApiResource) GetProducts() (interface{}, error) {
	var model domain.Product
	return api.GetAll(model)
}

func (api *ApiResource) GetProduct(id string) (interface{}, error) {
	var model domain.Product
	return api.Get(id, model, []string{"Routes" })
}

func (api *ApiResource) GetProductByCode(code string) (interface{}, error) {
	var model domain.Product
	return api.GetByCode(code, model, []string{"Routes" })
}

func (api *ApiResource) UpdateProduct(id string, value interface{}) (interface{}, error) {
	var model domain.Product
	return api.Update(id, value, model)
}

func (api *ApiResource) DeleteProduct(id string) (bool, error) {
	var model domain.Product
	return api.Delete(id, model)
}

func (api *ApiResource) CreateProduct(value *domain.Product) (interface{}, error) {
	var model domain.Product

	if exists, _ := api.Repository.GetByCode(value.Code, model, nil); exists.(*domain.Product).Code != "" {
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
