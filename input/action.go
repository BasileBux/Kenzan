package input

import (
	"fmt"

	f "github.com/basilebux/kenzan/files"
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
)

func save(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	err := f.WriteFile(state.AcitveFile, *text)
	if err != nil {
		fmt.Println("Couldn't save file")
	} else {
		state.SaveState = true
		state.SavedFile = make([]string, len(*text))
		copy(state.SavedFile, *text)
	}
}
