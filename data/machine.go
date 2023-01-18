package data

import "vending-mechine/pkg"

const (
	CocaName   = "coca"
	CoffeeName = "coffee"

	MachineName1 = "machine_1"
	MachineName2 = "machine_2"
)

var Mechins = []*pkg.Machine{
	{
		Name: MachineName1,
		Products: []*pkg.Product{
			{
				Name:  CoffeeName,
				Price: 20,
				Count: 1000,
			},
			{
				Name:  CocaName,
				Price: 10,
				Count: 1000,
			},
		},
	},
	{
		Name: MachineName2,
		Products: []*pkg.Product{
			{
				Name:  CoffeeName,
				Price: 20,
				Count: 1000,
			},
			{
				Name:  CocaName,
				Price: 10,
				Count: 1000,
			},
		},
	},
}
