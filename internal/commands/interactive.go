package commands

import (
	"fmt"

	"github.com/charmbracelet/bubbletea"
	"github.com/spf13/cobra"
	"weather-cli/internal/tui"
)

var InteractiveCmd = &cobra.Command{
	Use:   "interactive",
	Short: "Launch interactive weather app",
	Long:  "Launch a beautiful interactive terminal UI for browsing weather information",
	Run: func(cmd *cobra.Command, args []string) {
		// Create and run the TUI application
		model := tui.NewModel()
		program := tea.NewProgram(model, tea.WithAltScreen())
		
		if _, err := program.Run(); err != nil {
			fmt.Printf("Error running interactive mode: %v\n", err)
		}
	},
}
