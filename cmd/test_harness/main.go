package main

import (
	"flag"
	"io/ioutil"
	"log"
	"net/http"
	"regexp"
)

func checkRegex(str string) {
	r, err := regexp.MatchString("[0-9]+,[0-9]+,[0-9]+,[0-9]+,[0-9]+,[0-9]+,[0-9]+,[0-9]+", str)
	if err != nil {
		log.Fatal(err)
	}

	if r == false {
		log.Fatal("Invalid pattern")
	}
}

func sendRequest(req http.Request, client http.Client) {
	resp, _ := client.Do(&req)
	body, _ := ioutil.ReadAll(resp.Body)
	log.Println("URL: " + req.URL.Path)
	log.Println("Reponse: " + string(body) + "\n")
}

func main() {
	var initialFloat = flag.String("initial-float", "10,10,10,10,10,10,10,10", "initial cash float for the vending machine")
	var productName = flag.String("product-name", "coke", "name of product in the vending machine")
	var productPrice = flag.String("product-price", "1.5", "price of product in the vending machine")
	var coinDeposit = flag.String("coin-deposit", "1,1,1,1,1,1,1,1", "coins to deposit")

	flag.Parse()

	checkRegex(*initialFloat)
	checkRegex(*coinDeposit)

	client := &http.Client{}

	req, _ := http.NewRequest(http.MethodPost, "http://localhost:9090/coins/"+*initialFloat, nil)
	log.Println("Setting initial float")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodPut, "http://localhost:9090/products/"+*productName+","+*productPrice, nil)
	log.Println("Adding product")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/coins", nil)
	log.Println("Getting coins in machine")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/products", nil)
	log.Println("Getting products in machine")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodPut, "http://localhost:9090/coins/"+*coinDeposit, nil)
	log.Println("Depositing coins")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/coins", nil)
	log.Println("Getting coins in machine")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/credit", nil)
	log.Println("Get credit")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/products/buy-product/"+*productName, nil)
	log.Println("Buy product and get change")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/coins", nil)
	log.Println("Get coins in machine")
	sendRequest(*req, *client)

	req, _ = http.NewRequest(http.MethodGet, "http://localhost:9090/credit", nil)
	log.Println("Get credit")
	sendRequest(*req, *client)
}
