package fonts

import (
	"os"
	"path"
	"runtime"
	"slices"
	"strings"

	u "github.com/basilebux/kenzan/utils"
)

const FALLBACK_FONT string = "kenzan/fonts/JetBrainsMono-Medium.ttf"

type font struct {
	Path   string
	Name   string
	Styles []string
}

func GetFontPath(fontName string) string {
	if u.FileExists(fontName) {
		return fontName // fontName is a file path
	}
	var fontPath string
	if fontName == "" {
		fontPath, err := os.UserConfigDir()
		if err != nil {
			panic("Could not find os config file")
		}
		return path.Join(fontPath, FALLBACK_FONT)
	}
	var err error
	switch runtime.GOOS {
	case "linux":
		fontPath, err = linuxFonts(fontName)
	case "darwin":
		fontPath, err = darwinFonts(fontName)
	case "windows":
		fontPath, err = windowsFonts(fontName)
	}
	fileExt := strings.TrimSpace(path.Ext(fontPath))
	if err != nil || (fileExt != ".otf" && fileExt != ".ttf") {
		// Using fallback
		fontPath, err = os.UserConfigDir()
		if err != nil {
			panic("Could not find os config file")
		}
		return path.Join(fontPath, FALLBACK_FONT)
	}

	return fontPath
}

func findFontStyle(fonts []font, style string) []font {
	var foundFonts []font
	for _, f := range fonts {
		if slices.Contains(f.Styles, style) {
			foundFonts = append(foundFonts, f)
		}
	}
	return foundFonts
}
