package store

import (
	"fmt"
	"encoding/json"

	"github.com/andrepinto/goway/domain"
	"github.com/couchbase/gocb"
	"github.com/twinj/uuid"
)

type CouchbaseRepository struct {
	Cluster *gocb.Cluster
	Bucket  *gocb.Bucket
	options *CouchbaseRepositoryOptions
}

type CouchbaseRepositoryOptions struct {
	ConnectionString string
	BucketName       string
	BucketPassword   string
}

func NewCouchbaseRepository(options *CouchbaseRepositoryOptions) *CouchbaseRepository {

	cluster, err := gocb.Connect(options.ConnectionString)

	if (err != nil) {
		panic("ERROR_ON_OPEN_CLUSTER")
	}

	bucket, err := cluster.OpenBucket(options.BucketName, "")

	if (err != nil) {
		panic("ERROR_ON_OPEN_BUCKET")
	}

	manager := bucket.Manager("", "")

	manager.CreatePrimaryIndex("gateway", true, false)

	manager.CreateIndex("gateway_products", []string{"_type"}, true, false)

	return &CouchbaseRepository{
		Cluster: cluster,
		Bucket: bucket,
		options: options,
	}
}


func(l *CouchbaseRepository) GetAllProducts() []domain.ProductV1 {

	bucketName := l.options.BucketName
	result := []domain.ProductV1{}
	query := gocb.NewN1qlQuery(fmt.Sprintf("SELECT * FROM %s WHERE _type = $1", bucketName))
	rows, _ := l.Bucket.ExecuteN1qlQuery(query, []interface{}{ "product" })

	raw := rows.NextBytes()
	for ( len(raw) > 0 ) {
		res := map[string]domain.ProductV1{}
		json.Unmarshal(raw, &res)
		result = append(result, res[bucketName])
		raw = rows.NextBytes()
	}

	rows.Close()

	return result
}

func(l *CouchbaseRepository) GetAllClients() []domain.ClientV1 {
	bucketName := l.options.BucketName
	result := []domain.ClientV1{}
	query := gocb.NewN1qlQuery(fmt.Sprintf("SELECT * FROM %s WHERE _type = $1", bucketName))
	rows, _ := l.Bucket.ExecuteN1qlQuery(query, []interface{}{ "client" })

	raw := rows.NextBytes()
	for ( len(raw) > 0 ) {
		res := map[string]domain.ClientV1{}
		json.Unmarshal(raw, &res)
		result = append(result, res[bucketName])
		raw = rows.NextBytes()
	}

	rows.Close()

	return result
}

func (l *CouchbaseRepository) CreateProduct(product *domain.ProductV1) (bool, *domain.ProductV1) {

	if ( product.Id == "" ) {
		product.Id = uuid.NewV4().String()
	}

	_, err := l.Bucket.Upsert( fmt.Sprintf("product::%s", product.Id), product, 0 )
	if ( err != nil ) {
		return false, nil
	}

	return true, product
}

func (l *CouchbaseRepository) CreateClient(client *domain.ClientV1) (bool, *domain.ClientV1) {
	if ( client.ApiPath == "" ) {
		client.ApiPath = uuid.NewV4().String()
	}

	_, err := l.Bucket.Upsert( fmt.Sprintf("client::%s", client.ApiPath), client, 0 )
	if ( err != nil ) {
		return false, nil
	}

	return true, client
}
