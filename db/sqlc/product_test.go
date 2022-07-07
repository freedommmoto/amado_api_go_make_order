package db

import (
	"context"
	"database/sql"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMakeProduct(t *testing.T) {
	RandomMakeProduct(t)
}

func TestListProducts(t *testing.T) {

	for i := 0; i <= 10; i++ {
		RandomMakeProduct(t)
	}

	arg := ListProductsParams{
		Limit:  10,
		Offset: 0,
	}
	products, err := testQueries.ListProducts(context.Background(), arg)
	assert.NoError(t, err)
	for _, product := range products {
		assert.NotEmpty(t, product.IDProduct)
	}
}

func RandomMakeProduct(t *testing.T) Product {
	name := sql.NullString{
		String: "asd",
		Valid:  true,
	}

	arg := MakeNewProductParams{
		Name:  name,
		Stock: 1,
		Price: 22,
	}

	product, err := testQueries.MakeNewProduct(context.Background(), arg)

	assert.NoError(t, err)
	assert.NotEmpty(t, product)

	assert.Equal(t, arg.Name, product.Name)
	assert.Equal(t, arg.Stock, product.Stock)
	assert.Equal(t, arg.Price, product.Price)

	assert.NotZero(t, product.IDProduct)
	assert.NotZero(t, product.CreatedAt)

	return product
}
