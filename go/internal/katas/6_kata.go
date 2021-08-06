package katas

import (
	"fmt"
	"math"
)

func RunKata6() {
	fmt.Println(Movie(500, 15, 0.9))  // 43
	fmt.Println(Movie(0, 10, 0.95))   // 2
	fmt.Println(Movie(2500, 20, 0.9)) // 135
}

// Going to the cinema: https://www.codewars.com/kata/562f91ff6a8b77dfe900006e/train/go
func Movie(card, ticket int, perc float64) (count int) {
	priceWithCard := float64(card)
	priceNormal := 0

	for {
		count++
		priceNormal += ticket
		priceWithCard += float64(ticket) * math.Pow(perc, float64(count))
		if math.Ceil(priceWithCard) < float64(priceNormal) {
			return count
		}
	}
}
