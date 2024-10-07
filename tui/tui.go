package tui

import (
	"fmt"
	"log"

	"github.com/charmbracelet/bubbles/textinput"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/per1cycle/miui/tui/constants"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)


type Database struct {
	db *gorm.DB
}

func openSqlite() (*gorm.DB, error) {
	db, err := gorm.Open(sqlite.Open("miui.db"), &gorm.Config{})
	if err != nil {
		return db, fmt.Errorf("unable to open database: %w", err)
	}
	return db, nil
}

func Run() {

}

func Debug() {
	m := InitialProfileUrlInputModel()
	constants.TeaProgram = tea.NewProgram(m, tea.WithAltScreen())
	if _, err := constants.TeaProgram.Run(); err != nil {
		log.Fatal(err)
	}
}


func (m ProfileUrlInputModel) Init() tea.Cmd {
	return textinput.Blink
}

func (m ProfileUrlInputModel) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var cmd tea.Cmd

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.Type {
		case tea.KeyEnter, tea.KeyCtrlC, tea.KeyEsc:
			return m, tea.Quit
		}

	// We handle errors just like any other message
	case error:
		m.err = msg
		return m, nil
	}

	m.textInput, cmd = m.textInput.Update(msg)
	return m, cmd
}


func (m ProfileUrlInputModel) View() string {
	return fmt.Sprintf(
		m.textInput.View(),
	) + "\n"
}

