package model

import (
	"log"
	"draco/internal/config"
	"strings"
	"fmt"
)

type Stage int

const (
	StageMenu Stage = iota
	StageCommand
	StageArgument
	StageConfirm
	StageOutput
	StageDone
)

type Model struct {
	Config          config.Config
	MenuOptions     []string
	SelectedMenu    string
	Cursor          int
	SelectedCommand int
	CurrentStage    Stage
	SelectedArgs    []string
	SelectedPath    string
}

func InitialModel(conf config.Config) Model {
	// Extract the menu options from the config
	var menuOptions []string
	for menu := range conf.Menus {
		menuOptions = append(menuOptions, menu)
	}
	log.Println("Initializing model.")
	return Model{
		Config:       conf,
		MenuOptions:  menuOptions,
		Cursor:       0,
		CurrentStage: StageMenu,
	}
}

// GetCurrentOptions returns the list of options available for the current stage
func (m *Model) GetCurrentOptions() []string {
	switch m.CurrentStage {
	case StageMenu:
		return m.MenuOptions
	case StageCommand:
		var commands []string
		for _, cmd := range m.Config.Menus[m.SelectedMenu] {
			commands = append(commands, cmd.Name)
		}
		return commands
	case StageArgument:
		return m.Config.Menus[m.SelectedMenu][m.SelectedCommand].Arguments
	default:
		return nil
	}
}

// IsSelectingPath checks if the current stage involves selecting a path
func (m *Model) IsSelectingPath() bool {
	return len(m.SelectedArgs) >= len(m.Config.Menus[m.SelectedMenu][m.SelectedCommand].Arguments) &&
		len(m.Config.Menus[m.SelectedMenu][m.SelectedCommand].Paths) > 0
}

func (m *Model) BuildCommandString() string {
    // Check if the SelectedMenu is valid
    if len(m.MenuOptions) == 0 || m.SelectedMenu == "" {
        return ""
    }

    commands := m.Config.Menus[m.SelectedMenu]

    // Check if the SelectedCommand is valid
    if len(commands) == 0 || m.SelectedCommand < 0 || m.SelectedCommand >= len(commands) {
        return ""
    }

    command := commands[m.SelectedCommand].Name

    if len(m.SelectedArgs) == 0 {
        return command // Just the command name if no arguments selected
    }

    // Concatenate command name and selected arguments
    return fmt.Sprintf("%s %s", command, strings.Join(m.SelectedArgs, " "))
}
