package main

import (
	"encoding/json"
	"fmt"
	"log"
)

type Product struct {
	ProductID      int    `json:"productId"`
	Manufacturer   string `json:"manufacturer"`
	Sku            string `json:"sku"`
	Upc            string `json:"upc"`
	PricePerUnit   string `json:"pricePerUnit"`
	QuantityOnHand int    `json:"quantityOnHand"`
	ProductName    string `json:"productName"`
}

func main() {
	//marshal data into JSON
	/*productToJSON := &Product{
		ProductID:      123,
		Manufacturer:   "Big Box Company",
		PricePerUnit:   "12.99",
		Sku:            "4561qHJK",
		Upc:            "414654444566",
		QuantityOnHand: 28,
		ProductName:    "Gizmo",
	}
	*/
	//productToJSON, err := json.Marshal(product)
	/* if err != nil {
		log.Fatal(err)
	}

	*/

	//unmarshal JSON data

	productToDecode := `{
		"productId":456,
		"manufacturer":"Small Box Company",
		"sku":"4hslj90JKL",
		"upc":"424654445566",
		"pricePerUnit":"$9.99",
		"quantityOnHand":18,
		"productName":"Sprocket"
		}`

	product := Product{}

	err := json.Unmarshal([]byte(productToDecode), &product)
	if err != nil {
		log.Fatal(err)
	}

	fmt.Printf("json unmarshal product: %s \n", product.ProductName)

	//fmt.Println(string(productToJSON))
}
