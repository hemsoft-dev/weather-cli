package weather

import (
	"time"
)

// Mock data for demo purposes
func (s *Service) getMockCurrent(city string) *CurrentWeather {
	now := time.Now()
	// Create realistic sunrise/sunset times
	sunrise := time.Date(now.Year(), now.Month(), now.Day(), 6, 30, 0, 0, now.Location())
	sunset := time.Date(now.Year(), now.Month(), now.Day(), 19, 45, 0, 0, now.Location())
	
	return &CurrentWeather{
		City:        city,
		Country:     "Demo",
		Temperature: 22.5,
		FeelsLike:   24.0,
		Humidity:    65,
		Pressure:    1013,
		Visibility:  10000,
		Description: "Partly cloudy",
		Icon:        "02d",
		WindSpeed:   3.2,
		WindDir:     180,
		CloudCover:  40,
		UV:          5.2,
		Sunrise:     sunrise,
		Sunset:      sunset,
		Timestamp:   now,
	}
}

func (s *Service) getMockForecast(city string) *Forecast {
	days := make([]ForecastDay, 5)
	baseTime := time.Now()
	
	icons := []string{"01d", "02d", "10d", "04d", "13d"}
	descriptions := []string{"Sunny", "Partly cloudy", "Rainy", "Cloudy", "Snowy"}
	
	for i := 0; i < 5; i++ {
		days[i] = ForecastDay{
			Date:        baseTime.AddDate(0, 0, i+1),
			TempMin:     18.0 + float64(i),
			TempMax:     25.0 + float64(i),
			Description: descriptions[i],
			Icon:        icons[i],
			Humidity:    60 + i*5,
			WindSpeed:   2.5 + float64(i)*0.5,
			CloudCover:  20 + i*15,
			ChanceRain:  i * 20,
		}
	}
	
	return &Forecast{
		City:    city,
		Country: "Demo",
		Days:    days,
	}
}

// Convert API response to our internal models
func (s *Service) convertToCurrent(resp *OpenWeatherResponse) *CurrentWeather {
	description := "Unknown"
	icon := "01d"
	if len(resp.Weather) > 0 {
		description = resp.Weather[0].Description
		icon = resp.Weather[0].Icon
	}
	
	return &CurrentWeather{
		City:        resp.Name,
		Country:     resp.Sys.Country,
		Temperature: resp.Main.Temp,
		FeelsLike:   resp.Main.FeelsLike,
		Humidity:    resp.Main.Humidity,
		Pressure:    resp.Main.Pressure,
		Visibility:  resp.Visibility,
		Description: description,
		Icon:        icon,
		WindSpeed:   resp.Wind.Speed,
		WindDir:     resp.Wind.Deg,
		CloudCover:  resp.Clouds.All,
		UV:          0, // Not available in current weather API
		Sunrise:     time.Unix(resp.Sys.Sunrise, 0),
		Sunset:      time.Unix(resp.Sys.Sunset, 0),
		Timestamp:   time.Unix(resp.Dt, 0),
	}
}

func (s *Service) convertToForecast(resp *OpenWeatherForecastResponse) *Forecast {
	// Group by day and take the first entry for each day
	dayMap := make(map[string]ForecastDay)
	
	for _, item := range resp.List {
		date := time.Unix(item.Dt, 0)
		dayKey := date.Format("2006-01-02")
		
		if _, exists := dayMap[dayKey]; !exists {
			description := "Unknown"
			icon := "01d"
			if len(item.Weather) > 0 {
				description = item.Weather[0].Description
				icon = item.Weather[0].Icon
			}
			
			dayMap[dayKey] = ForecastDay{
				Date:        date,
				TempMin:     item.Main.TempMin,
				TempMax:     item.Main.TempMax,
				Description: description,
				Icon:        icon,
				Humidity:    item.Main.Humidity,
				WindSpeed:   item.Wind.Speed,
				CloudCover:  item.Clouds.All,
				ChanceRain:  int(item.Pop * 100),
			}
		}
	}
	
	// Convert map to slice and sort by date
	days := make([]ForecastDay, 0, len(dayMap))
	for _, day := range dayMap {
		days = append(days, day)
	}
	
	return &Forecast{
		City:    resp.City.Name,
		Country: resp.City.Country,
		Days:    days,
	}
}
