package weather

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/lipgloss"
)

// Renderer handles beautiful output formatting
type Renderer struct {
	// Color palette
	primaryColor   lipgloss.Color
	secondaryColor lipgloss.Color
	accentColor    lipgloss.Color
	textColor      lipgloss.Color
	subtleColor    lipgloss.Color
	
	// Base styles
	titleStyle       lipgloss.Style
	headerStyle      lipgloss.Style
	cardStyle        lipgloss.Style
	labelStyle       lipgloss.Style
	valueStyle       lipgloss.Style
	iconStyle        lipgloss.Style
	subtleStyle      lipgloss.Style
	highlightStyle   lipgloss.Style
}

// NewRenderer creates a new renderer with beautiful styling
func NewRenderer() *Renderer {
	r := &Renderer{
		primaryColor:   lipgloss.Color("#3B82F6"),   // Blue
		secondaryColor: lipgloss.Color("#10B981"),   // Emerald
		accentColor:    lipgloss.Color("#F59E0B"),    // Amber
		textColor:      lipgloss.Color("#1F2937"),    // Gray-800
		subtleColor:    lipgloss.Color("#6B7280"),    // Gray-500
	}
	
	// Initialize styles
	r.initStyles()
	return r
}

func (r *Renderer) initStyles() {
	r.titleStyle = lipgloss.NewStyle().
		Foreground(r.primaryColor).
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(r.primaryColor)
	
	r.headerStyle = lipgloss.NewStyle().
		Foreground(r.primaryColor).
		Bold(true).
		Padding(0, 1)
	
	r.cardStyle = lipgloss.NewStyle().
		Border(lipgloss.RoundedBorder()).
		BorderForeground(r.subtleColor).
		Padding(1, 2).
		Margin(1, 0)
	
	r.labelStyle = lipgloss.NewStyle().
		Foreground(r.subtleColor).
		Bold(true)
	
	r.valueStyle = lipgloss.NewStyle().
		Foreground(r.textColor).
		Bold(true)
	
	r.iconStyle = lipgloss.NewStyle().
		Foreground(r.accentColor).
		Bold(true)
	
	r.subtleStyle = lipgloss.NewStyle().
		Foreground(r.subtleColor)
	
	r.highlightStyle = lipgloss.NewStyle().
		Foreground(r.secondaryColor).
		Bold(true)
}

// RenderCurrent renders current weather beautifully
func (r *Renderer) RenderCurrent(weather *CurrentWeather) string {
	var output strings.Builder
	
	// Title
	title := fmt.Sprintf("🌤️  Current Weather - %s, %s", weather.City, weather.Country)
	output.WriteString(r.titleStyle.Render(title))
	output.WriteString("\n\n")
	
	// Main weather info card
	mainCard := r.buildMainWeatherCard(weather)
	output.WriteString(mainCard)
	output.WriteString("\n")
		// Details cards in a row
	detailsRow := lipgloss.JoinHorizontal(
		lipgloss.Top,
		r.buildWindCard(weather),
		" ",
		r.buildAtmosphereCard(weather),
		" ",
		r.buildSunCard(weather),
	)
	output.WriteString(detailsRow)
	output.WriteString("\n")
	
	// Footer with timestamp
	footer := r.subtleStyle.Render(fmt.Sprintf("Last updated: %s", 
		weather.Timestamp.Format("Jan 2, 2006 at 3:04 PM")))
	output.WriteString(footer)
	output.WriteString("\n")
	
	return output.String()
}

func (r *Renderer) buildMainWeatherCard(weather *CurrentWeather) string {
	icon := WeatherIcon[weather.Icon]
	if icon == "" {
		icon = "🌤️"
	}
	
	temp := r.highlightStyle.Render(fmt.Sprintf("%.1f°C", weather.Temperature))
	feelsLike := r.subtleStyle.Render(fmt.Sprintf("Feels like %.1f°C", weather.FeelsLike))
	description := r.valueStyle.Render(strings.Title(weather.Description))
	
	content := fmt.Sprintf("%s  %s\n%s\n%s", 
		r.iconStyle.Render(icon), temp, description, feelsLike)
	
	return r.cardStyle.Copy().
		Width(40).
		Align(lipgloss.Center).
		Render(content)
}

