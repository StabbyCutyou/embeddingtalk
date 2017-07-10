package main

import "fmt"

type BigPubCalculator struct {
	DefaultCalculator
}

func (c *BigPubCalculator) NetworkShare() float64 {
	fmt.Println("BigPub Calc NetworkShare: 0.10")
	return 0.10
}

func (c *BigPubCalculator) PublisherShare() float64 {
	fmt.Println("BigPub Calc PublisherShare: 0.475")
	return 0.475
}
