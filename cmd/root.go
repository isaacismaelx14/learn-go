package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use: "mycli",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains
	examples and usage of using your application.`,
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root called")
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}