package fonts

import (
	"runtime"

	u "github.com/basileb/kenzan/utils"
)

func GetFontPath(fontName string) string {
	if u.FileExists(fontName) {
		return fontName // fontName is a file path
	}
	var fontMap map[string]string
	switch runtime.GOOS {
	case "linux":
		fontMap = linuxFonts()
	case "darwin":
		fontMap = darwinFonts()
	case "windows":
		fontMap = windowsFonts()
	}
	if fontMap == nil {
		return fontName // This is weird
	}
	path, found := fontMap[fontName]
	if !found {
		return ""
	}
	return path
}
