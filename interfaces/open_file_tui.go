package interfaces

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type openFilePage struct {
	filePath string
	file os.File
}

func OpenFile(path string) openFilePage {
	return openFilePage{filePath: path}
}

func (o openFilePage) Init() tea.Cmd { 
	file, err := os.OpenFile(o.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	o.file = *file
	fmt.Println("yup file is here: " + o.filePath)
	return nil 
}

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
