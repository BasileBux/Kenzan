package fonts

// Implementation for this isn't done with Core Text which would be the correct way.
// instead, I use mdls and the implementation is really weak but it's ok to me.

import (
	"bytes"
	"encoding/xml"
	"fmt"
	"os"
	"os/exec"
	"path"
	"strings"
)

type Plist struct {
	XMLName xml.Name `xml:"plist"`
	Dict    Dict     `xml:"dict"`
}

type Dict struct {
	Keys    []string    `xml:"key"`
	Strings []string    `xml:"string"`
	Arrays  []ArrayNode `xml:"array"`
}

type ArrayNode struct {
	Strings []string `xml:"string"`
}

func darwinFonts(fontName string) (string, error) {
	fonts, err := getDarwinFontList(fontName)
	if err != nil || len(fonts) == 0 {
		return "", fmt.Errorf("Couldn't find your font. Using fallback instead")
	}
	if len(fonts) > 1 {
		for _, f := range fonts {
			if f.Name == fontName {
				return f.Path, nil
			}
		}
		regularFonts := findFontStyle(fonts, "Regular")
		if len(regularFonts) == 1 {
			return regularFonts[0].Path, nil
		}

		for _, f := range regularFonts {
			if strings.Contains(f.Path, "regular") || strings.Contains(f.Path, "Regular") {
				return f.Path, nil
			}
		}

		shortest := path.Base(regularFonts[0].Path)
		minId := 0
		for i, file := range regularFonts {
			current := path.Base(file.Path)
			if len(current) < len(shortest) {
				shortest = current
				minId = i
			}
		}
		return regularFonts[minId].Path, nil
	}

	return fonts[0].Path, nil
}

func getDarwinFontList(fontName string) ([]font, error) {
	homeFolder, err := os.UserHomeDir()
	if err != nil {
		panic("Could not find home folder")
	}
	userFonts := path.Join(homeFolder, "/Library/Fonts")
	cmd := exec.Command("find", "/System/Library/Fonts", "/System/Library/Fonts/Supplemental",
		"/Library/Fonts", userFonts, "-type", "f")
	var out bytes.Buffer
	cmd.Stdout = &out
	err = cmd.Run()
	if err != nil {
		return nil, fmt.Errorf("Couldn't find your font. Using fallback instead")
	}

	lines := strings.Split(out.String(), "\n")
	var fonts []font
	for _, l := range lines {
		if (strings.TrimSpace(path.Ext(l)) == ".ttf" || strings.TrimSpace(path.Ext(l)) == ".otf") &&
			strings.Contains(l, strings.Split(fontName, " ")[0]) {
			curFont, err := getDarwinFontInfo(l)
			if err == nil {
				fonts = append(fonts, curFont)
			}
		}
	}
	return fonts, nil
}

func getDarwinFontInfo(fontPath string) (font, error) {
	cmd := exec.Command("mdls", fontPath, "-p", "-")
	var out bytes.Buffer
	cmd.Stdout = &out
	err := cmd.Run()
	if err != nil {
		return font{}, fmt.Errorf("Could not find any info on this font")
	}

	xmlData := out.Bytes()

	var plist Plist
	err = xml.Unmarshal([]byte(xmlData), &plist)
	if err != nil {
		return font{}, err
	}

	var fullname string
	var types []string

	for i, key := range plist.Dict.Keys {
		switch key {
		case "com_apple_ats_name_full":
			if i-1 < len(plist.Dict.Arrays) {
				fullname = plist.Dict.Arrays[i-1].Strings[0] // I don't understand why it's i-1
			}

		case "com_apple_ats_name_style":
			if i-1 < len(plist.Dict.Arrays) {
				types = plist.Dict.Arrays[i-1].Strings // I don't understand why it's i-1
			}
		}
	}

	return font{Name: fullname, Path: fontPath, Styles: types}, nil
}
