package tui


import (
	"github.com/charmbracelet/bubbles/textinput"
)

type ProfileUrlInputModel struct {
	textInput textinput.Model
	err       error
}

func InitialProfileUrlInputModel() ProfileUrlInputModel {
	ti := textinput.New()
	ti.Placeholder = "input your sub link here"
	ti.Focus()
	ti.CharLimit = 156
	ti.Width = 20

	return ProfileUrlInputModel{
		textInput: ti,
		err:       nil,
	}
}

