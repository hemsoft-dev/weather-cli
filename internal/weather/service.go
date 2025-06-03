package weather

import (
	"encoding/json"
	"fmt"
	"net/http"
	"net/url"
	"os"
	"time"
)

// Service handles weather API interactions
type Service struct {
	apiKey  string
	baseURL string
	client  *http.Client
}

// NewService creates a new weather service
func NewService() *Service {
	apiKey := os.Getenv("WEATHER_API_KEY")
	if apiKey == "" {
		// For demo purposes, we'll use a mock service
		return &Service{
			apiKey:  "demo",
			baseURL: "https://api.openweathermap.org/data/2.5",
			client:  &http.Client{Timeout: 10 * time.Second},
		}
	}
	
	return &Service{
		apiKey:  apiKey,
		baseURL: "https://api.openweathermap.org/data/2.5",
		client:  &http.Client{Timeout: 10 * time.Second},
	}
}

// GetCurrent gets current weather for a city
func (s *Service) GetCurrent(city string) (*CurrentWeather, error) {
	if s.apiKey == "demo" {
		return s.getMockCurrent(city), nil
	}
	
	endpoint := fmt.Sprintf("%s/weather", s.baseURL)
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", s.apiKey)
	params.Add("units", "metric")
	
	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch weather data: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}
	
	var apiResp OpenWeatherResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	
	return s.convertToCurrent(&apiResp), nil
}

// GetForecast gets weather forecast for a city
func (s *Service) GetForecast(city string) (*Forecast, error) {
	if s.apiKey == "demo" {
		return s.getMockForecast(city), nil
	}
	
	endpoint := fmt.Sprintf("%s/forecast", s.baseURL)
	params := url.Values{}
	params.Add("q", city)
	params.Add("appid", s.apiKey)
	params.Add("units", "metric")
	
	url := fmt.Sprintf("%s?%s", endpoint, params.Encode())
	
	resp, err := s.client.Get(url)
	if err != nil {
		return nil, fmt.Errorf("failed to fetch forecast data: %w", err)
	}
	defer resp.Body.Close()
	
	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("API request failed with status: %d", resp.StatusCode)
	}
	
	var apiResp OpenWeatherForecastResponse
	if err := json.NewDecoder(resp.Body).Decode(&apiResp); err != nil {
		return nil, fmt.Errorf("failed to decode response: %w", err)
	}
	
	return s.convertToForecast(&apiResp), nil
}
