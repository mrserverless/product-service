package api

import (
        "testing"
        "github.com/stretchr/testify/assert"
)

func TestNewProductApi(t *testing.T) {
        // given
        seed := "seed.json"

        // when
        pApi := NewProductApi(seed)

        // then
        assert.Equal(t, len(pApi.Products), 5, "should have 4 products")
        assert.Equal(t, len(pApi.ProductMap), 5, "should have 4 in product map")

}

func TestProductApi_GetProducts(t *testing.T) {
        // given
        pApi := NewProductApi("seed.json")

        // when
        products := pApi.GetProducts();

        // then
        assert.Equal(t, len(products), 5, "should have 5 products")

}

func TestProductApi_GetProduct(t *testing.T) {
        // given
        pApi := NewProductApi("seed.json")
        codes := []string{"ult_small", "ult_medium", "ult_large", "ult_xlarge" }

        // then
        for _, code := range codes {
                p, _ := pApi.GetProductById(code)
                assert.Equal(t, code, p.Code)
        }
}
