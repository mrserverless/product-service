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
        router.GET("/products/:id", Products(productApi))

        log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}

func Products(productApi *api.ProductApi) fasthttprouter.Handle {
        return func(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
                if ps.ByName("id") != "" {
                        if product, err := productApi.GetProductById(ps.ByName("id")); err != nil {
                                HandleError(ctx, err)
                        } else {
                                JsonMarshall(ctx, product)
                        }
                } else {
                        JsonMarshall(ctx, productApi.GetProducts())
                }
        }
}

func JsonMarshall(ctx *fasthttp.RequestCtx, v interface{}) {
        if buf, err := ffjson.Marshal(v); err != nil {
                HandleError(ctx, err)
        } else {
                ctx.Success("application/json", buf)
                ffjson.Pool(buf)
        }
}

func HandleError(ctx *fasthttp.RequestCtx, err error) {
        if (err.Error() == "not found") {
                ctx.Error("product not found", 404)
        } else {
                log.Fatalf(err.Error())
                ctx.Error("unexpected internal error", 500)
        }
}