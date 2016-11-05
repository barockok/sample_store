package store

import (
  "github.com/barockok/sample_store/models"
)

type Store interface {
  CreateProduct(*models.Product) error
  GetProduct(id int64) (*models.Product, error)
  UpdateProduct(*models.Product) error
}