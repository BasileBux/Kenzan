package fonts

import "fmt"

// This will probably be implemented one day but right now, just use a complete path to your font
func windowsFonts(fontName string) (string, error) {
	fmt.Println("Font fetching is not implemented for windows yet. Using fallback or give full path in settings.")
	return fontName, nil
}
