package models

type PricingLambda struct {
	Region          string
	PerRequestPrice float64 // in USD per 1 million requests
	PerMemoryPrice  float64 // in USD per GB-second
}
