package presentaion

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
	"vending-mechine/domain"
)

var handler *CliHandler

type CliHandler struct {
	domain   *domain.VendingMachineDomain
	commands []*cobra.Command
}

func NewCliHandler(domain *domain.VendingMachineDomain) *CliHandler {

	handler := &CliHandler{domain: domain}
	rootCmd.AddCommand(machineList())
	rootCmd.AddCommand(buyCoffee())
	rootCmd.AddCommand(buyCoca())
	rootCmd.AddCommand(buyCake())

	return handler
}

var rootCmd = &cobra.Command{
	Use:   "vending",
	Short: "vending is a trusted machine for coffee,coca and cake",
	Long:  "vending is a trusted machine for coffee,coca and cake",
	Run:   func(cmd *cobra.Command, args []string) {},
}

func init() {
	handler = NewCliHandler(domain.NewVendingMachineDomain())
}

func CliRun() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}

}
