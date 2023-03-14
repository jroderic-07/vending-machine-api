# Vending Machine API
This project was created as my submission for the Oracle Analytics Coding challenge. It is an implementation of an API that can track changes and coins within a vending machine.
Three main requirements were implemented:
1. The API must be able to be initialised to a known state, for use when the machine is set up.
2. Coins deposited by the user are registered.
3. When an order is received, the correct amount of change is removed from the machine
Further functionality has also been implemented to add more usability when using and testing the API, further documentation can be found below.

## Installation
To install this project, clone this repository and use the Makefile:
```bash
git clone https://github.com/jroderic-07/vending-machine-api.git
make vending
```

## Instructions
For information about the command line arguments:
```bash
./bin/vending --help
```

If the binary is run without any arguments, by default port 8080 will be used, quantities for each type of coin will be set to 0 and no products will be added.
```bash
./bin/vending
```

Initial floats can be set using the REST API endpoint, for example:
```bash
curl -X POST localhost:<port>/coins/1,1,1,1,1,1,1,1
```

Products can be added using the REST API endpoint, for example:
```bash
curl -X PUT localhost<port>/products/coke,1.5
```

To run the command with arguments
```bash
./bin/vending -floats £2:::1,£1:::1,50p:::0,20p:::0,10p:::0,5p:::0,2p:::0,1p:::0 -products coke:::1.5 -port :9090
```

## Test harness
Rather than having to run curl commands to interact with the API, an interactive test harness has been included.
Once you have compiled the binary, you can run it providing an initial float, product name and price, and coins to deposit.
The test handler will:
1. Initialise the coin quantities using the intitial float
2. Create a product using the given name and price
3. Get all coin quantities
4. Get all product prices and names
5. Deposit coins
6. Get all coin quantities
7. Get user credit
8. Buy the product
9. Get all coin quantities
10. Get user credit

To compile the binary, run:
```bash
make test-harness
```

Ensure that the vending machine API is running on port 9090 in another terminal tab/window before using the test harness binary.

If you run the binary without any arguments, by default ten of each coin are passed as an initial float, the product added is coke for 1.50, and one of each coin is deposited.
```bash
./bin/test-harness
```

For more information about the command line arguments:
```
./bin/test-harness --help
```

## Tests
# Test harness
Default (expected input)
```bash
./bin/test-harness
```

Expected input
```bash
./bin/test-harness -initial-float 10,10,10,5,4,3,2,5 -product-name fanta -product-price 3 -coin-deposit 1,0,3,2,0,0,0,0
```

In-sufficient funds
Transaction will not go ahead
```bash
./bin/test-harness -initial-float 10,10,10,5,4,3,2,5 -product-name fanta -product-price 3 -coin-deposit 1,0,1,0,0,0,0,0
```

In-sufficient change
Transaction will not go ahead
```bash
./bin/test-harness -initial-float 10,0,0,0,0,0,0,0 -product-name fanta -product-price 3.5 -coin-deposit 2,0,0,0,0,0,0,0
```

In-sufficient change
Transaction will not go ahead
```bash
./bin/test-harness -initial-float 100,0,0,0,0,0,0,0 -product-name fanta -product-price 3.5 -coin-deposit 2,1,0,1,1,1,1,1
```

Erroneous (product price)
Default price will be used
```bash
./bin/test-harness -product-name fanta -product-price z
```

Erroneous (initial floats)
Test will not run
```bash
./bin/test-harness -initial-float z,z,z,z
```

Erroneous (coin deposit)
Test will not run
```bash
./bin/test-harness -coin-deposit z,z,z,z
```

## API Endpoints
Home page.
```
/
```

Get all product names and prices in vending machine.
```
GET /products
```

Get name and price of a particular product.
```
GET /products/{id}
```

Buy a particular product, using credit from the coins deposited so far. Will give back all change.
```
GET /products/buy-product/{id}
```

Add a product, using name and price accepted as parameters
```
PUT /products/{id}{price}
```

Get quantities of each coin type.
```
GET /coins
```

Get quantity of a particular coin type.
```
GET /coins/{id}
```

Initialise state of the vending machine by passing an initial cash float. Quantities of each type of coin accepted as parameters. Clears all previous coins.
```
POST /coins/{£2},{£1},{50p},{20p},{10p},{5p},{2p},{1p}
```

Deposit coins. Acceps quantities of each type of coin as parameters. Updates current credit.
```
PUT /coins/{£2},{£1},{50p},{20p},{10p},{5p},{2p},{1p}
```

Get credit, in other words how much money you have deposited, that can be used to buy products.
```
GET /credit
```

## Design Decisions
I decided to use Go for this task, my reasons being:
- It is the programming language that I am most comfortable using.
- It is has a relatively simple syntax, meaning that less time is spent worrying about syntax and more time is spent actually implementing features.
- There is a comprehensive standard library. This is useful in corporate setting where CSSAP approvals are an issue.
- It compiles and runs very quickly.
- It has a large amount of community support.

I decided to implement a REST API for this task, my reasons being:
- It is the most popular communication standard at the moment.
- They are simple to implement.
- They are scalable.
- It is seperates out components of the system. In this project, the API and vending machine components are completely decoupled, meaning that they can be worked on independently and easily debugged. 
I followed this un-official project structure standard: https://github.com/golang-standards/project-layout.

I included a makfile in this repo that can create two binaries: the vending machine and the test harnenss.
The makefile compiles these files: in cmd/vending/main.go or /cmd/test_harness/main.go.
The main.go file creates a vending machine object, and then an API object passing the vending machine object as a parameter.
The vending machine object is created using a builder design pattern.

## Dependencies
- A Unix-like system
- Go

## References
- https://github.com/golang-standards/project-layout
- https://refactoring.guru/design-patterns/go 
