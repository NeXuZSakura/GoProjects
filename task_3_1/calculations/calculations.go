package calculations

import (
	"fmt"
	"log"
)

func Calculate(n int64, flag bool) int64 {
	if flag {
		fmt.Println("Start calculations...")
		fmt.Printf("Calculate < %d >\n", n)
	}
	var result int64
	if n >= 0 {
		if n != 0 {
			for i := int64(1); i <= n; i++ {
				result *= i
			}
		} else {
			result = 1
		}
	} else {
		log.Fatal("Error: number need to be positive!")
	}
	if flag {
		fmt.Println("Calculations complete!")
	}
	return result
}
