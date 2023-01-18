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
				fmt.Println(machine.Name)
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
		Short: "print available machines name",
		Args:  cobra.ExactArgs(2),
		Run: func(cmd *cobra.Command, args []string) {
			machine := args[0]
			coin, err := strconv.ParseInt(args[1], 10, 32)
			if err != nil {

				fmt.Println("coin should be number")
				cmd.Help()
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
