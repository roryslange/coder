package interfaces

import (
	tea "github.com/charmbracelet/bubbletea"
)

type openFilePage struct {
	filePath string
}

func (o openFilePage) Init() tea.Cmd { return nil }

func (o openFilePage) View() string { return "file" }

func (o openFilePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyMsg:
        switch msg.(tea.KeyMsg).String() {
        case "ctrl+c":
            return o, tea.Quit
        }
    }
    return o, nil
}
