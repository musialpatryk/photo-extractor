package main

import (
	"errors"
	"fmt"
	"os"

	"github.com/charmbracelet/huh"
)

const (
	DefaultInPath  = "input"
	DefaultOutPath = "output"
)

var PhotoExtensions = []string{".jpg", ".jpeg", ".png", ".webp"}

type GroupMode int

const (
	ByMonth GroupMode = iota
	ByYear
)

type Config struct {
	InputDir     string
	OutputDir    string
	Extensions   []string
	GroupingMode GroupMode
	ClearOutput  bool
}

func loadConfig() Config {
	cfg := Config{
		Extensions:  PhotoExtensions,
		ClearOutput: false,
	}

	confirmed := true
	form := huh.NewForm(
		huh.NewGroup(
			huh.NewInput().
				Title("Source Directory").
				Placeholder(DefaultInPath).
				Value(&cfg.InputDir).
				Validate(dirExists),

			huh.NewInput().
				Title("Destination Directory").
				Placeholder(DefaultOutPath).
				Value(&cfg.OutputDir).
				Validate(dirExists),

			huh.NewSelect[GroupMode]().
				Title("How to group them?").
				Options(
					huh.NewOption("By Month (2024/05)", ByMonth),
					huh.NewOption("By Year (2024)", ByYear),
				).
				Value(&cfg.GroupingMode),

			huh.NewConfirm().
				Title("Clear destination directory?").
				Description("This will delete ALL files in the output folder before starting!").
				Affirmative("Yes, clean it").
				Negative("No, keep existing files").
				Value(&cfg.ClearOutput),
		),
		huh.NewGroup(
			huh.NewConfirm().
				Title("Ready to extract?").
				DescriptionFunc(func() string {
					modeStr := "By Month"
					if cfg.GroupingMode == ByYear {
						modeStr = "By Year"
					}
					cleanStr := "No"
					if cfg.ClearOutput {
						cleanStr = "Yes (DANGEROUS)"
					}

					return fmt.Sprintf(
						"Summary:\n• Source: %s\n• Destination: %s\n• Grouping: %s\n• Clear Output: %s",
						cfg.InputDir, cfg.OutputDir, modeStr, cleanStr,
					)
				}, &cfg.GroupingMode). // Przekazujemy tylko jeden wskaźnik
				Affirmative("Yes, let's go!").
				Negative("No, take me out").
				Value(&confirmed),
		),
	)

	err := form.Run()
	if err != nil {
		handleFormError(err)
	}

	if !confirmed {
		fmt.Println("Extraction cancelled!")
		os.Exit(0)
	}

	return cfg
}

func dirExists(path string) error {
	info, err := os.Stat(path)
	if err != nil {
		return fmt.Errorf("directory not found: %s", path)
	}
	if !info.IsDir() {
		return fmt.Errorf("path is not a directory")
	}
	return nil
}

func handleFormError(err error) {
	if errors.Is(err, huh.ErrUserAborted) {
		fmt.Println("\n[!] Aborted by user.")
		os.Exit(0)
	}
	fmt.Fprintf(os.Stderr, "Error: %v\n", err)
	os.Exit(1)
}
