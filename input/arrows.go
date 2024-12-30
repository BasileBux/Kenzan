package input

import (
	"strings"

	r "github.com/basilebux/kenzan/renderer"
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

// These functions are really baddly written but work good enough

func leftArrow(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	state.Update.Cursor = true
	state.ForceQuit = false
	if state.Nav.AbsoluteSelectedRow >= 1 {
		if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
			state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
		}

		// control + right moves whole words
		if rl.IsKeyDown(rl.KeyLeftControl) {
			jumpTo := strings.LastIndex((*text)[state.Nav.SelectedLine][:state.Nav.AbsoluteSelectedRow-1], " ")
			if jumpTo == -1 {
				state.Nav.AbsoluteSelectedRow = 0
				state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
				state.Nav.ScrollOffset.X = 0
			} else {
				offset := state.Nav.AbsoluteSelectedRow - (jumpTo + 1)
				state.Nav.AbsoluteSelectedRow = jumpTo + 1
				state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
				r.ScrollLeft(offset, state, style)

				if (*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow] == ' ' &&
					(*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow] >= 32 &&
					(*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow] <= 126 {
					for {
						if state.Nav.AbsoluteSelectedRow >= 0 &&
							(*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow] == ' ' {
							state.Nav.AbsoluteSelectedRow--
							state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
							r.ScrollLeft(1, state, style)
						} else {
							state.Nav.AbsoluteSelectedRow++
							state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
							r.ScrollRight(1, state, style)
							break
						}
					}
				}
			}
		} else if state.Nav.AbsoluteSelectedRow > 0 {
			state.Nav.AbsoluteSelectedRow--
			r.ScrollLeft(1, state, style)
		}
	} else if state.Nav.SelectedLine >= 1 {
		// when on left line end, go up end
		state.Nav.SelectedLine--
		state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
		r.ResetHorizontalScrollRight(float32(state.Nav.AbsoluteSelectedRow), state, style)
		r.ScrollUp(1, state, style)
	}
	state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
}

func rightArrow(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	state.Update.Cursor = true
	state.ForceQuit = false
	if state.Nav.AbsoluteSelectedRow < len((*text)[state.Nav.SelectedLine]) {
		if state.Nav.AbsoluteSelectedRow > len((*text)[state.Nav.SelectedLine]) {
			state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
		}

		// control + right moves whole words
		if rl.IsKeyDown(rl.KeyLeftControl) {
			jumpTo := strings.Index((*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow+1:], " ")
			if jumpTo == -1 {
				offset := len((*text)[state.Nav.SelectedLine]) - state.Nav.AbsoluteSelectedRow
				state.Nav.AbsoluteSelectedRow = len((*text)[state.Nav.SelectedLine])
				r.ScrollRight(offset, state, style)
				state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
			} else {
				offset := jumpTo + state.Nav.AbsoluteSelectedRow + 1 - state.Nav.AbsoluteSelectedRow
				state.Nav.AbsoluteSelectedRow = jumpTo + state.Nav.AbsoluteSelectedRow + 1
				r.ScrollRight(offset, state, style)
				state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
				for {
					if (*text)[state.Nav.SelectedLine][state.Nav.AbsoluteSelectedRow] == ' ' {
						if state.Nav.AbsoluteSelectedRow < len((*text)[state.Nav.SelectedLine])-1 {
							state.Nav.AbsoluteSelectedRow++
							r.ScrollRight(1, state, style)
							state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
						} else {
							state.Nav.AbsoluteSelectedRow++
							r.ScrollRight(1, state, style)
							state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
							break
						}
					} else {
						break
					}
				}
			}
		} else {
			state.Nav.AbsoluteSelectedRow++
			state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
			r.ScrollRight(1, state, style)
		}
	} else if state.Nav.SelectedLine < len((*text))-1 {
		// when on right line end, go down and 0
		state.Nav.SelectedLine++
		state.Nav.AbsoluteSelectedRow = 0
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow

		// going to begining of next line so reset X scroll offset and scroll down
		state.Nav.ScrollOffset.X = 0
		r.ScrollDown(1, state, style)
		state.Nav.SelectedRow = state.Nav.AbsoluteSelectedRow
	}
}
