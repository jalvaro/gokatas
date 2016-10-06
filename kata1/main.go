package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"
)

type Product struct {
	Sku      string `xml:"sku" json:"sku"`
	Quantity int    `xml:"quantity" json:"quantity"`
}

type Stock struct {
	ProductList []Product `xml:"Product" json:"products"`
}

func randomSleep() {
    time.Sleep(time.Duration(rand.Int31n(1000)) * time.Millisecond)
}

type Parser struct {
    xmlData []byte
    close chan Parser
    pos int
    res []byte
}

func (p Parser) parse() {

    fmt.Println(p.pos)
	
    randomSleep()

	stock := Stock{}

	err := xml.Unmarshal(p.xmlData, &stock)

	if err != nil {
        return
	}

	randomSleep()

	fmt.Println(stock.ProductList[0].Sku)
	fmt.Println(stock)

	res, err := json.Marshal(stock)

    if err != nil {
        return
    }

	randomSleep()

    p.res = res

    p.close <- p
}

func main() {

	xmlData := []byte(`<?xml version="1.0" encoding="UTF-8" ?>
<ProductList>
    <Product>
        <sku>ABC123</sku>
        <quantity>2</quantity>
    </Product>
    <Product>
        <sku>ABC124</sku>
        <quantity>20</quantity>
    </Product>
</ProductList>`)

    var parser Parser
    done := make(chan Parser, 10)

    for i:=0; i<10; i++ {
        parser = Parser{xmlData, done, i, []byte{}}
        go parser.parse()
    }

	var p = <- done

    fmt.Println("Winner: ", p.pos)
    fmt.Println(string(p.res))
}
