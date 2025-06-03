package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"weather-cli/internal/weather"
)

var CurrentCmd = &cobra.Command{
	Use:   "current [city]",
	Short: "Get current weather for a city",
	Long:  "Display the current weather conditions for the specified city with beautiful formatting",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := strings.Join(args, " ")
		
		// Create weather service
		service := weather.NewService()
		
		// Get current weather
		currentWeather, err := service.GetCurrent(city)
		if err != nil {
			fmt.Printf("Error getting weather data: %v\n", err)
			return
		}
		
		// Display with beautiful formatting
		renderer := weather.NewRenderer()
		output := renderer.RenderCurrent(currentWeather)
		fmt.Print(output)
	},
}
