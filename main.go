package main

import (
	"fmt"
	"math/rand"
	"time"
)

var availableSymbols = []string{"BTC", "DOT", "SOL", "SHIB", "ETH", "THETA"}
var actualPrices = make(map[string]float64)
var pricesChannles = make(map[string](chan float64))

func main() {
	go updatePrice("BTC")
	go updatePrice("ETH")
	
	fmt.Println("The rest of my application can continue")

	select {}
}

func updatePrice(symbol string) {
	ticker := time.NewTicker(1 * time.Second)
	for _ = range ticker.C {
		if _, ok := actualPrices[symbol]; !ok {
			actualPrices[symbol] = 5.000
			fmt.Println(symbol, actualPrices[symbol]);
			continue
		}

		percentage := (rand.Float64() * (0.001 + 5.000)) + 0.001
		operation := rand.Intn(1)
		if operation == 0 {
			actualPrices[symbol] += (actualPrices[symbol] * (percentage / 100.00))
		} else {
			actualPrices[symbol] -= (actualPrices[symbol] * (percentage / 100.00))
		}
		fmt.Println(symbol, actualPrices[symbol]);
	}
}
