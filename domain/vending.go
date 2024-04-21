package domain

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"vending-mechine/data"
	"vending-mechine/pkg"
)

type VendingMachineDomain struct {
	Machines []*pkg.Machine
}

func NewVendingMachineDomain() *VendingMachineDomain {
	return &VendingMachineDomain{
		data.Mechins,
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
	dir, _ := os.Getwd()
	filePath := filepath.Join(dir, "/data/machines.json")
	// Find the machine in the JSON data
	machineIndex := -1
	for i, machine := range v.Machines {
		if machine.Name == mch.Name {
			machineIndex = i
			break
		}
	}
	if machineIndex == -1 {
		return "", errors.New("machine not found")
	}

	// Find the product in the machine's products
	for i, product := range v.Machines[machineIndex].Products {
		if product.Name == productName {
			if product.Count < 1 {
				return "", errors.New("product not available")
			} else if product.Price > coin {
				return "", errors.New("coin amount is not enough")
			}
			//product.Count-- // Update product count in memory
			leftCount := product.Count - 1

			// Update the Count value of the product in the JSON data
			v.Machines[machineIndex].Products[i].Count--

			// Marshal the updated data back to JSON
			updatedData, err := json.MarshalIndent(v.Machines, "", "    ")
			if err != nil {
				return "", fmt.Errorf("error marshalling updated data: %v", err)
			}

			// Write the updated data back to the JSON file
			if err := ioutil.WriteFile(filePath, updatedData, 0644); err != nil {
				return "", fmt.Errorf("error writing updated data to JSON file: %v", err)
			}
			message := fmt.Sprintf("Operation done - Left count: %d", leftCount)
			return message, nil
		}
	}

	return "", errors.New("product not found")
}
