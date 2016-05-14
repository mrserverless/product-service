package api

import (
        "github.com/yunspace/product-service/models"
        "io/ioutil"
        "log"
        "github.com/pquerna/ffjson/ffjson"
)

type ProductApi struct {
        Products []models.Product
}

func NewProductApi(seedfile string) *ProductApi {
        if buf, err := ioutil.ReadFile(seedfile); err != nil {
                log.Fatalf("failed to read seed file: %s", seedfile)
                return nil
        } else {
                products := []models.Product{}
                ffjson.Unmarshal(buf, &products)
                return &ProductApi{products}
        }
}

func (pApi ProductApi) GetProducts() []models.Product {
        return pApi.Products
}
//
//func (pApi ProductApi) GetProduct(id string) *models.Product {
//        return &pApi.products[0]
//}