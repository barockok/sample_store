package datastore

import (
	"testing"
	"github.com/barockok/sample_store/models"
	"github.com/franela/goblin"
)

func TestProducts(t *testing.T) {
	db := openTest()
	defer db.Close()
	s := From(db)

	g := goblin.Goblin(t)
	g.Describe("Products", func() {
		g.BeforeEach(func() {
			db.Exec("DELETE FROM products")
		})

		g.It("Should create", func() {
			product := models.Product{
        Price: 12000,
        Name: "Sample Product 1",
        Description: "Sample Product Description 1",
			}

			err := s.CreateProduct(&product)
			g.Assert(err == nil).IsTrue()
			g.Assert(product.ID != 0).IsTrue()
		})

    g.It("Should get", func ()  {
  		product := models.Product{
        Price: 12000,
        Name: "Sample Product 1",
        Description: "Sample Product Description 1",
  		}

  		s.CreateProduct(&product)
      getProduct, err := s.GetProduct(product.ID)

      g.Assert(err== nil).IsTrue()
      g.Assert(getProduct.Name).Equal(product.Name)
      g.Assert(getProduct.Price).Equal(product.Price)
      g.Assert(getProduct.Description).Equal(product.Description)
    })

    g.It("Should update", func ()  {
  		product := models.Product{
        Price: 12000,
        Name: "Sample Product 1",
        Description: "Sample Product Description 1",
  		}

  		s.CreateProduct(&product)
      product.Name = "Updated Product Name"
      updateErr := s.UpdateProduct(&product)
      getProduct, _ := s.GetProduct(product.ID)

      g.Assert(updateErr== nil).IsTrue()
      g.Assert(getProduct.Name).Equal("Updated Product Name")
    })
  })
}
