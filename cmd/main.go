package main

import (
	"fmt"
	"log"
	"os"
	"path/filepath"

	"draco/internal/config"
	"draco/internal/model"
	"draco/internal/ui"

	tea "github.com/charmbracelet/bubbletea"
)

func setupLogging() {
	logDir := "/usr/var/log/draco"
	logFile := "draco.log"

	// Create log directory if it doesn't exist
	if _, err := os.Stat(logDir); os.IsNotExist(err) {
		if err := os.MkdirAll(logDir, 0755); err != nil {
			fmt.Fprintf(os.Stderr, "Failed to create log directory: %v\n", err)
			os.Exit(1)
		}
	}

	// Open log file
	logPath := filepath.Join(logDir, logFile)
	logFileHandle, err := os.OpenFile(logPath, os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open log file: %v\n", err)
		os.Exit(1)
	}

	// Set log output to the log file
	log.SetOutput(logFileHandle)
	log.Println("Logging started.")
}

func main() {
	// Set up logging
	setupLogging()

	// Load the configuration
	conf, err := config.LoadConfig("config.json")
	if err != nil {
		log.Fatalf("Config Load Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Config Load Error: %v\n", err)
		os.Exit(1)
	}

	// Initialize the model
	m := model.InitialModel(conf)

	// Start the TUI program
	p := tea.NewProgram(ui.NewModel(m, conf))
	if err := p.Start(); err != nil {
		log.Fatalf("Program Start Error: %v\n", err)
		fmt.Fprintf(os.Stderr, "Program Start Error: %v\n", err)
		os.Exit(1)
	}

	log.Println("Program exited successfully.")
}
