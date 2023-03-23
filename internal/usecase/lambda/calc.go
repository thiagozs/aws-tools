package usecase

import (
	"fmt"

	"github.com/thiagozs/aws-tools/internal/config"
)

type UseCaseCalcLambda struct{}

func NewUseCaseCalcLambda() *UseCaseCalcLambda {
	return &UseCaseCalcLambda{}
}

func (cc UseCaseCalcLambda) CalculateCost(region string, memory int,
	requests int, executionTime int) (float64, error) {

	pricing, ok := config.PricingLambidaByRegion[region]
	if !ok {
		return 0.0, fmt.Errorf("unknown region %s", region)
	}

	gbSec := float64(requests) * float64(memory) / 1024.0

	if executionTime > 1 {
		gbSec = gbSec * (float64(executionTime) / 1000.0) // convert to seconds
	}

	computeCharges := gbSec * pricing.PerMemoryPrice
	requestCharges := float64(requests) / 1000000.0 * pricing.PerRequestPrice
	totalCharges := computeCharges + requestCharges

	return totalCharges, nil
}
