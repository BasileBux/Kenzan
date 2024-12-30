package input

import (
	r "github.com/basilebux/kenzan/renderer"
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
)

// Behaves exactly like "x" in vim
func erase(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	state.Update.SyntaxHighlight = true
	state.SaveState = false
	state.ForceQuit = false
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
		state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
	}
	if state.Nav.AbsoluteSelectedRow <= 0 {
		return
	}
	(*text)[state.Nav.SelectedLine] = (*text)[state.Nav.SelectedLine][:state.Nav.AbsoluteSelectedRow-1] +
		(*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow:]

	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
		state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
	}
	r.ScrollLeft(1, state, style)
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
}

/*
*  BUG: complete freeze in ../syntax_highlighting_samples/main.c on line 14
*		when erasing ']' in "*argv[])" -> only when there is '[' on the left of it and
*		no ')' on the right of it.

*	#DEBUGGING PROCESS:
*		- Not inside of erase() or backspace() function
 */
func backspace(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	if len((*text)[state.Nav.SelectedLine]) <= 0 {
		state.Nav.AbsoluteSelectedRow = 0
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
	}
	if state.Nav.AbsoluteSelectedRow <= 0 {
		if state.Nav.SelectedLine <= 0 {
			return
		}
		state.Update.SyntaxHighlight = true
		state.SaveState = false
		state.ForceQuit = false
		currentLine := (*text)[state.Nav.SelectedLine]
		begin := (*text)[:state.Nav.SelectedLine]
		end := (*text)[state.Nav.SelectedLine+1:]
		*text = begin
		*text = append(*text, end...)
		state.Nav.SelectedLine--
		state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
		(*text)[state.Nav.SelectedLine] += currentLine
		r.ResetHorizontalScrollRight(float32(state.Nav.AbsoluteSelectedRow), state, style)
		r.ScrollUp(1, state, style)
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
		return
	}
	erase(text, state, style)
	state.Nav.AbsoluteSelectedRow--
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
}

func deleteAction(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	if state.Nav.AbsoluteSelectedRow >= len((*text)[state.Nav.SelectedLine]) {
		if state.Nav.SelectedLine >= len(*text)-1 {
			return
		}
		state.Update.SyntaxHighlight = true
		state.SaveState = false
		state.ForceQuit = false
		lineToMove := (*text)[state.Nav.SelectedLine+1]
		begin := (*text)[:state.Nav.SelectedLine+1]
		end := (*text)[state.Nav.SelectedLine+2:]
		*text = begin
		*text = append(*text, end...)
		(*text)[state.Nav.SelectedLine] += lineToMove
		return
	}
	endOfLine := state.Nav.AbsoluteSelectedRow+1 >= len((*text)[state.Nav.SelectedLine])
	state.Nav.AbsoluteSelectedRow++
	erase(text, state, style)
	if !endOfLine {
		state.Nav.AbsoluteSelectedRow--
	}
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
}
