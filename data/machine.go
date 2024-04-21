package data

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"os"
	"path/filepath"
	"vending-mechine/pkg"
)

const (
	CocaName   = "coca"
	CoffeeName = "coffee"
	CakeName   = "cake"

	//MachineName1 = "machine_1"
	//MachineName2 = "machine_2"
)

var Mechins []*pkg.Machine

func LoadDataFromJSONFile(filePath string) error {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		return err
	}

	err = json.Unmarshal(data, &Mechins)
	if err != nil {
		return err
	}

	return nil
}

func init() {
	dir, err := os.Getwd()
	filePath := filepath.Join(dir, "/data/machines.json")

	err = LoadDataFromJSONFile(filePath)
	if err != nil {
		log.Fatalf("Error loading data from JSON file: %v", err)
	}
}
