package main

import (
        "log"

        "github.com/buaazp/fasthttprouter"
        "github.com/valyala/fasthttp"
        "github.com/pquerna/ffjson/ffjson"

        "github.com/yunspace/product-service/api"
        "github.com/yunspace/product-service/models"
)

func main() {
        router := fasthttprouter.New()
        pApi := api.NewProductApi("api/seed.json")

        // CRUD
        router.GET("/products/", Products(pApi))
        router.GET("/products/:id",  Products(pApi))

        log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

// Products handler
func Products(pApi *api.ProductApi) fasthttprouter.Handle {
        return func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
                //p := pApi.GetProducts()
                JsonMarshall(ctx, &[]models.Product{})
        }

}

func JsonMarshall(ctx *fasthttp.RequestCtx, v interface{}) {
        if bytes, err := ffjson.Marshal(v); err != nil {
                log.Fatalf("failed to marshal Product json: %s", err)
                ctx.Error("failed to marshal Product json", 500)
        } else {
                ctx.Success("application/json", bytes)
                ffjson.Pool(bytes)
        }
}

