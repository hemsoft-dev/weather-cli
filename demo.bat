@echo off
echo.
echo ===========================================
echo Weather CLI Demo - Beautiful Terminal App
echo ===========================================
echo.

echo 1. Current Weather (Demo Data):
echo.
weather-cli.exe current "New York"

echo.
echo.
echo 2. 5-Day Forecast (Demo Data):
echo.
weather-cli.exe forecast "Tokyo"

echo.
echo.
echo 3. To try the interactive mode, run:
echo    weather-cli.exe interactive
echo.
echo 4. To use with real weather data:
echo    set WEATHER_API_KEY=your-openweathermap-api-key
echo.
pause
