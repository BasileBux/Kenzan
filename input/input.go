package input

import (
	st "github.com/basilebux/kenzan/settings"
	t "github.com/basilebux/kenzan/types"
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Input struct {
	ModKey  int32 // set to -1 if no mod key
	Key     int32
	Execute func(text *[]string, state *t.ProgramState, style *st.WindowStyle)
}

func (i *Input) set(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	if i.check() {
		i.Execute(text, state, style)
	}
}

func (i *Input) check() bool {
	if i.ModKey != -1 {
		return rl.IsKeyDown(i.ModKey) && rl.IsKeyPressed(i.Key)
	}
	return rl.IsKeyPressedRepeat(i.Key) || rl.IsKeyPressed(i.Key)
}

func (i *Input) exec(text *[]string, state *t.ProgramState, style *st.WindowStyle) {
	i.Execute(text, state, style)
}

func InputManager(text *[]string, keyMaps []Input, state *t.ProgramState, style *st.WindowStyle) {
	textInput(text, state, style)
	for _, i := range keyMaps {
		i.set(text, state, style)
	}
}
