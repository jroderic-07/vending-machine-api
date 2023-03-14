package vending_machine

import (
	"strconv"
	"strings"

	"vending/pkg/error_catching"
)

var CoinTypes = []string{"£2", "£1", "50p", "20p", "10p", "5p", "2p", "1p"}

// Vending machine builder interface
type IVendingMachineBuilder interface {
	setfloats(floatsFlag []string)
	setproducts(productsFlag []string)
	getVendingMachine() VendingMachine
}

func GetBuilder() IVendingMachineBuilder {
	return newVendingMachineBuilder()
}

// Vending machine builder
type VendingMachineBuilder struct {
	coins    map[string]int
	products map[string]float64
}

func newVendingMachineBuilder() *VendingMachineBuilder {
	return &VendingMachineBuilder{}
}

func (b *VendingMachineBuilder) setfloats(floatsFlag []string) {
	floats := make(map[string]int)
	for _, coin := range CoinTypes {
		floats[coin] = 0
	}

	var floatAndQuantity []string
	var quantity int
	var err error

	for _, float := range floatsFlag {
		if error_catching.CheckPattern(float) {
			floatAndQuantity = strings.Split(float, ":::")
			quantity, err = strconv.Atoi(floatAndQuantity[1])
			if err == nil {
				floats[floatAndQuantity[0]] = quantity
			}
		}
	}

	b.coins = floats
}

func (b *VendingMachineBuilder) setproducts(productsFlag []string) {
	products := make(map[string]float64)
	var productAndPrice []string
	var price float64
	var err error

	for _, product := range productsFlag {
		if error_catching.CheckPattern(product) {
			productAndPrice = strings.Split(product, ":::")
			price, err = strconv.ParseFloat(productAndPrice[1], 64)
			if err == nil {
				products[productAndPrice[0]] = price
			}
		}
	}

	b.products = products
}

func (b *VendingMachineBuilder) getVendingMachine() VendingMachine {
	return VendingMachine{
		coins:    b.coins,
		products: b.products,
	}
}

// Vending machine product
type VendingMachine struct {
	coins    map[string]int
	products map[string]float64
	credit   int
}

func (b *VendingMachine) GetCoins() map[string]int {
	return b.coins
}

func (b *VendingMachine) GetProducts() map[string]float64 {
	return b.products
}

func (b *VendingMachine) SetFloatsAPI(floats map[string]int) {
	b.coins = floats
	b.credit = 0
}

func (b *VendingMachine) AddProduct(productName string, productPrice float64) {
	b.products[productName] = productPrice
}

func (b *VendingMachine) DepositCoins(floats map[string]int) {
	for coinType, coinQuantity := range floats {
		b.coins[coinType] = b.coins[coinType] + coinQuantity
	}

	var value string
	var valueInt int
	var err error
	for _, coinType := range CoinTypes {
		if coinType == "£2" {
			b.credit += 200
		} else if coinType == "£1" {
			b.credit += 100
		} else {
			value = strings.TrimSuffix(coinType, "p")
			valueInt, err = strconv.Atoi(value)
			if err == nil {
				b.credit += (valueInt * floats[coinType])
			}
		}
	}
}

func (b *VendingMachine) GetCredit() int {
	return b.credit
}

func (b *VendingMachine) BuyProduct(productPrice int) (string, map[string]int) {
	if productPrice > b.credit {
		return "Insufficient funds", nil
	}

	changeRequired := b.credit - productPrice
	change := make(map[string]int)
	coinsCopy := b.coins
	var value int
	var valueStr string
	var err error

	for _, coinType := range CoinTypes {
		if coinType == "£2" {
			value = 200
		} else if coinType == "£1" {
			value = 100
		} else {
			valueStr = strings.TrimSuffix(coinType, "p")
			value, err = strconv.Atoi(valueStr)
			if err != nil {
				continue
			}
		}

		for changeRequired >= value {
			if coinsCopy[coinType] != 0 {
				coinsCopy[coinType] -= 1
				change[coinType] = change[coinType] + 1
				changeRequired = changeRequired - value
			}
		}
	}

	if changeRequired > 0 {
		return "Insufficient change", nil
	}

	b.coins = coinsCopy
	b.credit = 0
	return "", change
}

// Director
type Director struct {
	builder IVendingMachineBuilder
}

func NewDirector(b IVendingMachineBuilder) *Director {
	return &Director{
		builder: b,
	}
}

func (d *Director) BuildVendingMachine(floats []string, products []string) VendingMachine {
	d.builder.setfloats(floats)
	d.builder.setproducts(products)

	return d.builder.getVendingMachine()
}
