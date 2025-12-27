package interfaces

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
)

type openFilePage struct {
	filePath string
	file *os.File
	lines []string
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
	for {
		line, err := o.readWriter.Reader.ReadString('\n')
		if err != nil && err != io.EOF {
            return fmt.Sprintf("error reading file: %s", err.Error())
        }
		if len(line) > 0 {		
			o.lines = append(o.lines, line)
		}
		if err == io.EOF {
			break
		}
	}
	return strings.Join(o.lines, "")
}

func (o *openFilePage) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
    switch msg := msg.(type) {
    case tea.KeyMsg:
        switch msg.Type {
        case tea.KeyCtrlC:
			e := o.file.Close()
			if e != nil {
				panic(e)
			}
            return o, tea.Quit
        }
    }
    return o, nil
}