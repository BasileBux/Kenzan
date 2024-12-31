package input

import (
	"strings"

	r "github.com/basilebux/kenzan/renderer"
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func textInput(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	char := rl.GetCharPressed()
	for char > 0 {
		// refuse non ascii and non printable chars
		if char >= 32 && char <= 126 {
			state.SaveState = false
			state.ForceQuit = false
			state.Update.SyntaxHighlight = true
			if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
				state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
			}
			if state.Nav.AbsoluteSelectedRow < len((*text)[state.Nav.SelectedLine]) {
				(*text)[state.Nav.SelectedLine] = (*text)[state.Nav.SelectedLine][:state.Nav.AbsoluteSelectedRow] +
					string(rune(char)) + (*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow:]
				state.Nav.AbsoluteSelectedRow++
				state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
			} else {
				(*text)[state.Nav.SelectedLine] += string(rune(char))
				state.Nav.AbsoluteSelectedRow++
				state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow

				// Scroll right if needed
				r.ScrollRight(1, state, style)
			}
		}
		char = rl.GetCharPressed()
	}
}

func enter(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	if state.ForceQuit {
		state.Terminate = true
		return
	}
	state.Update.SyntaxHighlight = true
	state.SaveState = false
	state.Update.Cursor = true

	newText := make([]string, len(*text)+1)
	copy(newText, (*text)[:state.Nav.SelectedLine+1])
	copy(newText[state.Nav.SelectedLine+2:], (*text)[state.Nav.SelectedLine+1:])
	*text = newText

	remainingString := (*text)[state.Nav.SelectedLine][state.Nav.SelectedRow:]
	if len(remainingString) != 0 {
		(*text)[state.Nav.SelectedLine] = (*text)[state.Nav.SelectedLine][:state.Nav.SelectedRow]
		(*text)[state.Nav.SelectedLine+1] += remainingString
	}

	state.Nav.SelectedLine++
	state.Nav.AbsoluteSelectedRow = 0
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow

	r.ScrollDown(1, state, style)
	state.Nav.ScrollOffset.X = 0

	// Check and change width of line numbers gutter if max nb changes
	r.UpdateLineNumWidth(len(*text), state, style)
}

func tab(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	state.Update.SyntaxHighlight = true
	state.SaveState = false
	state.ForceQuit = false
	begin := (*text)[state.Nav.SelectedLine][:state.Nav.SelectedRow]
	end := (*text)[state.Nav.SelectedLine][state.Nav.SelectedRow:]
	(*text)[state.Nav.SelectedLine] = begin + strings.Repeat(" ", 4) + end
	state.Nav.AbsoluteSelectedRow += 4
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
	r.ScrollRight(4, state, style)
}
