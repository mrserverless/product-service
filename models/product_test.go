package models

import (
        "testing"
        "github.com/pquerna/ffjson/ffjson"
        "io/ioutil"
)

const (
        productArray = `[
    {
      "code": "ult_small",
      "name": "Unlimited 1GB",
      "description": "",
      "price": 24.90,
      "expiry": 30,
      "is_plan": true,
      "is_ulimited": true,
      "size_mb": 1024,
      "4g": true,
      "auto_renew": true,
      "terms_url": "http://amaysim.com.au/terms-conditions/special-conditi"
    }
  ]`
)

func TestMarshallProductSlice(t *testing.T) {

        p := []Product{
                Product{"ult_small", "Unlimited 1GB", "", 24.90, 30, true, true, 1024, true, true, ""},
        }
        expected := `[{"Code":"ult_small","Name":"Unlimited 1GB","Description":"","Price":24.9,"Expiry":30,"IsPlan":true,"IsUnlimited":true,"SizeMb":1024,"Is4g":true,"AutoRenew":true,"TermsUrl":""}]`

        if buf, _ := ffjson.Marshal(&p); string(buf) != expected {
                t.Fatalf("failed to marshal Product array to JSON %s", string(buf))
        }

}

func TestUnmarshallProductSlice(t *testing.T) {
        p := []Product{}
        if ffjson.Unmarshal([]byte(productArray), &p); (len(p) != 1) {
                t.Fatalf("incorrect number of products %d", len(p))
        }
}

func TestUnmarshallProductFile(t *testing.T) {
        products := []Product{}
        buf, _ := ioutil.ReadFile("../api/seed.json")
        if ffjson.Unmarshal(buf, &products); len(products) != 5 {
                t.Fatalf("incorrect number of products: %d", len(products))
        }
}