package main

import (
	"fmt"
	"math/rand"
	"time"
)

var availableSymbols = []string{"BTC", "DOT", "SOL", "SHIB", "ETH", "THETA"}
var actualPrices = make(map[string]float64)

func main() {
	pricesChannel := make(chan )

	ticker := time.NewTicker(1 * time.Second)
	for t := range ticker.C {
		updatePrice("BTC")
		fmt.Println(t, "Princes: ", actualPrices)
	}
}

func updatePrice(symbol string) bool {
	if _, ok := actualPrices[symbol]; !ok {
		actualPrices[symbol] = 5.000
		return true
	}

	percentage := (rand.Float64() * (0.001 + 5.000)) + 0.001
	operation := rand.Intn(1)
	if operation == 0 {
		fmt.Println("sube")
		fmt.Println((percentage / 100.00))
		actualPrices[symbol] += (actualPrices[symbol] * (percentage / 100.00))
	} else {
		fmt.Println("baja")
		actualPrices[symbol] -= (actualPrices[symbol] * (percentage / 100.00))
	}

	return true
}
