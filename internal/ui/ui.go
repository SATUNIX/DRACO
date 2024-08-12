package ui

import (
	"bytes"
	"fmt"
	"log"
	"os/exec"
	"strings"

	"draco/internal/config"
	"draco/internal/model"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)
/*
 *
 // * TO ALL YE WHO FOLLOW, HERE IS A FROG FOR YOUR TROUBLES / TRAVELS

    _____
   /     \______
  | o     |     \____
  /\_____/           \___
 /                       \
|_______/                 \
  \______   _       ___    \
        /\_//      /   \    |
       // //______/    /___/
      /\/\/\      \   / \ \
                    \ \   \ \
                      \ \   \ \
                        \ \  \ \
                         \ \ /\/\
                         /\/\

 */
// Model wraps the state and handles user input and rendering.
type Model struct {
	Model  model.Model
	Styles Styles
	Output string // Store the output of the command
	Error  error  // Store any error from the command
	Width  int    // Terminal width
	Height int    // Terminal height
}

// Styles holds the styles for the TUI components.
type Styles struct {
	CursorStyle    lipgloss.Style
	SelectedStyle  lipgloss.Style
	TextStyle      lipgloss.Style
	BorderStyle    lipgloss.Style
	ErrorStyle     lipgloss.Style
	HighlightStyle lipgloss.Style // New style for highlighting selected text
}

// NewModel creates a new UI model with styles.
func NewModel(m model.Model, conf config.Config) Model {
	styles := Styles{
		CursorStyle:    lipgloss.NewStyle().Foreground(lipgloss.Color(conf.ColorPalette.CursorColor)),
		SelectedStyle:  lipgloss.NewStyle().Foreground(lipgloss.Color(conf.ColorPalette.SelectedColor)),
		TextStyle:      lipgloss.NewStyle().Foreground(lipgloss.Color(conf.ColorPalette.TextColor)),
		ErrorStyle:     lipgloss.NewStyle().Foreground(lipgloss.Color("#FF0000")), // Red for errors
		HighlightStyle: lipgloss.NewStyle().Background(lipgloss.Color("#404040")), // Grey background for highlight
	}
	log.Println("MULTINOMIAL ANALYSIS EXECUTE\n\tBEGIN INTERFACE\n\t\tNEURAL LINK INITIALISE\n =============================")
	return Model{Model: m, Styles: styles}
}


// Init initializes the TUI with no command.
func (m Model) Init() tea.Cmd {
	return tea.EnterAltScreen // Ensure the app takes over the terminal screen
}

// Update updates the state based on user input.
func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			log.Println("God's in his heaven, all is right with the world....")
			return m, tea.Quit
		case "up":
			if m.Model.CurrentStage != model.StageOutput && m.Model.Cursor > 0 {
				m.Model.Cursor--
				log.Println("Going Up!")
			}
		case "down":
			if m.Model.CurrentStage != model.StageOutput && m.Model.Cursor < len(m.Model.GetCurrentOptions())-1 {
				m.Model.Cursor++
				log.Println("G0ing Down!.")
			}
		case "right":
			m.handleRightKey()
		case "left":
			m.handleLeftKey()
		case "enter":
			m.handleEnterKey()
		case "ctrl+left":
			m.handleCtrlLeftKey() // New
		}
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		m.Styles.BorderStyle = lipgloss.NewStyle().Border(lipgloss.NormalBorder(), true).Padding(1, 2).Align(lipgloss.Left).Width(m.Width - 4).Height(m.Height - 4).BorderForeground(m.Styles.CursorStyle.GetForeground())
		log.Printf("DIMENSION MODIFY: %dx%d\n", m.Width, m.Height)
	}

	return m, nil
}

// handleCtrlLeftKey clears all selected arguments
func (m *Model) handleCtrlLeftKey() {
	m.Model.SelectedArgs = nil
	log.Println("ALL ARGS CLEARED")
}

