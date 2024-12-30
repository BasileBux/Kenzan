package input

import (
	r "github.com/basilebux/kenzan/renderer"
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
)

func moveLeft(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	state.Update.Cursor = true
	state.ForceQuit = false
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
		state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
	}
	if state.Nav.AbsoluteSelectedRow >= 1 {
		state.Nav.AbsoluteSelectedRow--
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
		r.ScrollLeft(1, state, style)
	}
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
}

func moveRight(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	state.Update.Cursor = true
	state.ForceQuit = false
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
		state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
	}
	if state.Nav.AbsoluteSelectedRow < len((*text)[state.Nav.SelectedLine]) {
		state.Nav.AbsoluteSelectedRow++
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
		r.ScrollRight(1, state, style)
	}
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
}

func moveUp(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	if state.Nav.SelectedLine <= 0 {
		return
	}
	state.Update.Cursor = true
	state.ForceQuit = false
	state.Nav.SelectedLine--
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine])-1 {
		state.Nav.SelectedRow = len((*text)[state.Nav.SelectedLine])
	} else {
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
	}

	r.ScrollUp(1, state, style)
}

func moveDown(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	if state.Nav.SelectedLine >= len(*text)-1 {
		return
	}
	state.Update.Cursor = true
	state.ForceQuit = false
	state.Nav.SelectedLine++
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine])-1 {
		state.Nav.SelectedRow = len((*text)[state.Nav.SelectedLine])
	} else {
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
	}

	r.ScrollDown(1, state, style)
}
