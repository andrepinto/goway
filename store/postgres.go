package store

import (
	"encoding/json"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/andrepinto/goway/domain"
)

type PostgresRepository struct {
	DB *gorm.DB
}

type PostgresRepositoryOptions struct {
	ConnectionString string
}

func NewPostgresRepository(options *PostgresRepositoryOptions) *PostgresRepository {
	db, err := gorm.Open("postgres", options.ConnectionString)

	if err != nil {
		panic("failed to connect database")
	}

	db.LogMode(true)

	return &PostgresRepository{
		DB: db,
	}
}

func (rp *PostgresRepository) GetAllProducts() []domain.ProductV1 {
	var products []domain.Product
	if err := rp.DB.Preload("Routes.InjectData").Preload("Routes").Find(&products).Error; err != nil {
		return nil
	} else {
		var products_v1 []domain.ProductV1

		p, err := json.Marshal(products)
		if err != nil {
			return nil
		}

		err2 := json.Unmarshal(p, &products_v1)
		if err2 != nil {
			return nil
		}
		return products_v1
	}
	return nil
}

func (rp *PostgresRepository) GetAllClients() []domain.ClientV1 {
	var clients []domain.Client
	if err := rp.DB.
		Preload("InjectData", func(db *gorm.DB) *gorm.DB {
		return db.Order("injects.referrer_id").Order("injects.order ASC")
	}).
		Preload("Routes.InjectData", func(db *gorm.DB) *gorm.DB {
		return db.Order("injects.referrer_id").Order("injects.order ASC")
	}).
		Preload("Routes").
		Find(&clients).Error; err != nil {
		return nil
	} else {
		var clients_v1 []domain.ClientV1

		c, err := json.Marshal(clients)
		if err != nil {
			return nil
		}

		err2 := json.Unmarshal(c, &clients_v1)
		if err2 != nil {
			return nil
		}
		return clients_v1
	}
	return nil
}

func (rp *PostgresRepository) CreateProduct(product *domain.ProductV1) (bool, *domain.ProductV1) {
	return true, nil
}

func (rp *PostgresRepository) CreateClient(client *domain.ClientV1) (bool, *domain.ClientV1) {
	return true, nil
}
