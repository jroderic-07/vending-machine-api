package api

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"strconv"

	"github.com/jroderic-07/vending-machine-api/internal/vending_machine"

	"github.com/gorilla/mux"
)

type api struct {
	router         mux.Router
	port           string
	vendingMachine *vending_machine.VendingMachine
}

func (a *api) log(r *http.Request) {
	log.Println("URl: " + r.URL.Path + " Method: " + r.Method)
}

func (a *api) homeLink(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	fmt.Fprintf(w, "Welcome to my vending machine REST api!")
}

func (a *api) initCoints(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	coins := make(map[string]int)
	var err error
	var quantity int
	for _, coin := range vending_machine.CoinTypes {
		quantity, err = strconv.Atoi(mux.Vars(r)[coin])
		if err == nil {
			coins[coin] = quantity
		} else {
			log.Println(err)
		}
	}
	a.vendingMachine.SetFloatsAPI(coins)
}

func (a *api) getProducts(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	json.NewEncoder(w).Encode(a.vendingMachine.GetProducts())
}

func (a *api) getCoints(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	json.NewEncoder(w).Encode(a.vendingMachine.GetCoins())
}

func (a *api) getOneProduct(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	productID := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(a.vendingMachine.GetProducts()[productID])
}

func (a *api) getOneCoin(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	floatID := mux.Vars(r)["id"]
	json.NewEncoder(w).Encode(a.vendingMachine.GetCoins()[floatID])
}

func (a *api) depositCoins(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	coins := make(map[string]int)
	var err error
	var quantity int
	for _, coin := range vending_machine.CoinTypes {
		quantity, err = strconv.Atoi(mux.Vars(r)[coin])
		if err == nil {
			coins[coin] = quantity
		} else {
			log.Println(err)
		}
	}
	a.vendingMachine.DepositCoins(coins)
}

func (a *api) getCredit(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	json.NewEncoder(w).Encode(float64(a.vendingMachine.GetCredit()) / 100)

}

func (a *api) buyProduct(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	price := int(a.vendingMachine.GetProducts()[mux.Vars(r)["id"]] * 100)
	output, change := a.vendingMachine.BuyProduct(price)

	if output != "" {
		json.NewEncoder(w).Encode(output)
	} else {
		json.NewEncoder(w).Encode(change)
	}
}

func (a *api) addProduct(w http.ResponseWriter, r *http.Request) {
	a.log(r)

	price, err := strconv.ParseFloat(mux.Vars(r)["price"], 64)

	if err == nil {
		a.vendingMachine.AddProduct(mux.Vars(r)["id"], price)
	} else {
		log.Println(err)
	}
}

func (a *api) createAPI() *api {
	a.router = *mux.NewRouter()
	a.router.HandleFunc("/", a.homeLink)

	a.router.HandleFunc("/products", a.getProducts).Methods("GET")
	a.router.HandleFunc("/products/{id}", a.getOneProduct).Methods("GET")
	a.router.HandleFunc("/products/buy-product/{id}", a.buyProduct).Methods("GET")
	a.router.HandleFunc("/products/{id},{price}", a.addProduct).Methods("PUT")

	a.router.HandleFunc("/coins", a.getCoints).Methods("GET")
	a.router.HandleFunc("/coins/{id}", a.getOneCoin).Methods("GET")
	a.router.HandleFunc("/coins/{£2},{£1},{50p},{20p},{10p},{5p},{2p},{1p}", a.initCoints).Methods("POST")
	a.router.HandleFunc("/coins/{£2},{£1},{50p},{20p},{10p},{5p},{2p},{1p}", a.depositCoins).Methods("PUT")

	a.router.HandleFunc("/credit", a.getCredit).Methods("GET")

	return a
}

func (a *api) setPort(port string) *api {
	a.port = port
	return a
}

func (a *api) setVendingMachine(vendingMachine *vending_machine.VendingMachine) *api {
	a.vendingMachine = vendingMachine
	return a
}

func (a *api) ServeAPI() {
	log.Println("Starting REST api on port " + a.port)
	log.Fatal(http.ListenAndServe(a.port, &a.router))
}

func New(vendingMachineArgument *vending_machine.VendingMachine, portArgument string) *api {
	api := &api{}
	return api.createAPI().setPort(portArgument).setVendingMachine(vendingMachineArgument)
}
