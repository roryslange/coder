package interfaces

import (
	"fmt"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type simplePage struct {
	text string
}

func SimpleProgram(message string) simplePage {
	return simplePage{text: message}
}

func (s simplePage) Init() tea.Cmd {return nil}

func (s simplePage) View() string {
    textLen := len(s.text)
    topAndBottomBar := strings.Repeat("*", textLen + 4)
    return fmt.Sprintf(
        "%s\n* %s *\n%s\n\nPress Ctrl+C to exit",
        topAndBottomBar, s.text, topAndBottomBar,
    )
}

func (s simplePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyMsg:
        switch msg.(tea.KeyMsg).String() {
        case "ctrl+c":
            return s, tea.Quit
        }
    }
    return s, nil
}