# Weather CLI 🌤️

A beautiful command-line weather application written in Go, featuring rich terminal output similar to Spectre.Console for .NET.

## Features

- 🎨 **Beautiful Terminal UI** - Rich colors, icons, and layouts using Charm libraries
- 🌡️ **Current Weather** - Get current weather conditions for any city
- 📅 **5-Day Forecast** - View detailed weather forecasts
- 🖥️ **Interactive Mode** - Full-screen TUI for browsing weather data
- 🌍 **Global Coverage** - Weather data for cities worldwide
- ⚡ **Fast & Responsive** - Quick API responses with elegant loading states

## Installation

### From Source

```bash
git clone <your-repo-url>
cd weather-cli
go mod tidy
go build -o weather-cli
```

### Using Go Install

```bash
go install github.com/yourusername/weather-cli@latest
```

## Usage

### Command Line Interface

Get current weather:
```bash
weather-cli current "New York"
weather-cli current London
weather-cli current "San Francisco"
```

Get 5-day forecast:
```bash
weather-cli forecast "Tokyo"
weather-cli forecast Paris
```

### Interactive Mode

Launch the beautiful full-screen interface:
```bash
weather-cli interactive
```

In interactive mode:
- Enter city names to search
- Press `F` to switch to forecast view
- Press `C` to switch to current weather
- Press `N` to search for a new city
- Press `Esc` to go back
- Press `Ctrl+C` or `Q` to quit

## Configuration

### API Key Setup

To use real weather data, set up a free API key from [OpenWeatherMap](https://openweathermap.org/api):

1. Sign up at OpenWeatherMap
2. Get your free API key
3. Set the environment variable:

```bash
# Windows (PowerShell)
$env:WEATHER_API_KEY = "your-api-key-here"

# Windows (Command Prompt)
set WEATHER_API_KEY=your-api-key-here

# macOS/Linux
export WEATHER_API_KEY="your-api-key-here"
```

Without an API key, the app will use demo data for testing.

## Screenshots

### Current Weather Display
```
┌─────────────────────────────────────────────────┐
│  🌤️  Current Weather - London, GB              │
└─────────────────────────────────────────────────┘

          ┌──────────────────────────────────────┐
          │              ⛅  22.5°C              │
          │            Partly cloudy             │
          │          Feels like 24.0°C           │
          └──────────────────────────────────────┘

┌──────────────────┐ ┌────────────────────┐ ┌──────────────────┐
│ 💨 Wind          │ │ 🌡️  Atmosphere     │ │ 🌅 Sun Times     │
│ Speed: 3.2 m/s   │ │ Humidity: 65%      │ │ Sunrise: 6:23 AM │
│ Direction: S     │ │ Pressure: 1013 hPa │ │ Sunset: 8:45 PM  │
│                  │ │ Clouds: 40%        │ │                  │
└──────────────────┘ └────────────────────┘ └──────────────────┘
```

### 5-Day Forecast
```
┌─────────────────────────────────────────────────┐
│  📅  5-Day Forecast - London, GB               │
└─────────────────────────────────────────────────┘

┌────────────┐ ┌────────────┐ ┌────────────┐ ┌────────────┐ ┌────────────┐
│    Mon     │ │    Tue     │ │    Wed     │ │    Thu     │ │    Fri     │
│   Dec 4    │ │   Dec 5    │ │   Dec 6    │ │   Dec 7    │ │   Dec 8    │
│ ☀️  Sunny   │ │ ⛅ Cloudy  │ │ 🌧️ Rainy   │ │ ☁️  Overcast│ │ 🌨️ Snowy   │
│  25° / 18° │ │  26° / 19° │ │  27° / 20° │ │  28° / 21° │ │  29° / 22° │
│ 💨 2.5 m/s │ │ 💨 3.0 m/s │ │ 💨 3.5 m/s │ │ 💨 4.0 m/s │ │ 💨 4.5 m/s │
│            │ │            │ │ 🌧️ 40%     │ │            │ │ 🌧️ 80%     │
└────────────┘ └────────────┘ └────────────┘ └────────────┘ └────────────┘
```

## Dependencies

- [Cobra](https://github.com/spf13/cobra) - CLI framework
- [Bubble Tea](https://github.com/charmbracelet/bubbletea) - TUI framework
- [Lip Gloss](https://github.com/charmbracelet/lipgloss) - Terminal styling
- [Bubbles](https://github.com/charmbracelet/bubbles) - TUI components

## Development

### Build
```bash
go build -o weather-cli
```

### Run Tests
```bash
go test ./...
```

### Run with Demo Data
```bash
# No API key needed - uses mock data
./weather-cli current "Demo City"
./weather-cli forecast "Demo City"
./weather-cli interactive
```

## Contributing

1. Fork the repository
2. Create a feature branch
3. Commit your changes
4. Push to the branch
5. Create a Pull Request

## License

MIT License - see LICENSE file for details.

## Acknowledgments

- Weather data provided by [OpenWeatherMap](https://openweathermap.org/)
- Beautiful terminal output powered by [Charm](https://charm.sh/) libraries
- Inspired by [Spectre.Console](https://spectreconsole.net/) for .NET
