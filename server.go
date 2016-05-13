package main

import (
        "log"

        "github.com/buaazp/fasthttprouter"
        "github.com/valyala/fasthttp"
        "github.com/pquerna/ffjson/ffjson"

        //"github.com/yunspace/product-service/api"
        "github.com/yunspace/product-service/models"
)

func Products(ctx *fasthttp.RequestCtx, ps fasthttprouter.Params) {
        var p models.Product

        JsonMarshall(ctx, p)
}

func JsonMarshall(ctx *fasthttp.RequestCtx, p models.Product) {
        ctx.SetContentType("application/json")
        if err := ffjson.NewEncoder(ctx).Encode(p); err != nil {
                log.Fatalf("error in ffjson.Encoder.Encode: %s", err)
        }
}

func main() {
        router := fasthttprouter.New()
        router.GET("/products/", Products)
        router.GET("/products/:id", Products)

        log.Fatal(fasthttp.ListenAndServe(":8080", router.Handler))
}