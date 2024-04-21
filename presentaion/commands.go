package presentaion

import (
	"fmt"
	"github.com/spf13/cobra"
	"strconv"
)

func machineList() *cobra.Command {
	return &cobra.Command{
		Use:   "machines",
		Short: "print available machines name",
		Args:  cobra.ExactArgs(0),
		Run: func(cmd *cobra.Command, args []string) {
			for _, machine := range handler.domain.Machines {
				fmt.Printf("Machine Name: %s\n", machine.Name)
				for _, product := range machine.Products {
					fmt.Printf("Product: %s, Count: %d\n", product.Name, product.Count)
				}
			}
		},
	}
}
func buyCoffee() *cobra.Command {
	return &cobra.Command{
		Use:   "coffee",
		Short: "buy coffee",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			machine := args[0]
			coin, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				fmt.Println("coin should be number")
				return
			}
			message, err := handler.domain.BuyCoffee(machine, int32(coin))
			if err != nil {
				fmt.Println(err.Error())
				cmd.Help()
				return
			}
			fmt.Println(message)
		},
	}
}

func buyCoca() *cobra.Command {
	return &cobra.Command{
		Use:   "coca",
		Short: "buy coca",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			machine := args[0]
			coin, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {
				fmt.Println("coin should be number")
				cmd.Help()
				return
			}
			message, err := handler.domain.BuyCoca(machine, int32(coin))
			if err != nil {
				fmt.Println(err.Error())
				cmd.Help()
				return
			}
			fmt.Println(message)
		},
	}
}
func buyCake() *cobra.Command {
	return &cobra.Command{
		Use:   "cake",
		Short: "buy cake",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			machine := args[0]
			coin, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {

				fmt.Println("coin should be number")
				cmd.Help()
				return
			}
			message, err := handler.domain.BuyCake(machine, int32(coin))
			if err != nil {
				fmt.Println(err.Error())
				cmd.Help()
				return
			}
			fmt.Println(message)
		},
	}
}
