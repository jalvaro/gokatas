package main

import (
	"encoding/json"
	"encoding/xml"
	"fmt"
	"math/rand"
	"time"
	"net/http"
	"io/ioutil"
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
	close   chan Parser
	pos     int
	res     []byte
}

func (p Parser) parse() {
	randomSleep()
	stock := Stock{}

	err := xml.Unmarshal(p.xmlData, &stock)
	if err != nil {
		return
	}

	randomSleep()

	res, err := json.Marshal(stock)
	if err != nil {
		return
	}

	randomSleep()
	p.res = res
	p.close <- p
}

func main() {

	res, err := http.Get("http://127.0.0.1:8080/ping")
	if err != nil {
		panic(err)
	}

	if (res.StatusCode != 200){
		panic(fmt.Sprintf("Error. StatusCode: %d", res.StatusCode))
	}

	xmlData, err := ioutil.ReadAll(res.Body)
	res.Body.Close()
	if err != nil {
		panic(err)
	}

	var parser Parser
	done := make(chan Parser, 10)

	for i := 0; i < 10; i++ {
		parser = Parser{xmlData, done, i, []byte{}}
		go parser.parse()
	}

	var p = <-done

	fmt.Println("Winner thread:", p.pos)
	fmt.Println(string(p.res))
}
