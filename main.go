package main

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
	"weather-cli/internal/commands"
)

var rootCmd = &cobra.Command{
	Use:   "weather-cli",
	Short: "A beautiful weather CLI application",
	Long: `Weather CLI is a beautifully designed command-line application 
that provides current weather information and forecasts with rich terminal output.`,
}

func main() {
	// Add subcommands
	rootCmd.AddCommand(commands.CurrentCmd)
	rootCmd.AddCommand(commands.ForecastCmd)
	rootCmd.AddCommand(commands.InteractiveCmd)

	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}
