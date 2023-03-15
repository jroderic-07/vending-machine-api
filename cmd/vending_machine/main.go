package main

import (
	"flag"
	"log"
	"runtime"
	"strings"

	"github.com/jroderic-07/vending-machine-api/internal/api"
	"github.com/jroderic-07/vending-machine-api/internal/vending_machine"
	"github.com/jroderic-07/vending-machine-api/pkg/error_catching"
)

type arrayFlags []string

func (l *arrayFlags) String() string {
	return "array of flags"
}

func (l *arrayFlags) Set(value string) error {
	for _, str := range strings.Split(value, ",") {
		*l = append(*l, str)
	}

	return nil
}

var floatList arrayFlags
var productList arrayFlags

func main() {
	var cpuNum = flag.Int("cpus", 0, "Number of CPUs to use")
	var port = flag.String("port", ":8080", "Port to run the REST API on")
	flag.Var(&floatList, "floats", "Quantites of each type of coin given to vending machine as float, order: £2, £1, 50p, 20p, 10p, 5p, 2p, 1p, format COIN:::quantity - seperated by comas")
	flag.Var(&productList, "products", "List of products and prices that the vending machine will sell, format: PRODUCT:::PRICE - seperated by comas")

	flag.Parse()

	portAvailable, err := error_catching.CheckPort(*port)
	if portAvailable == false {
		log.Fatalln(err)
	}

	if *cpuNum != 0 {
		runtime.GOMAXPROCS(*cpuNum)
	}

	vendingMachineBuilder := vending_machine.GetBuilder()
	director := vending_machine.NewDirector(vendingMachineBuilder)
	vendingMachine := director.BuildVendingMachine(floatList, productList)

	restApi := api.New(&vendingMachine, *port)
	restApi.ServeAPI()
}
