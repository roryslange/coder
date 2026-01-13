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
	file *os.File
}

var CURSOR_CHARACTER rune = '\u258C'

func OpenModel(filepath string) *openModel {
	return &openModel{path: filepath}
}

func (m *openModel) Init() tea.Cmd {
	m.lines = m.initFileContents(m.path)
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
				m.closeFile()
				return m, tea.Quit
			case tea.KeyRight:
				m.cursor.column = min(m.cursor.column + 1, len(m.lines[m.cursor.row]) - 1)
				m.updateViewport()
			case tea.KeyLeft:
				m.cursor.column = max(m.cursor.column - 1, 0)
				m.updateViewport()
			case tea.KeyUp:
				if m.cursor.row > 0 {
					m.cursor.row--
					m.cursor.column = min(m.cursor.column, len(m.lines[m.cursor.row]) - 1)
				}
				m.updateViewport()
			case tea.KeyDown:
				if m.cursor.row <= len(m.lines) {
					m.cursor.row++
					m.cursor.column = min(m.cursor.column, len(m.lines[m.cursor.row]) - 1)
				}
				m.updateViewport()
			}
		case tea.WindowSizeMsg:
			m.viewport = viewport.New(msg.Width, msg.Height - 1)
			m.viewport.YOffset = 0
			m.updateViewport()
		}
	return m, nil
}

func (m *openModel) initFileContents(filepath string) [][]rune {
	var err error
	m.file, err = os.OpenFile(filepath, os.O_RDWR|os.O_CREATE, 0644)
	if err != nil {
		panic(err)
	}
	reader := bufio.NewReader(m.file)
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
	return buffer
}

func (m *openModel) updateViewport() {
	var buf strings.Builder
	for i, line := range m.lines {
		if i == m.cursor.row {

			line[m.cursor.column] = CURSOR_CHARACTER
		}
		buf.WriteString(string(line))
	}

	m.viewport.SetContent(buf.String())
}

func (m *openModel) closeFile() {
	err := m.file.Close()
	if err != nil {
		panic(err)
	}
}