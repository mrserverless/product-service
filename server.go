package main

import (
        "log"

        "github.com/buaazp/fasthttprouter"
        "github.com/valyala/fasthttp"
        "github.com/pquerna/ffjson/ffjson"

        "github.com/yunspace/product-service/api"
)

func main() {
        router := fasthttprouter.New()
        productApi := api.NewProductApi("api/seed.json")

        // CRUD
        router.GET("/products/", Products(productApi))
        router.GET("/products/:id",  Products(productApi))

        log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

// Products handler
func Products(productApi *api.ProductApi) fasthttprouter.Handle {
        return func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
                if ps.ByName("id") != "" {
                        JsonMarshall(ctx, productApi.GetProductById(ps.ByName("id")))
                } else {
                        JsonMarshall(ctx, productApi.GetProducts())
                }
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

