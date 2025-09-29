package tui

import (
	"fmt"
	"time"

	"github.com/ashish0kumar/typtea/internal/game"

	tea "github.com/charmbracelet/bubbletea"
)

// Model represents the state of the typing test application
type Model struct {
	game        *game.TypingGame
	width       int
	height      int
	showResults bool
	finalStats  game.TypingStats
	duration    int
	language    string
	tabPressed  bool
}

// tickMsg is a message type used to handle periodic updates in the application
type tickMsg time.Time

// NewModel initializes a new Model instance with the specified duration and language
func NewModel(duration int, language string) (*Model, error) {
	if err := game.SetLanguage(language); err != nil {
		return nil, fmt.Errorf("failed to load language '%s': %v", language, err)
	}

	return &Model{
		game:     game.NewTypingGame(duration),
		duration: duration,
		language: language,
	}, nil
}

// restartTest resets the game state for a new typing test session
func (m *Model) restartTest() {
	m.game = game.NewTypingGame(m.duration)
	m.showResults = false
	m.finalStats = game.TypingStats{}
	m.tabPressed = false
}

// Init initializes the model and starts the tick command for periodic updates
func (m Model) Init() tea.Cmd {
	return tea.Batch(
		tickCmd(),
		tea.EnterAltScreen,
	)
}

// tickCmd returns a command that sends a tick message every 100 milliseconds
func tickCmd() tea.Cmd {
	return tea.Tick(100*time.Millisecond, func(t time.Time) tea.Msg {
		return tickMsg(t)
	})
}
