package interfaces

import (
	"bufio"
	"io"
	"os"
	"strings"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

type cursor struct {
	row int
	column int
}

type openModel struct {
	path string
	lines [][]rune
	cursor cursor
	viewport viewport.Model
}

func OpenModel(filepath string) *openModel {
	return &openModel{path: filepath}
}

func (m *openModel) Init() tea.Cmd {
	m.lines = initFileContents(m.path)
	m.cursor = cursor{0,0}
	return nil
}

func (m *openModel) View() string {
	return m.viewport.View()
}

func (m *openModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
		case tea.KeyMsg:
			switch msg.Type {
			case tea.KeyCtrlC:
				return m, tea.Quit
			}
		case tea.WindowSizeMsg:
			m.viewport = viewport.New(msg.Width, msg.Height - 1)
			m.viewport.YOffset = 0
			m.updateViewport()
		}
	return m, nil
}

func initFileContents(filepath string) [][]rune {
	file, err := os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(file)
	var buffer [][]rune

	for {
		line, err := reader.ReadString('\n')
		if err != nil && err != io.EOF {
			panic(err)
		}
		if len(line) > 0 {
			buffer = append(buffer, []rune(line))
		}
		if err == io.EOF {
			break
		}
	}

	err = file.Close()
	if err != nil {
		panic(err)
	}

	return buffer
}

func (m *openModel) updateViewport() {
	var buf strings.Builder
	for _, line := range m.lines {
		buf.WriteString(string(line))
	}
	m.viewport.SetContent(buf.String())
}