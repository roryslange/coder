package interfaces

import (
	"bufio"
	"io"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/bubbles/viewport"
)

type openFilePage struct {
	filePath string
	file *os.File
	lines []string
	readWriter *bufio.ReadWriter
	Viewport viewport.Model
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

	for {
		line, err := o.readWriter.Reader.ReadString('\n')
		if err != nil && err != io.EOF {
            panic(err)
        }
		if len(line) > 0 {		
			o.lines = append(o.lines, line)
		}
		if err == io.EOF {
			break
		}
	}

	return nil 
}

func (o *openFilePage) View() string {
	return o.Viewport.View()
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
	case tea.WindowSizeMsg:
		o.Viewport = viewport.New(msg.Width, msg.Height-1)
		o.Viewport.YPosition = 0
		o.updateViewport()
    }
    return o, nil
}

func (o *openFilePage) updateViewport() {
	var buf strings.Builder
	for _, line := range o.lines {
		buf.WriteString(line)
		buf.WriteByte('\n')
	}
	o.Viewport.SetContent(buf.String())
}