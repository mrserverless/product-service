package api

import (
        "testing"
)

func TestNewProductApi(t *testing.T) {
        // given
        seed := "seed.json"

        // when
        if pApi := NewProductApi(seed); len(pApi.Products) != 4 {
                t.Log(pApi)
                t.Errorf("ProductApi products length incorrect: %d", len(pApi.Products))
        }
}


func TestProductApi_GetProducts(t *testing.T) {

}