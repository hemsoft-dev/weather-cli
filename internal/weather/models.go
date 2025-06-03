package weather

import "time"

// CurrentWeather represents current weather conditions
type CurrentWeather struct {
	City        string    `json:"city"`
	Country     string    `json:"country"`
	Temperature float64   `json:"temperature"`
	FeelsLike   float64   `json:"feels_like"`
	Humidity    int       `json:"humidity"`
	Pressure    int       `json:"pressure"`
	Visibility  int       `json:"visibility"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	WindSpeed   float64   `json:"wind_speed"`
	WindDir     int       `json:"wind_direction"`
	CloudCover  int       `json:"cloud_cover"`
	UV          float64   `json:"uv_index"`
	Sunrise     time.Time `json:"sunrise"`
	Sunset      time.Time `json:"sunset"`
	Timestamp   time.Time `json:"timestamp"`
}

// ForecastDay represents weather forecast for a single day
type ForecastDay struct {
	Date        time.Time `json:"date"`
	TempMin     float64   `json:"temp_min"`
	TempMax     float64   `json:"temp_max"`
	Description string    `json:"description"`
	Icon        string    `json:"icon"`
	Humidity    int       `json:"humidity"`
	WindSpeed   float64   `json:"wind_speed"`
	CloudCover  int       `json:"cloud_cover"`
	ChanceRain  int       `json:"chance_rain"`
}

// Forecast represents multi-day weather forecast
type Forecast struct {
	City    string        `json:"city"`
	Country string        `json:"country"`
	Days    []ForecastDay `json:"days"`
}

// OpenWeatherResponse represents the API response structure
type OpenWeatherResponse struct {
	Name string `json:"name"`
	Sys  struct {
		Country string `json:"country"`
		Sunrise int64  `json:"sunrise"`
		Sunset  int64  `json:"sunset"`
	} `json:"sys"`
	Main struct {
		Temp      float64 `json:"temp"`
		FeelsLike float64 `json:"feels_like"`
		Humidity  int     `json:"humidity"`
		Pressure  int     `json:"pressure"`
	} `json:"main"`
	Weather []struct {
		Main        string `json:"main"`
		Description string `json:"description"`
		Icon        string `json:"icon"`
	} `json:"weather"`
	Wind struct {
		Speed float64 `json:"speed"`
		Deg   int     `json:"deg"`
	} `json:"wind"`
	Clouds struct {
		All int `json:"all"`
	} `json:"clouds"`
	Visibility int   `json:"visibility"`
	Dt         int64 `json:"dt"`
}

// OpenWeatherForecastResponse represents the forecast API response
type OpenWeatherForecastResponse struct {
	City struct {
		Name    string `json:"name"`
		Country string `json:"country"`
	} `json:"city"`
	List []struct {
		Dt   int64 `json:"dt"`
		Main struct {
			TempMin  float64 `json:"temp_min"`
			TempMax  float64 `json:"temp_max"`
			Humidity int     `json:"humidity"`
		} `json:"main"`
		Weather []struct {
			Main        string `json:"main"`
			Description string `json:"description"`
			Icon        string `json:"icon"`
		} `json:"weather"`
		Wind struct {
			Speed float64 `json:"speed"`
		} `json:"wind"`
		Clouds struct {
			All int `json:"all"`
		} `json:"clouds"`
		Pop float64 `json:"pop"` // Probability of precipitation
	} `json:"list"`
}

// WeatherIcon maps condition codes to Unicode weather symbols
var WeatherIcon = map[string]string{
	"01d": "☀️",  // clear sky day
	"01n": "🌙",  // clear sky night
	"02d": "⛅",  // few clouds day
	"02n": "☁️",  // few clouds night
	"03d": "☁️",  // scattered clouds
	"03n": "☁️",  // scattered clouds
	"04d": "☁️",  // broken clouds
	"04n": "☁️",  // broken clouds
	"09d": "🌧️", // shower rain
	"09n": "🌧️", // shower rain
	"10d": "🌦️", // rain day
	"10n": "🌧️", // rain night
	"11d": "⛈️",  // thunderstorm
	"11n": "⛈️",  // thunderstorm
	"13d": "🌨️", // snow
	"13n": "🌨️", // snow
	"50d": "🌫️", // mist
	"50n": "🌫️", // mist
}
