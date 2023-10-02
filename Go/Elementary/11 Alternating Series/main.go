package main

import (
	"fmt"
	"math"
)

func main() {

	const ONE_MILLION int = 1000000
	var sum float64 = 0.0

	for k := 2; k <= ONE_MILLION; k++ {
		sum += math.Pow(-1.0, float64(k)+1.0) / (2.0 * (float64(k) - 1.0))
	}

	fmt.Printf("The sum of ((-1)^{k+1})/(2 * k-1) where 1 >= k <= 1,000,000 is %f\n", sum)
	fmt.Printf("This, multiplied by 4, is %f", sum*4)

}
