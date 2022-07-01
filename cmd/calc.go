package cmd

import (
	"fmt"
	"os"
	"strconv"

	"github.com/isaacismaelx14/mycli/pkg/calc"
	"github.com/spf13/cobra"
)
var add bool
var substract bool

var calcCmd = &cobra.Command{
	Use:   "calc",
	Short: "A simple calculator",
	Long: `Addition, subtraction, multiplication and division of two numbers.`,
	Run: func(cmd *cobra.Command, args []string) {
		initCalc(args)
	},
}

func init() {
	cobra.OnInitialize()
	rootCmd.AddCommand(calcCmd)
	calcCmd.PersistentFlags().BoolVarP(&add, "add", "a", false, "verbose output")
	calcCmd.PersistentFlags().BoolVarP(&substract, "subtract", "s", false, "Verbose output")
}

func initCalc(args []string){
	x, xErr := strconv.Atoi(args[0])
	y, yErr := strconv.Atoi(args[1])

	if (xErr  != nil || yErr != nil) {
		fmt.Println("Error: Invalid input")
		os.Exit(1)
	}
	
	if add {
		fmt.Println(calc.Add(x, y))
	} else if substract {
		fmt.Println(calc.Substract(x, y))
	} else {
		fmt.Println("No operation selected")
	}
}
