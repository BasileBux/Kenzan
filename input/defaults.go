package input

import rl "github.com/gen2brain/raylib-go/raylib"

func InitDefaultKeyMaps() []Input {
	keyMaps := []Input{
		{ModKey: rl.KeyLeftControl, Key: rl.KeyS, Execute: save},

		{ModKey: -1, Key: rl.KeyBackspace, Execute: backspace},

		{ModKey: -1, Key: rl.KeyDelete, Execute: deleteAction},

		{ModKey: -1, Key: rl.KeyEnter, Execute: enter},

		{ModKey: -1, Key: rl.KeyTab, Execute: tab},

		{ModKey: -1, Key: rl.KeyLeft, Execute: leftArrow},

		{ModKey: -1, Key: rl.KeyRight, Execute: rightArrow},

		{ModKey: -1, Key: rl.KeyUp, Execute: moveUp},

		{ModKey: -1, Key: rl.KeyDown, Execute: moveDown},
	}
	return keyMaps
}
