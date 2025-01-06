package renderer

import (
	"strings"

	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

func CalculateCursorPos(userText []string, nav *t.NavigationData, indentSize int, cache *t.Cache, userStyle *st.WindowStyle) {
	// If we're on an empty line or newline, place cursor at the start
	if len(userText[nav.SelectedLine]) <= 0 || userText[nav.SelectedLine] == "\n" ||
		nav.AbsoluteSelectedRow == 0 {
		cursorVerticalPos := int32(userStyle.PaddingTop) +
			int32(nav.SelectedLine)*int32(userStyle.FontSize) +
			int32(nav.SelectedLine*int(userStyle.FontSpacing)) -
			int32(nav.ScrollOffset.Y*float32(userStyle.FontSize+userStyle.FontSpacing))

		cache.Cursor.X = int(userStyle.PaddingLeft)
		cache.Cursor.Y = int(cursorVerticalPos)
		return
	}

	// Get the text up to the cursor position
	textBeforeCursor := userText[nav.SelectedLine][:nav.SelectedRow]

	// Replace '\t' with spaces for rendering
	tabRender := strings.Repeat(" ", indentSize)
	textBeforeCursor = strings.Replace(textBeforeCursor, "\t", tabRender, -1)

	// Measure the exact width of the text before the cursor
	cursorPos := rl.MeasureTextEx(
		userStyle.Font,
		textBeforeCursor,
		userStyle.FontSize,
		userStyle.FontSpacing,
	)

	// Calculate positions
	cursorHorizontalPos := int32(userStyle.PaddingLeft) +
		int32(cursorPos.X) -
		int32(nav.ScrollOffset.X*float32(cursorPos.X/float32(max(1, len(textBeforeCursor)))))

	cursorVerticalPos := int32(userStyle.PaddingTop) +
		int32(nav.SelectedLine)*int32(userStyle.FontSize) +
		int32(nav.SelectedLine*int(userStyle.FontSpacing)) -
		int32(nav.ScrollOffset.Y*float32(userStyle.FontSize+userStyle.FontSpacing))

	cache.Cursor.X = int(cursorHorizontalPos)
	cache.Cursor.Y = int(cursorVerticalPos)
}
