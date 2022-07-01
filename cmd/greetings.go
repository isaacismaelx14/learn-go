package cmd

import (
	"fmt"

	"github.com/isaacismaelx14/mycli/pkg/greetings"
	"github.com/spf13/cobra"
)

var greetingsCmd = &cobra.Command{
	Use:   "greetings",
	Short: "A brief description of your application",
	Long: `A longer description that spans multiple lines and likely contains`,
	Run: func(cmd *cobra.Command, args []string) {
		initGreetigns(args)
	},
}

func init() {
	rootCmd.AddCommand(greetingsCmd)
}

func initGreetigns(args []string) {
	fmt.Print(greetings.Get(args[0]))
}

