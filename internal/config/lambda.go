package config

import "github.com/thiagozs/aws-tools/internal/models"

var PricingLambidaByRegion = map[string]models.PricingLambda{
	"us-east-1": {
		Region:          "US East (N. Virginia)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"us-east-2": {
		Region:          "US East (Ohio)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"us-west-1": {
		Region:          "US West (N. California)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"us-west-2": {
		Region:          "US West (Oregon)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"af-south-1": {
		Region:          "Africa (Cape Town)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001833,
	},
	"ap-east-1": {
		Region:          "Asia Pacific (Hong Kong)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00002333,
	},
	"ap-south-1": {
		Region:          "Asia Pacific (Mumbai)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"ap-northeast-3": {
		Region:          "Asia Pacific (Osaka-Local)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00002,
	},
	"ap-northeast-2": {
		Region:          "Asia Pacific (Seoul)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"ap-southeast-1": {
		Region:          "Asia Pacific (Singapore)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"ap-southeast-2": {
		Region:          "Asia Pacific (Sydney)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"ap-northeast-1": {
		Region:          "Asia Pacific (Tokyo)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"ca-central-1": {
		Region:          "Canada (Central)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"eu-central-1": {
		Region:          "EU (Frankfurt)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"eu-west-1": {
		Region:          "EU (Ireland)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"eu-west-2": {
		Region:          "EU (London)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"eu-south-1": {
		Region:          "EU (Milan)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00002,
	},
	"eu-west-3": {
		Region:          "EU (Paris)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"eu-north-1": {
		Region:          "EU (Stockholm)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
	"me-south-1": {
		Region:          "Middle East (Bahrain)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00002,
	},
	"sa-east-1": {
		Region:          "South America (São Paulo)",
		PerRequestPrice: 0.20,
		PerMemoryPrice:  0.00001667,
	},
}
