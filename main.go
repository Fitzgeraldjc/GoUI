package main

import (
	"log"
	"os"

	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

var (
	titleStyle = lipgloss.NewStyle().
			Bold(true).
			Foreground(lipgloss.Color("cyan")).
			Padding(0, 1)

	helpStyle = lipgloss.NewStyle().
			Foreground(lipgloss.Color("241")).
			Italic(true)

	viewportStyle = lipgloss.NewStyle().
			Border(lipgloss.RoundedBorder()).
			BorderForeground(lipgloss.Color("62")).
			Padding(1, 2)
)

const filename = "book.txt"

func main() {
	p := tea.NewProgram(InitialModel())
	if _, err := p.Run(); err != nil {
		log.Fatal(err)
	}
}

type model struct {
	viewport viewport.Model
}

func ParseFile(filename string) ([]byte, error) {
	f, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return f, nil
}

func InitialModel() model {
	content, err := ParseFile(filename)
	if err != nil {
		panic(err)
	}
	vp := viewport.New(1, 1)
	vp.SetContent(string(content))
	return model{
		viewport: vp,
	}
}

func (m model) Init() tea.Cmd {
	return nil
}

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "q", "ctrl+c", "esc":
			return m, tea.Quit
		}

	case tea.WindowSizeMsg:
		headerHeight := lipgloss.Height(titleStyle.Render("Peanits Reader"))
		helpHeight := lipgloss.Height(helpStyle.Render("↑/↓: scroll • q: quit"))
		verticalFrameSize := 4
		horizontalFrameSize := 6
		m.viewport.Width = msg.Width - horizontalFrameSize
		m.viewport.Height = msg.Height - headerHeight - helpHeight - verticalFrameSize
	}

	m.viewport, cmd = m.viewport.Update(msg)
	return m, cmd
}

func (m model) View() string {
	title := titleStyle.Render("Peanits Reader")

	viewport := viewportStyle.Render(m.viewport.View())
	help := helpStyle.Render("↑/↓: scroll • q: quit")
	return lipgloss.JoinVertical(
		lipgloss.Left,
		title,
		viewport,
		help,
	)
}
