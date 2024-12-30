package renderer

import (
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
)

// Arbitrary padding to allign scroll to side
const UP_PADDING int = -1
const DOWN_PADDING int = 0
const LEFT_PADDING int = -1
const RIGHT_PADDING int = 0

func ResetHorizontalScrollRight(lineSize float32, state *t.ProgramState, style *st.WindowStyle) {
	if lineSize > float32(state.ViewPortSteps.X)-float32(RIGHT_PADDING) {
		state.Nav.ScrollOffset.X = lineSize - float32(state.ViewPortSteps.X) + float32(RIGHT_PADDING) + float32(style.Cursor.HorizontalPadding)
		state.Update.SyntaxHighlight = true
	}
}

func ScrollLeft(size int, state *t.ProgramState, style *st.WindowStyle) {
	if state.Nav.ScrollOffset.X > float32(size-1) {
		if state.Nav.SelectedRow < int(state.Nav.ScrollOffset.X+float32(LEFT_PADDING)+float32(style.Cursor.HorizontalPadding)) {
			state.Nav.ScrollOffset.X -= float32(size)
			state.Update.SyntaxHighlight = true
		}
	} else {
		state.Nav.ScrollOffset.X = 0
		state.Update.SyntaxHighlight = true
	}
}

func ScrollRight(size int, state *t.ProgramState, style *st.WindowStyle) {
	if state.Nav.AbsoluteSelectedRow > int(state.Nav.ScrollOffset.X)+state.ViewPortSteps.X-RIGHT_PADDING-int(style.Cursor.HorizontalPadding) {
		state.Nav.ScrollOffset.X += float32(size)
		state.Update.SyntaxHighlight = true
	}
}

func ScrollUp(size int, state *t.ProgramState, style *st.WindowStyle) {
	if int(state.Nav.ScrollOffset.Y) > (size) {
		if state.Nav.SelectedLine < int(state.Nav.ScrollOffset.Y)+int(style.Cursor.VerticalPadding)+UP_PADDING {
			state.Nav.ScrollOffset.Y -= float32(size)
			state.Update.SyntaxHighlight = true
		}
	} else {
		state.Nav.ScrollOffset.Y = 0
		state.Update.SyntaxHighlight = true
	}
}

func ScrollDown(size int, state *t.ProgramState, style *st.WindowStyle) {
	if state.Nav.SelectedLine > int(state.Nav.ScrollOffset.Y)+state.ViewPortSteps.Y-DOWN_PADDING-int(style.Cursor.VerticalPadding) {
		state.Nav.ScrollOffset.Y += float32(size)
		state.Update.SyntaxHighlight = true
	}
}
