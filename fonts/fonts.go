package fonts

import (
	"os"
	"path"
	"runtime"
	"strings"

	u "github.com/basileb/kenzan/utils"
)

const FALLBACK_FONT string = "kenzan/fonts/GeistMono-Regular.otf"

func GetFontPath(fontName string) string {
	if u.FileExists(fontName) {
		return fontName // fontName is a file path
	}
	var fontPath string
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
