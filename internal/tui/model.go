package tui

import (
	"fmt"
	"strings"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
	"weather-cli/internal/weather"
)

// Model represents the state of our TUI application
type Model struct {
	textInput     textinput.Model
	weather       *weather.CurrentWeather
	forecast      *weather.Forecast
	weatherSvc    *weather.Service
	renderer      *weather.Renderer
	loading       bool
	err           error
	mode          string // "input", "current", "forecast"
	city          string
}

type weatherMsg struct {
	current  *weather.CurrentWeather
	forecast *weather.Forecast
	err      error
}

// NewModel creates a new TUI model
func NewModel() Model {
	ti := textinput.New()
	ti.Placeholder = "Enter city name..."
	ti.Focus()
	ti.CharLimit = 50
	ti.Width = 30

	return Model{
		textInput:  ti,
		weatherSvc: weather.NewService(),
		renderer:   weather.NewRenderer(),
		mode:       "input",
	}
}

// Init initializes the model
func (m Model) Init() tea.Cmd {
	return textinput.Blink
}

// Update handles messages and updates the model
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "enter":
			if m.mode == "input" && m.textInput.Value() != "" {
				m.city = strings.TrimSpace(m.textInput.Value())
				m.loading = true
				m.mode = "current"
				return m, m.fetchWeather()
			}
		case "f":
			if m.mode == "current" && m.weather != nil {
				m.mode = "forecast"
				m.loading = true
				return m, m.fetchForecast()
			}
		case "c":
			if m.mode == "forecast" && m.weather != nil {
				m.mode = "current"
			}
		case "n":
			if m.mode == "current" || m.mode == "forecast" {
				m.mode = "input"
				m.textInput.SetValue("")
				m.weather = nil
				m.forecast = nil
				m.err = nil
			}
		case "esc":
			if m.mode != "input" {
				m.mode = "input"
				m.textInput.SetValue("")
				m.weather = nil
				m.forecast = nil
				m.err = nil
			}
		}

	case weatherMsg:
		m.loading = false
		if msg.err != nil {
			m.err = msg.err
		} else {
			if msg.current != nil {
				m.weather = msg.current
			}
			if msg.forecast != nil {
				m.forecast = msg.forecast
			}
		}
		return m, nil
	}

	// Update text input
	if m.mode == "input" {
		m.textInput, cmd = m.textInput.Update(msg)
	}

	return m, cmd
}

// View renders the TUI
func (m Model) View() string {
	var s strings.Builder

	// Header
	headerStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#3B82F6")).
		Bold(true).
		Padding(1, 2).
		Border(lipgloss.RoundedBorder()).
		BorderForeground(lipgloss.Color("#3B82F6")).
		Margin(1, 0)

	header := headerStyle.Render("🌤️  Interactive Weather CLI")
	s.WriteString(header)
	s.WriteString("\n\n")

	switch m.mode {
	case "input":
		s.WriteString(m.renderInput())
	case "current":
		s.WriteString(m.renderCurrent())
	case "forecast":
		s.WriteString(m.renderForecast())
	}

	// Instructions
	s.WriteString("\n")
	s.WriteString(m.renderInstructions())

	return s.String()
}

func (m Model) renderInput() string {
	var s strings.Builder
	
	promptStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280")).
		Margin(1, 0)
	
	s.WriteString(promptStyle.Render("Enter a city name to get weather information:"))
	s.WriteString("\n\n")
	s.WriteString(m.textInput.View())
	s.WriteString("\n")
	
	return s.String()
}

func (m Model) renderCurrent() string {
	var s strings.Builder
	
	if m.loading {
		loadingStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Bold(true)
		s.WriteString(loadingStyle.Render("🔄 Loading weather data..."))
		return s.String()
	}
	
	if m.err != nil {
		errorStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#EF4444")).
			Bold(true)
		s.WriteString(errorStyle.Render(fmt.Sprintf("❌ Error: %v", m.err)))
		return s.String()
	}
	
	if m.weather != nil {
		s.WriteString(m.renderer.RenderCurrent(m.weather))
	}
	
	return s.String()
}

func (m Model) renderForecast() string {
	var s strings.Builder
	
	if m.loading {
		loadingStyle := lipgloss.NewStyle().
			Foreground(lipgloss.Color("#10B981")).
			Bold(true)
		s.WriteString(loadingStyle.Render("🔄 Loading forecast data..."))
		return s.String()
	}
	
	if m.forecast != nil {
		s.WriteString(m.renderer.RenderForecast(m.forecast))
	}
	
	return s.String()
}

func (m Model) renderInstructions() string {
	instructionStyle := lipgloss.NewStyle().
		Foreground(lipgloss.Color("#6B7280")).
		Border(lipgloss.NormalBorder()).
		BorderForeground(lipgloss.Color("#E5E7EB")).
		Padding(1).
		Margin(1, 0)
	
	var instructions string
	switch m.mode {
	case "input":
		instructions = "• Press Enter to search\n• Press Ctrl+C or Q to quit"
	case "current":
		instructions = "• Press F for forecast\n• Press N for new city\n• Press Esc to go back\n• Press Ctrl+C or Q to quit"
	case "forecast":
		instructions = "• Press C for current weather\n• Press N for new city\n• Press Esc to go back\n• Press Ctrl+C or Q to quit"
	}
	
	return instructionStyle.Render("📖 Instructions:\n" + instructions)
}

// fetchWeather fetches current weather data
func (m Model) fetchWeather() tea.Cmd {
	return func() tea.Msg {
		current, err := m.weatherSvc.GetCurrent(m.city)
		return weatherMsg{current: current, err: err}
	}
}

// fetchForecast fetches forecast data
func (m Model) fetchForecast() tea.Cmd {
	return func() tea.Msg {
		forecast, err := m.weatherSvc.GetForecast(m.city)
		return weatherMsg{forecast: forecast, err: err}
	}
}
