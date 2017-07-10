package main

import "fmt"

// Calculator is
type Calculator interface {
	Calculate(float64) (float64, float64, float64)
}

// DefaultCalculator is
type DefaultCalculator struct{}

// NewCalculator is
func NewCalculator(calcType string) Calculator {
	switch calcType {
	case "cust":
		return &BigPubCalculator{}
	default:
		return &DefaultCalculator{}
	}
}

// Calculate is
func (c *DefaultCalculator) Calculate(sale float64) (float64, float64, float64) {
	netShare := c.NetworkShare() * sale
	adShare := c.AdvertiserShare() * sale
	pubShare := c.PublisherShare() * sale
	return netShare, adShare, pubShare
}

// NetworkShare is
func (c *DefaultCalculator) NetworkShare() float64 {
	fmt.Println("Default Calc NetworkShare: 0.15")
	return 0.15
}

// AdvertiserShare is
func (c *DefaultCalculator) AdvertiserShare() float64 {
	fmt.Println("Default Calc AdvertiserShare: 0.425")
	return 0.425
}

// PublisherShare is
func (c *DefaultCalculator) PublisherShare() float64 {
	fmt.Println("Default Calc PublisherShare: 0.425")
	return 0.425
}
