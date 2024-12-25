package fonts

import (
	"fmt"
	"path"
	"runtime"
	"strings"

	u "github.com/basileb/kenzan/utils"
)

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
		// Error -> use fallback
		fmt.Println("Using fallback")
		fmt.Println("File ext: ", path.Ext(fontPath))
		return fontName
	}

	return fontPath
}
