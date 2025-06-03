#!/bin/bash

echo ""
echo "==========================================="
echo "Weather CLI Demo - Beautiful Terminal App"
echo "==========================================="
echo ""

echo "1. Current Weather (Demo Data):"
echo ""
./weather-cli current "New York"

echo ""
echo ""
echo "2. 5-Day Forecast (Demo Data):"
echo ""
./weather-cli forecast "Tokyo"

echo ""
echo ""
echo "3. To try the interactive mode, run:"
echo "   ./weather-cli interactive"
echo ""
echo "4. To use with real weather data:"
echo "   export WEATHER_API_KEY=your-openweathermap-api-key"
echo ""
