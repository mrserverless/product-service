# GoLang Product Service

[![Build Status](https://snap-ci.com/yunspace/product-service/branch/master/build_image)](https://snap-ci.com/yunspace/product-service/branch/master)
[![Docker Repository on Quay](https://quay.io/repository/yunspace/product-service/status "Docker Repository on Quay")](https://quay.io/repository/yunspace/product-service)

A simple Product Service written in GoLang using the best libraries:

- fasthttp
- fasthttprouter
- ffjson
- testify/assert

## Instalation

    go get github.com/yunspace/product-service

## Usage

Docker:

    docker run -d -p 8080:8080 quay.io/yunspace/product-service

Commandline: since there is no DB, you must run the Go executable from a directory where api/seed.json exists:

    product-service

This Service is deployed at `https://www.yunspace.com/products/`. Test using either

 1. PostMan: [![Run in Postman](https://run.pstmn.io/button.svg)](https://app.getpostman.com/run-collection/4ee7471904c06b844a99)
 2. [Command Line Client](https://github.com/yunspace/product-client)