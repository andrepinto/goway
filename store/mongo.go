package store

import (
	"github.com/andrepinto/goway/domain"
	"gopkg.in/mgo.v2/bson"
	"gopkg.in/mgo.v2"
)

type MongodbRepository struct {
	Session      *mgo.Session
	DatabaseName string
}

type MongodbRepositoryOptions struct {
	Url          string
	DatabaseName string
}

const (
	PRODUCT_COLLECTIONS = "products"
)

func NewMongodbRepository(options *MongodbRepositoryOptions) *MongodbRepository {

	if (len(options.DatabaseName) == 0) {
		panic("NO DATABASE")
	}

	session, err := mgo.Dial(options.Url)

	if err != nil {
		panic(err.Error())
	}

	return &MongodbRepository{
		Session:      session,
		DatabaseName: options.DatabaseName,
	}
}

func (l *MongodbRepository) Create() {

}

func (l *MongodbRepository) GetAllProducts() []domain.ProductV1 {

	var products []domain.ProductV1

	c := l.Session.DB(l.DatabaseName).C(PRODUCT_COLLECTIONS)

	err := c.Find(bson.M{}).All(&products)

	if err != nil {
		panic(err.Error())
	}

	return products
}

func (l *MongodbRepository) GetAllClients() []domain.ClientV1 {
	return nil
}

func (l *MongodbRepository) CreateProduct(product *domain.ProductV1) (bool, *domain.ProductV1) {
	return true, nil
}
func (l *MongodbRepository) CreateClient(client *domain.ClientV1) (bool, *domain.ClientV1) {
	return true, nil
}
