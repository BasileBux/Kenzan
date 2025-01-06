package input

import (
	"math"
	"strings"

	r "github.com/basilebux/kenzan/renderer"
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
	u "github.com/basilebux/kenzan/utils"
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
	offsetTab(-1, text, state)
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
	offsetTab(1, text, state)
	state.Nav.SelectedLine++
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine])-1 {
		state.Nav.SelectedRow = len((*text)[state.Nav.SelectedLine])
	} else {
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
	}

	r.ScrollDown(1, state, style)
}

// Correct navigation up and down on tab symbols
func offsetTab(increment int, text *[]string, state *t.ProgramState) {
	if len((*text)[state.Nav.SelectedLine]) <= 0 {
		return
	}
	end := state.Nav.AbsoluteSelectedRow
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine])-1 {
		end = len((*text)[state.Nav.SelectedLine])
	}
	tabNbCurrentLine := strings.Count((*text)[state.Nav.SelectedLine][:end], "\t")
	if tabNbCurrentLine > 0 {
		state.Nav.AbsoluteSelectedRow += (state.Indent.Size - 1) * tabNbCurrentLine
	}
	end = state.Nav.AbsoluteSelectedRow
	if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine+increment]) {
		end = len((*text)[state.Nav.SelectedLine+increment])
	}
	tabNbNextLine := strings.Count((*text)[state.Nav.SelectedLine+increment][:end], "\t")
	if tabNbNextLine > 0 {
		tabPos := u.FindAllIndices((*text)[state.Nav.SelectedLine+increment][:end], "\t")
		found := -1
		for i, pos := range tabPos {
			offsetSelected := state.Nav.AbsoluteSelectedRow - (i * (state.Indent.Size - 1))
			if offsetSelected > pos && offsetSelected <= pos+state.Indent.Size {
				state.Nav.AbsoluteSelectedRow = pos + 1
				found = i
				break
			}
		}
		var calcOffsets []int
		if found < 0 {
			newSel := -1
			for i := range len(tabPos) {
				calcOffsets = append(calcOffsets, state.Nav.AbsoluteSelectedRow-int(math.Abs(float64((state.Indent.Size-1)*(i)))))
				newSel = calcOffsets[i]
				if calcOffsets[i] <= tabPos[i] {
					found = 1
					break
				}
			}
			if found < 0 { // No clue but this works
				newSel -= state.Indent.Size - 1
			}
			state.Nav.AbsoluteSelectedRow = newSel
		}
	}
	if state.Nav.AbsoluteSelectedRow < 0 {
		state.Nav.AbsoluteSelectedRow = 0
	}
}
