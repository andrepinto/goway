package product

import "github.com/andrepinto/goway/domain"

type IProductRepository interface {

	GetAllProducts() []domain.ProductV1
	GetAllClients()	 []domain.ClientV1
	CreateProduct(product *domain.ProductV1) (bool, *domain.ProductV1)
	CreateClient(client *domain.ClientV1) (bool, *domain.ClientV1)

}
