package datastore

import (
	"github.com/barockok/sample_store/models"
	"github.com/russross/meddler"
)

func (db *datastore) CreateProduct(product *models.Product) error {
	return meddler.Insert(db, productTable, product)
}

func (db *datastore) GetProduct(id int64) (*models.Product, error) {
	var product = new(models.Product)
	var err = meddler.Load(db, productTable, product, id)
	return product, err
}

func (db *datastore) UpdateProduct(product *models.Product) error {
	return meddler.Update(db, productTable, product)
}

const productTable = "products"