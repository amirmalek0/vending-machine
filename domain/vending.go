package domain

import (
	"errors"
	"fmt"
	"vending-mechine/data"
	"vending-mechine/pkg"
)

type VendingMachineDomain struct {
	Machines []*pkg.Machine
}

func NewVendingMachineDomain() *VendingMachineDomain {

	machines := data.Mechins

	return &VendingMachineDomain{
		machines,
	}
}

func (v *VendingMachineDomain) MachineList() []*pkg.Machine {
	return v.Machines
}

func (v *VendingMachineDomain) GetMachineByName(name string) (*pkg.Machine, error) {
	for _, machine := range v.Machines {
		if machine.Name == name {
			return machine, nil
		}
	}
	return nil, errors.New("machine not found")
}

func (v *VendingMachineDomain) BuyCoffee(machine string, coin int32) (string, error) {
	mch, err := v.GetMachineByName(machine)
	if err != nil {
		return "", errors.New("invalid machine name")
	}
	message, err := v.BuyOperation(mch, coin, data.CoffeeName)
	if err != nil {
		return "", err
	}
	return message, nil
}

func (v *VendingMachineDomain) BuyCoca(machine string, coin int32) (string, error) {
	mch, err := v.GetMachineByName(machine)
	if err != nil {
		return "", errors.New("invalid machine name")
	}
	message, err := v.BuyOperation(mch, coin, data.CocaName)
	if err != nil {
		return "", err
	}
	return message, nil
}
func (v *VendingMachineDomain) BuyCake(machine string, coin int32) (string, error) {
	mch, err := v.GetMachineByName(machine)
	if err != nil {
		return "", errors.New("invalid machine name")
	}
	message, err := v.BuyOperation(mch, coin, data.CakeName)
	if err != nil {
		return "", err
	}
	return message, nil
}

func (v *VendingMachineDomain) BuyOperation(mch *pkg.Machine, coin int32, productName string) (string, error) {
	for _, product := range mch.Products {
		if product.Name == productName {
			if product.Count < 1 {
				return "", errors.New("product not available")
			} else if product.Price > coin {
				return "", errors.New("coin amount is not enough")
			}
			product.Count--
			leftCount := product.Count
			message := fmt.Sprintf("Operation done - Left count: %d", leftCount)
			return message, nil
		}
	}
	return "", errors.New("invalid product name")
}
