# Weather CLI - Features & Usage Guide

## 🎯 Overview

Weather CLI is a beautiful, feature-rich command-line weather application written in Go. It provides an elegant terminal interface similar to Spectre.Console for .NET, offering both command-line and interactive modes.

## ✨ Key Features

### 🎨 Beautiful UI Components
- **Rich Color Scheme**: Modern blue/emerald/amber color palette
- **Weather Icons**: Unicode emoji icons for different weather conditions
- **Styled Cards**: Rounded borders with elegant spacing
- **Responsive Layout**: Adapts to terminal width
- **Loading States**: Smooth loading animations

### 🌤️ Weather Data
- **Current Weather**: Temperature, feels-like, humidity, pressure, wind
- **5-Day Forecast**: Daily temperature ranges and conditions
- **Global Coverage**: Any city worldwide
- **Weather Icons**: Visual weather condition indicators
- **Detailed Metrics**: Wind speed/direction, cloud cover, sunrise/sunset

### 🖥️ Interface Modes
- **Command Line**: Quick weather checks
- **Interactive TUI**: Full-screen browsing experience
- **Responsive Design**: Works on any terminal size

## 🚀 Quick Start

### 1. Build the Application
```bash
go build -o weather-cli
```

### 2. Basic Usage
```bash
# Current weather
./weather-cli current "London"

# 5-day forecast  
./weather-cli forecast "Tokyo"

# Interactive mode
./weather-cli interactive
```

### 3. With Real Weather Data
```bash
# Set API key
export WEATHER_API_KEY="your-openweathermap-api-key"

# Or on Windows
set WEATHER_API_KEY=your-openweathermap-api-key
```

## 📋 Command Reference

### Current Weather
```bash
weather-cli current [city]
```
**Examples:**
- `weather-cli current London`
- `weather-cli current "New York"`
- `weather-cli current "San Francisco"`

**Output includes:**
- Current temperature and "feels like"
- Weather description with icon
- Wind speed and direction
- Humidity, pressure, cloud cover
- Sunrise and sunset times

### Forecast
```bash
weather-cli forecast [city]
```
**Examples:**
- `weather-cli forecast Paris`
- `weather-cli forecast "Los Angeles"`

**Output includes:**
- 5-day weather outlook
- Daily high/low temperatures
- Weather conditions with icons
- Wind speed and precipitation chance

### Interactive Mode
```bash
weather-cli interactive
```
**Features:**
- Search for cities interactively
- Switch between current and forecast views
- Beautiful full-screen interface
- Keyboard navigation

**Controls:**
- `Enter`: Search for weather
- `F`: Switch to forecast view
- `C`: Switch to current weather view
- `N`: Search for new city
- `Esc`: Go back
- `Ctrl+C` or `Q`: Quit

## 🎨 Visual Design

### Color Palette
- **Primary**: Blue (#3B82F6) - Headers and titles
- **Secondary**: Emerald (#10B981) - Highlights and success
- **Accent**: Amber (#F59E0B) - Icons and special elements
- **Text**: Gray-800 (#1F2937) - Main text
- **Subtle**: Gray-500 (#6B7280) - Secondary text

### Typography
- **Bold headers** for section titles
- **Highlighted values** for temperatures
- **Subtle labels** for data descriptions
- **Monospace alignment** for consistent spacing

### Layout
- **Card-based design** with rounded borders
- **Horizontal layouts** for related information
- **Consistent spacing** and padding
- **Responsive design** that adapts to terminal width

## 🔧 Configuration

### Environment Variables
```bash
WEATHER_API_KEY          # OpenWeatherMap API key
DEFAULT_CITY             # Default city for quick access
DEFAULT_UNITS            # metric, imperial, or kelvin
SHOW_EMOJI               # true/false for weather icons
COLOR_OUTPUT             # true/false for colored output
COMPACT_MODE             # true/false for compact display
```

### API Setup
1. Sign up at [OpenWeatherMap](https://openweathermap.org/api)
2. Get your free API key
3. Set the `WEATHER_API_KEY` environment variable
4. Without an API key, the app uses demo data

## 🛠️ Development

### Project Structure
```
weather-cli/
├── main.go                 # Application entry point
├── internal/
│   ├── commands/          # CLI command handlers
│   │   ├── current.go     # Current weather command
│   │   ├── forecast.go    # Forecast command
│   │   └── interactive.go # Interactive mode command
│   ├── weather/           # Weather service and models
│   │   ├── service.go     # API client
│   │   ├── models.go      # Data structures
│   │   ├── renderer.go    # Beautiful output formatting
│   │   └── mock.go        # Demo data
│   └── tui/               # Terminal UI
│       └── model.go       # Bubble Tea TUI model
├── go.mod                 # Go module definition
└── README.md             # Documentation
```

### Dependencies
- **[Cobra](https://github.com/spf13/cobra)**: CLI framework and command handling
- **[Bubble Tea](https://github.com/charmbracelet/bubbletea)**: TUI framework
- **[Lip Gloss](https://github.com/charmbracelet/lipgloss)**: Terminal styling and layout
- **[Bubbles](https://github.com/charmbracelet/bubbles)**: TUI components

### Building
```bash
# Development build
go build -o weather-cli

# Cross-compilation examples
GOOS=windows GOARCH=amd64 go build -o weather-cli.exe
GOOS=darwin GOARCH=amd64 go build -o weather-cli-mac
GOOS=linux GOARCH=amd64 go build -o weather-cli-linux
```

## 🎯 Design Philosophy

This weather CLI application was designed with these principles:

1. **Beauty First**: Every interface element should be visually appealing
2. **User Experience**: Intuitive navigation and clear information hierarchy
3. **Performance**: Fast responses and efficient API usage
4. **Accessibility**: Works in any terminal with proper fallbacks
5. **Extensibility**: Clean architecture for easy feature additions

## 🆚 Comparison to Spectre.Console

Like Spectre.Console for .NET, this Go application provides:

- ✅ Rich terminal output with colors and styling
- ✅ Component-based UI elements (cards, tables, panels)
- ✅ Responsive layout system
- ✅ Interactive terminal interfaces
- ✅ Beautiful progress indicators and loading states
- ✅ Consistent design language across all outputs

## 🚀 Future Enhancements

Potential features for future versions:

- **Weather Alerts**: Severe weather notifications
- **Historical Data**: Past weather information
- **Weather Maps**: ASCII art weather maps
- **Multiple Locations**: Save and switch between favorite cities
- **Themes**: Customizable color schemes
- **Export Options**: Save weather data to files
- **Weather Widgets**: Mini dashboard view
- **Geolocation**: Auto-detect current location

## 🤝 Contributing

We welcome contributions! Areas where you can help:

- **New Features**: Weather alerts, maps, themes
- **UI Improvements**: Better layouts, animations
- **Performance**: Caching, faster API calls
- **Documentation**: Examples, tutorials, guides
- **Testing**: Unit tests, integration tests
- **Cross-platform**: Windows/macOS/Linux compatibility

See `CONTRIBUTING.md` for development guidelines.
