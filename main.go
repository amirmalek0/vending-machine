package main

import (
	"os"
	"vending-mechine/domain"
	"vending-mechine/presentaion"
)

/*import (
	domain2 "vending-mechine/domain"
	"vending-mechine/presentaion"
)

func main() {
	domain := domain2.NewVendingMachineDomain()
}
*/

func main() {
	machineDomain := domain.NewVendingMachineDomain()
	if len(os.Args) <= 1 {
		presentaion.NewHttpServer("0.0.0.0", "8000", machineDomain).Start()
	}
	presentaion.CliRun()
}
