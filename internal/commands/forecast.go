package commands

import (
	"fmt"
	"strings"

	"github.com/spf13/cobra"
	"weather-cli/internal/weather"
)

var ForecastCmd = &cobra.Command{
	Use:   "forecast [city]",
	Short: "Get weather forecast for a city",
	Long:  "Display the weather forecast for the specified city with beautiful formatting",
	Args:  cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		city := strings.Join(args, " ")
		
		// Create weather service
		service := weather.NewService()
		
		// Get forecast
		forecast, err := service.GetForecast(city)
		if err != nil {
			fmt.Printf("Error getting forecast data: %v\n", err)
			return
		}
		
		// Display with beautiful formatting
		renderer := weather.NewRenderer()
		output := renderer.RenderForecast(forecast)
		fmt.Print(output)
	},
}