func (m *Model) handleLeftKey() {
	if m.Model.CurrentStage != model.StageOutput {
		if m.Model.CurrentStage == model.StageArgument {
			m.Model.CurrentStage = model.StageCommand
			log.Println("AWAITING ORDERS")
		} else if m.Model.CurrentStage == model.StageCommand {
			m.Model.CurrentStage = model.StageMenu
			log.Println("WEAPONS MATRIX RE-INITIALISE")
		}
		m.Model.Cursor = 0
	}
}


func (m *Model) handleEnterKey() {
	if m.Model.CurrentStage == model.StageArgument && len(m.Model.SelectedArgs) > 0 {
		m.Model.CurrentStage = model.StageConfirm
	}
}

func (m Model) View() string {
	var s strings.Builder

	currentCommand := m.Model.BuildCommandString()
	maxLabelWidth := len("SELECT MENU")

	s.WriteString(fmt.Sprintf("%-*s %s\n", maxLabelWidth, m.Styles.TextStyle.Render("CURRENT COMMAND:"), currentCommand))

	switch m.Model.CurrentStage {
	case model.StageMenu:
		m.renderMenu(&s, maxLabelWidth)
	case model.StageCommand:
		m.renderCommands(&s, maxLabelWidth)
	case model.StageArgument:
		m.renderArguments(&s, maxLabelWidth)
	case model.StageConfirm:
		m.renderConfirmation(&s)
	case model.StageOutput:
		m.renderOutput(&s)
	}

	paddingLeft := (m.Width - lipgloss.Width(s.String())) / 4 // Adjust this to position more to the left side

	if paddingLeft < 0 {
		paddingLeft = 0
	}

	return m.Styles.BorderStyle.Render(fmt.Sprintf("%s%s", strings.Repeat(" ", paddingLeft), s.String()))
}

func (m *Model) handleRightKey() {
	switch m.Model.CurrentStage {
	case model.StageMenu:
		m.Model.SelectedMenu = m.Model.MenuOptions[m.Model.Cursor]
		m.Model.CurrentStage = model.StageCommand
		m.Model.Cursor = 0
		log.Printf("MENU SELECTED: %s\n", m.Model.SelectedMenu)
	case model.StageCommand:
		m.Model.SelectedCommand = m.Model.Cursor
		m.Model.CurrentStage = model.StageArgument
		m.Model.Cursor = 0
		log.Printf("COMMAND SELECTED: %s\n", m.Model.Config.Menus[m.Model.SelectedMenu][m.Model.SelectedCommand].Name)
	case model.StageArgument:
		currentOptions := m.Model.GetCurrentOptions()
		if len(currentOptions) > 0 {
			m.Model.SelectedArgs = append(m.Model.SelectedArgs, currentOptions[m.Model.Cursor])
			log.Printf("ARG ADDED: %s\n", currentOptions[m.Model.Cursor])
		}
	case model.StageConfirm:
		cmdStr := m.Model.BuildCommandString()
		log.Printf("EXECUTING: %s\n", cmdStr)
		output, err := executeCommandInShell(cmdStr)
		m.handleCtrlLeftKey() // Yeah yeah i know... what are you gonna do about it ?
		m.Output = output
		m.Error = err
		if err != nil {
			log.Printf("MISS: %v\n", err)
		} else {
			log.Println("ORDER EXECUTED, BOMBS AWAY!")
		}
		m.Model.CurrentStage = model.StageOutput
	case model.StageOutput:
		m.Model.CurrentStage = model.StageMenu
		m.Model.Cursor = 0
		log.Println("RETURNING TO THE MATRIX")
	}
}

