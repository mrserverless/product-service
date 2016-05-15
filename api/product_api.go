package api

import (
        "github.com/yunspace/product-service/models"
        "io/ioutil"
        "log"
        "github.com/pquerna/ffjson/ffjson"
        "errors"
)

type ProductApi struct {
        Products   []models.Product
        ProductMap map[string]models.Product
}

func NewProductApi(seedfile string) *ProductApi {
        if buf, err := ioutil.ReadFile(seedfile); err != nil {
                log.Fatalf("failed to read seed file: %s", seedfile)
                return nil
        } else {
                products := []models.Product{}
                ffjson.Unmarshal(buf, &products)
                return &ProductApi{products, mapProducts(products)}
        }
}

func mapProducts(products []models.Product) map[string]models.Product {
        productMap := map[string]models.Product{}
        for _, product := range products {
                productMap[product.Code] = product
        }
        return productMap
}

func (pApi *ProductApi) GetProducts() []models.Product {
        return pApi.Products
}

func (pApi *ProductApi) GetProductById(id string) (*models.Product, error) {
        if val, ok := pApi.ProductMap[id]; ok {
                return &val, nil
        } else {
                return nil, errors.New("not found")
        }
}