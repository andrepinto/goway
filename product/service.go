package product

import "github.com/andrepinto/goway/domain"

func(r *ProductResource) GetAllProducts() []domain.ProductV1{
	return r.GetRepository().GetAllProducts()
}

func(r *ProductResource) GetAllClients() []domain.ClientV1{
	return r.GetRepository().GetAllClients()
}