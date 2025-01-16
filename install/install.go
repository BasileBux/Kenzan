package main

import (
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
)

func isCommandAvailable(command string) bool {
	_, err := exec.LookPath(command)
	return err == nil
}

func printErr(s string) {
	fmt.Println("\033[31mError: " + s + "\033[0m")
}

func createDir(path string) error {
	return os.MkdirAll(path, 0755)
}

func moveFonts(sourceDir, destDir string) error {
	if err := os.MkdirAll(destDir, os.ModePerm); err != nil {
		return fmt.Errorf("could not create destination directory: %w", err)
	}

	files, err := os.ReadDir(sourceDir)
	if err != nil {
		return fmt.Errorf("could not read source directory: %w", err)
	}

	for _, file := range files {
		sourceFilePath := filepath.Join(sourceDir, file.Name())
		destFilePath := filepath.Join(destDir, file.Name())

		if err := os.Rename(sourceFilePath, destFilePath); err != nil {
			return fmt.Errorf("could not move file %s to %s: %w", sourceFilePath, destFilePath, err)
		}
	}

	return nil
}

func main() {

	// Check if font finding util is available. If not, exit.
	switch runtime.GOOS {
	case "linux":
		if !isCommandAvailable("fc-list") {
			printErr("fc-list not found. You need that in order for the software to work." +
				"\nTo install on:\n" + "Debian/Ubuntu: sudo apt install fontconfig\n" +
				"Arch: sudo pacman -S fontconfig\n" + "Fedora: sudo dnf install fontconfig\n" +
				"OpenSUSE: sudo zypper install fontconfig" + "\nFor other distros, check your package manager.")
			return
		}
	case "darwin":
		if !isCommandAvailable("mdfind") {
			printErr("mdfind not found. You need that in order for the software to work. It should be available by default.")
			return
		}

		// No windows as font discovery is not implemented there yet.
	}
	fmt.Print("Font discovery utility found successfuly\n")

	configDir, err := os.UserConfigDir()
	if err != nil {
		printErr("Could not find os config directory")
		return
	}
	cacheDir, err := os.UserCacheDir()
	if err != nil {
		printErr("Could not find os cache directory")
		return
	}
	homeDir, err := os.UserHomeDir()
	if err != nil {
		printErr("Could not find os home directory")
		return
	}

	// create config and cache directories
	err = createDir(filepath.Join(configDir, "kenzan"))
	if err != nil {
		printErr("Could not create config directories")
		return
	}
	fmt.Println("Created config directory at:", filepath.Join(configDir, "kenzan"), " successfully")
	err = createDir(filepath.Join(cacheDir, "kenzan"))
	if err != nil {
		printErr("Could not create config/cache directories")
		return
	}
	fmt.Println("Created cache directory at:", filepath.Join(cacheDir, "kenzan"), " successfully")

	// move fonts to config dir
	fontsSourceDir := filepath.Join("fonts", "JetBrainsMono")
	fontsDestDir := filepath.Join(configDir, "kenzan", "Font")

	if err := moveFonts(fontsSourceDir, fontsDestDir); err != nil {
		printErr(fmt.Sprintf("Failed to move font files: %v", err))
		return
	}
	fmt.Println("Moved fonts to:", fontsDestDir, " successfully")

	// create empty config file `.json`
	file, err := os.Create(filepath.Join(configDir, "kenzan", "settings.json"))
	if err != nil {
		printErr(fmt.Sprintf("Could not create file: %v", err))
		return
	}
	defer file.Close()
	fmt.Println("Created config file at:", filepath.Join(configDir, "kenzan", "settings.json"), " successfully")

	// create empty cache file `.json`
	file, err = os.Create(filepath.Join(cacheDir, "kenzan", "cache.json"))
	if err != nil {
		printErr(fmt.Sprintf("Could not create file: %v", err))
		return
	}
	defer file.Close()
	fmt.Println("Created cache file at:", filepath.Join(cacheDir, "kenzan", "cache"), " successfully")

	// compile in right dir
	var outputPath string
	switch runtime.GOOS {
	case "darwin": // macOS
		outputPath = filepath.Join(homeDir, "bin", "kenzan")
	case "linux":
		outputPath = filepath.Join(homeDir, ".local", "bin", "kenzan")
	case "windows":
		outputPath = filepath.Join(os.Getenv("ProgramFiles"), "kenzan", "kenzan.exe")
	default:
		printErr("Unsupported OS for binary installation")
		return
	}

	cmd := exec.Command("go", "build", "-o", outputPath, "-ldflags=-w -s", ".")
	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	if err := cmd.Run(); err != nil {
		printErr(fmt.Sprintf("Failed to build the application: %v", err))
		return
	}

	fmt.Println("Build successful! Binary installed at:", outputPath, " make sure it in your PATH")

	fmt.Println("Application is now installed! You can run it by typing `kenzan` in your terminal.")
}