func (r *Renderer) buildWindCard(weather *CurrentWeather) string {
	content := fmt.Sprintf("%s %s\n\n%s %.1f m/s\n%s %s",
		r.iconStyle.Render("💨"), r.headerStyle.Render("Wind"),
		r.labelStyle.Render("Speed:"), weather.WindSpeed,
		r.labelStyle.Render("Direction:"), r.getWindDirection(weather.WindDir))
	
	return r.cardStyle.Copy().Width(18).Height(5).Render(content)
}

func (r *Renderer) buildAtmosphereCard(weather *CurrentWeather) string {
	content := fmt.Sprintf("%s %s\n\n%s %d%%\n%s %d hPa\n%s %d%%",
		r.iconStyle.Render("🌡️"), r.headerStyle.Render("Atmosphere"),
		r.labelStyle.Render("Humidity:"), weather.Humidity,
		r.labelStyle.Render("Pressure:"), weather.Pressure,
		r.labelStyle.Render("Clouds:"), weather.CloudCover)
	
	return r.cardStyle.Copy().Width(20).Height(5).Render(content)
}

func (r *Renderer) buildSunCard(weather *CurrentWeather) string {
	content := fmt.Sprintf("%s %s\n\n%s %s\n%s %s",
		r.iconStyle.Render("🌅"), r.headerStyle.Render("Sun Times"),
		r.labelStyle.Render("Sunrise:"), weather.Sunrise.Format("3:04 AM"),
		r.labelStyle.Render("Sunset:"), weather.Sunset.Format("3:04 PM"))
	
	return r.cardStyle.Copy().Width(18).Height(5).Render(content)
}

// RenderForecast renders weather forecast beautifully
func (r *Renderer) RenderForecast(forecast *Forecast) string {
	var output strings.Builder
	
	// Title
	title := fmt.Sprintf("📅  5-Day Forecast - %s, %s", forecast.City, forecast.Country)
	output.WriteString(r.titleStyle.Render(title))
	output.WriteString("\n\n")
	
	// Forecast cards
	var cards []string
	for i, day := range forecast.Days {
		if i >= 5 { // Limit to 5 days
			break
		}
		cards = append(cards, r.buildForecastCard(day))
	}
	
	forecastRow := lipgloss.JoinHorizontal(lipgloss.Top, cards...)
	output.WriteString(forecastRow)
	output.WriteString("\n")
	
	return output.String()
}

func (r *Renderer) buildForecastCard(day ForecastDay) string {
	icon := WeatherIcon[day.Icon]
	if icon == "" {
		icon = "🌤️"
	}
	
	dayName := day.Date.Format("Mon")
	date := day.Date.Format("Jan 2")
	tempHigh := r.highlightStyle.Render(fmt.Sprintf("%.0f°", day.TempMax))
	tempLow := r.subtleStyle.Render(fmt.Sprintf("%.0f°", day.TempMin))
	description := strings.Title(day.Description)
	
	var rainInfo string
	if day.ChanceRain > 0 {
		rainInfo = fmt.Sprintf("🌧️ %d%%", day.ChanceRain)
	}
	
	content := fmt.Sprintf("%s\n%s\n%s  %s\n%s / %s\n%s\n%s",
		r.headerStyle.Render(dayName),
		r.subtleStyle.Render(date),
		r.iconStyle.Render(icon), description,
		tempHigh, tempLow,
		r.subtleStyle.Render(fmt.Sprintf("💨 %.1f m/s", day.WindSpeed)),
		r.subtleStyle.Render(rainInfo))
	
	return r.cardStyle.Copy().
		Width(18).
		Height(8).
		Render(content)
}

func (r *Renderer) getWindDirection(degrees int) string {
	directions := []string{"N", "NNE", "NE", "ENE", "E", "ESE", "SE", "SSE", 
		"S", "SSW", "SW", "WSW", "W", "WNW", "NW", "NNW"}
	index := int((float64(degrees) + 11.25) / 22.5)
	return directions[index%16]
}
