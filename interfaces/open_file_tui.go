package interfaces

import (
	"bufio"
	"os"

	tea "github.com/charmbracelet/bubbletea"
)

type openFilePage struct {
	filePath string
	file *os.File
	readWriter *bufio.ReadWriter
}

func OpenFile(path string) *openFilePage {
	return &openFilePage{filePath: path}
}

func (o *openFilePage) Init() tea.Cmd { 
	var err error
	o.file, err = os.OpenFile(o.filePath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}

	o.readWriter = bufio.NewReadWriter(bufio.NewReader(o.file), bufio.NewWriter(o.file))


	return nil 
}

func (o *openFilePage) View() string {
	return ""
}

func (o *openFilePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg.(type) {
    case tea.KeyMsg:
        switch msg.(tea.KeyMsg).String() {
        case "ctrl+c":
			e := o.file.Close()
			if e != nil {
				panic(e)
			}

            return o, tea.Quit
        }
    }
    return o, nil
}
