package main

import (
	"fmt"
	"os"
	"path/filepath"
	"time"
	"vending-mechine/data"
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
	dir, _ := os.Getwd()
	filePath := filepath.Join(dir, "/data/machines.json")
	machineDomain := domain.NewVendingMachineDomain()
	// Load data from the JSON file initially
	if err := data.LoadDataFromJSONFile(filePath); err != nil {
		fmt.Printf("Error loading data from JSON file: %v\n", err)
		return
	}
	// Start a goroutine to watch for file changes and reload data
	go watchFileChanges(filePath, machineDomain)
	if len(os.Args) <= 1 {
		presentaion.NewHttpServer("0.0.0.0", "8000", machineDomain).Start()
	}
	presentaion.CliRun()
}

func watchFileChanges(filePath string, machineDomain *domain.VendingMachineDomain) {
	var lastModTime time.Time

	for {
		fileInfo, err := os.Stat(filePath)
		if err != nil {
			fmt.Printf("Error getting file info: %v\n", err)
			time.Sleep(5 * time.Second) // Retry after 5 seconds
			continue
		}
		if lastModTime.IsZero() || fileInfo.ModTime().After(lastModTime) {
			// File has been modified, reload the data
			if err := data.LoadDataFromJSONFile(filePath); err != nil {
				fmt.Printf("Error reloading data from JSON file: %v\n", err)
			}
			lastModTime = fileInfo.ModTime()
		}
		time.Sleep(100 * time.Millisecond)
	}
}