func (m *Model) renderMenu(s *strings.Builder, maxLabelWidth int) {
	s.WriteString(fmt.Sprintf("%-*s\n", maxLabelWidth, m.Styles.TextStyle.Render("SELECTION INTERFACE ONE")))
	for i, menu := range m.Model.MenuOptions {
		cursor := "  "
		if m.Model.Cursor == i {
			cursor = m.Styles.CursorStyle.Render(">>>")
			menu = m.Styles.HighlightStyle.Render(m.Styles.SelectedStyle.Render(menu)) // Highlight selected menu
		} else {
			menu = m.Styles.TextStyle.Render(menu)
		}
		s.WriteString(fmt.Sprintf("%s %s\n", cursor, menu))
	}
}

func (m *Model) renderCommands(s *strings.Builder, maxLabelWidth int) {
	s.WriteString(fmt.Sprintf("%-*s\n", maxLabelWidth, m.Styles.TextStyle.Render(":")))
	for i, cmd := range m.Model.Config.Menus[m.Model.SelectedMenu] {
		cursor := "  "
		if m.Model.Cursor == i {
			cursor = m.Styles.CursorStyle.Render(">>>")
			cmdName := fmt.Sprintf("%-20s - %s", cmd.Name, cmd.Description)
			cmdName = m.Styles.HighlightStyle.Render(m.Styles.SelectedStyle.Render(cmdName))
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, cmdName))
		} else {
			cmdName := fmt.Sprintf("%-20s - %s", cmd.Name, cmd.Description)
			s.WriteString(fmt.Sprintf("%s %s\n", cursor, cmdName))
		}
	}
}

func (m *Model) renderArguments(s *strings.Builder, maxLabelWidth int) {
	s.WriteString(fmt.Sprintf("%-*s %s\n", maxLabelWidth, m.Styles.TextStyle.Render("SELECTED C0MMAND [^ ̮ ^]"), m.Model.Config.Menus[m.Model.SelectedMenu][m.Model.SelectedCommand].Name))
	s.WriteString(m.Styles.TextStyle.Render("\n ARROW KEYS TO SELECT \n"))
	for i, arg := range m.Model.Config.Menus[m.Model.SelectedMenu][m.Model.SelectedCommand].Arguments {
		cursor := "  "
		if m.Model.Cursor == i {
			cursor = m.Styles.CursorStyle.Render(">>>")
			arg = m.Styles.HighlightStyle.Render(m.Styles.SelectedStyle.Render(arg))
		} else {
			arg = m.Styles.TextStyle.Render(arg)
		}
		s.WriteString(fmt.Sprintf("%s %s\n", cursor, arg))
	}
	s.WriteString(m.Styles.TextStyle.Render("\n\n [^ ̮ ^] \n\nINSTRUCTIONS: \n RIGHT = ADD ARG \n ENTER = CONFIRM \n LEFT = BACK \n CTRL+LEFT = CLEAR ARGS \n CTRL + C EXIT PROGRAM"))
}

func (m *Model) renderConfirmation(s *strings.Builder) {
	cmdStr := m.Model.BuildCommandString()
	s.WriteString(m.Styles.TextStyle.Render("CONFIRM COMMAND?:\n"))
	s.WriteString(m.Styles.SelectedStyle.Render(cmdStr))
	s.WriteString(m.Styles.TextStyle.Render("\nRIGHT = EXECUTE \n LEFT = BACK\n"))
}

func (m *Model) renderOutput(s *strings.Builder) {
	if m.Error != nil {
		s.WriteString(m.Styles.ErrorStyle.Render(fmt.Sprintf("ERROR: %v\n", m.Error)))
	} else {
		s.WriteString(m.Styles.TextStyle.Render(fmt.Sprintf("OUTPUT:\n%s\n", m.Output)))
	}
	s.WriteString(m.Styles.TextStyle.Render("\nRIGHT = MAIN MENU\n"))

}

func executeCommandInShell(command string) (string, error) {
	var outBuf, errBuf bytes.Buffer
	cmd := exec.Command("bash", "-c", command)

	cmd.Stdout = &outBuf
	cmd.Stderr = &errBuf

	err := cmd.Run()
	if err != nil {
		return errBuf.String(), err
	}
	return outBuf.String(), nil
}
