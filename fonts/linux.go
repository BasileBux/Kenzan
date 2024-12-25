package fonts

import (
	"bytes"
	"fmt"
	"os/exec"
	"strings"
)

func linuxFonts(fontName string) (string, error) {
	fonts, err := getLinuxFontList(fontName)
	if err != nil || len(fonts) == 0 {
		// fallback
		return "", fmt.Errorf("Couldn't find your font. Using fallback instead")
	}
	if len(fonts) == 1 {
		return fonts[0].Path, nil
	}
	return getLinuxDefaultFont(fonts)
}

// Rules for what a default font is are in docs/fonts.md
func getLinuxDefaultFont(fonts []font) (string, error) {
	regularFonts := findFontStyle(fonts, "Regular")
	if len(regularFonts) == 1 {
		return regularFonts[0].Path, nil
	}
	if len(regularFonts) > 1 {
		for _, f := range regularFonts {
			if len(f.Styles) == 1 {
				return f.Path, nil
			}
		}
		regularMediumFonts := findFontStyle(regularFonts, "Medium")
		if len(regularMediumFonts) == 1 {
			return regularMediumFonts[0].Path, nil
		}
	}

	var mediumFonts []font
	mediumFonts = findFontStyle(fonts, "Medium")
	if len(mediumFonts) != 1 && len(regularFonts) == 0 {
		return findLinuxDefaultFileName(fonts)
	} else {
		return mediumFonts[0].Path, nil
	}
}

func findLinuxDefaultFileName(fonts []font) (string, error) {
	var validFonts []font
	for _, f := range fonts {
		if strings.Contains(f.Path, "Regular") {
			validFonts = append(validFonts, f)
		}
	}
	if len(validFonts) == 1 {
		return validFonts[0].Path, nil
	}
	validFonts = nil
	for _, f := range fonts {
		if strings.Contains(f.Path, "Medium") {
			validFonts = append(validFonts, f)
		}
	}
	if len(validFonts) == 1 {
		return validFonts[0].Path, nil
	}
	return fonts[0].Path, fmt.Errorf("Could not determine default font. Using something.")
}

func getLinuxFontList(fontName string) ([]font, error) {
	cmd := exec.Command("fc-list")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Couldn't find your font. Using fallback instead")
	}

	lines := strings.Split(out.String(), "\n")
	var fonts []font

	for _, line := range lines {
		if strings.Contains(line, fontName) {
			parts := strings.Split(line, ":")
			if len(parts) < 3 {
				continue
			}

			filePath := strings.TrimSpace(parts[0])
			name := strings.TrimSpace(strings.Split(parts[1], ",")[0])
			styles := strings.Split(strings.TrimSpace(parts[2][len("style="):]), ",")

			if name == fontName {
				fonts = append(fonts, font{
					Path:   filePath,
					Name:   name,
					Styles: styles,
				})
			}
		}
	}
	return fonts, nil
}
