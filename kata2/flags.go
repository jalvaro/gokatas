package main

import (
	"flag"
	"fmt"
	"math/rand"
	"time"
  "github.com/gin-gonic/gin"
)

func generateRandomSku() string {
	// <Product>
	//	<sku>TPGwnKVpvPRbsdemjwfHfgGdDofuUzAgWlTFiAhB</sku>
	//	<quantity>60</quantity>
	//</Product>
	template := "\n\t<Product>\n\t\t<sku>%s</sku>\n\t\t<quantity>%d</quantity>\n\t</Product>"

	return fmt.Sprintf(template, generateRandomString(40), rand.Intn(100))
}

func generateProducts() string {
  nProducts := rand.Intn(10)
  products := ""

  for i:= 0; i<nProducts; i++ {
    products = products + generateRandomSku()
  }

  return "<ProductList>" + products + "\n</ProductList>\n"
}

func generateRandomString(length int) string {
	const letterBytes = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ"
	str := make([]byte, length)

	for i := 0; i < length; i++ {
		str[i] = letterBytes[rand.Intn(len(letterBytes))]
	}
	return string(str)
}

func fakeLoad() {
	p := rand.Intn(100)
	sleep := 0

	if p >= 5 && p < 25 {
		sleep = rand.Intn(10)
	} else if p < 75 {
		sleep = 50 + rand.Intn(50)
	} else if p >= 75 {
		sleep = 200 + rand.Intn(500)
	}

	fmt.Println("Percentage:", p, ", Sleep:", sleep, "ms")
	time.Sleep(time.Duration(sleep) * time.Millisecond)
}

func ex1(length int, port int) {
  str := generateRandomString(length)
	sku := generateRandomSku()

	fakeLoad()

  fmt.Println("Port: ", port)
	fmt.Println("random string: ", str)
	fmt.Println("random sku: ", sku)
}

func ex2(port int) {
  // ab -c60 -n1000 http://127.0.0.1:8080/ping
  r := gin.Default()
  r.GET("/ping", func(c *gin.Context) {
      if (rand.Intn(100) < 10) {
        c.JSON(500, gin.H{
            "message": "Error!!!",
        })
      } else {
        fakeLoad()

        c.String(200, generateProducts())
      }
  })
  r.Run() // listen and server on 0.0.0.0:8080
}

func main() {
	// go run flags.go -h
	// go run flags.go -port=1222
	port := flag.Int("port", 8081, "the port of the service")
	length := flag.Int("length", 10, "the length of the random string")
	flag.Parse()

	rand.Seed(time.Now().UTC().UnixNano())

	ex1(*port, *length)
  ex2(*port)
}
