package main

import "fmt"

func main() {
	c1 := NewCalculator("")
	c2 := NewCalculator("cust")
	sale := 10.50
	n, a, p := c1.Calculate(sale)
	fmt.Printf("Default Calc Results: Network %f Advertiser %f Publisher %f\n", n, a, p)
	n, a, p = c2.Calculate(sale)
	fmt.Printf("Default Calc Results: Network %f Advertiser %f Publisher %f\n", n, a, p)
}
