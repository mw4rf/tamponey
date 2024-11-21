package main

import (
	"fmt"
	"os"
	"regexp"

	"github.com/pdfcpu/pdfcpu/pkg/api"
	"github.com/pdfcpu/pdfcpu/pkg/pdfcpu/types"
	"github.com/spf13/viper"
)


func main() {
	// Get the stamp template and configuration from the configuration file
	configDir, err := os.UserConfigDir()
	if err != nil {
		error("Error getting user config directory: " + err.Error())
	}

	viper.SetConfigName("config")
	// viper.SetConfigType("yaml")
	viper.AddConfigPath(configDir + "/tamponey")

	if err := viper.ReadInConfig(); err != nil {
		error("Error reading configuration file: " + err.Error())
	}
	// Get the stamp template and configuration
	// for configuration, see https://pdfcpu.io/core/watermark
	stampTmpl := viper.GetString("stamp.template")
	stampCfg := viper.GetString("stamp.configuration")

	// Find all PDF files in the directory given by the first command-line argument or, if none, in the current working directory.
	if len(os.Args) > 1 {
		err := os.Chdir(os.Args[1])
		if err != nil {
			error("Error changing directory: " + err.Error())
		}
	}
	files, err := os.ReadDir(".")
	if err != nil {
		error("Error reading directory: " + err.Error())
	}

	// Loop over all files in the directory. Print the name of each file and whether it is a directory or a regular file.
	for _, file := range files {
		// Exclude directories
		if file.IsDir() {
			continue
		}
		// Exclude non-PDF files
		if file.Name()[len(file.Name())-4:] != ".pdf" {
			continue
		}

		// Extract the number at the start of the filename (e.g. '12 my file.pdf' -> 12)
		number := ""
		re := regexp.MustCompile(`^\d+`)
		number = re.FindString(file.Name())

		// If it's empty, skip the file
		if number == "" {
			warning("No number found in " + file.Name())
			continue
		}

		// Process the file
		stampTxt := fmt.Sprintf(stampTmpl, number)
		stamp(file.Name(), "stamped_" + file.Name(), stampTxt, stampCfg)
		message("Processed " + file.Name())
	}
}

func stamp(inFile string, outFile string, stampTxt string, stampCfg string) {
	onTop := true
	update := false
	wm, _ := api.TextWatermark(stampTxt, stampCfg, onTop, update, types.POINTS)
	pages := []string{"1",}
	api.AddWatermarksFile(inFile, outFile, pages, wm, nil)
}

func warning(message string) {
	yellow := "\033[33m"
	reset := "\033[0m"
	fmt.Printf("%s[WARNING] %s%s\n", yellow, message, reset)
}

func error(message string) {
	red := "\033[31m"
	reset := "\033[0m"
	fmt.Printf("%s[ERROR] %s%s\n", red, message, reset)
	os.Exit(1)
}

func message(message string) {
	blue := "\033[34m"
	reset := "\033[0m"
	fmt.Printf("%s[INFO] %s%s\n", blue, message, reset)
}
