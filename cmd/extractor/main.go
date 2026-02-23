package main

import (
	"fmt"
	"os"
	"path/filepath"

	"photo-extractor/internal/fs"
	"photo-extractor/internal/organizer"
	"photo-extractor/internal/scanner"

	"github.com/schollz/progressbar/v3"
)

func main() {
	cfg := loadConfig()

	if cfg.ClearOutput {
		fmt.Println("Cleaning destination directory...")
		os.RemoveAll(cfg.OutputDir)
		os.MkdirAll(cfg.OutputDir, 0755)
	}

	foundFiles, err := scanner.ScanFiles(cfg.InputDir, 1, cfg.Extensions)
	if err != nil {
		fmt.Println("Something went wrong while scaning input directory")
		os.Exit(1)
	}

	progress := progressbar.Default(int64(len(foundFiles)))
	buckets := organizer.Organize(
		fromFileInfo(foundFiles, cfg.GroupingMode),
	)
	for _, bucket := range buckets {
		targetDir := filepath.Join(cfg.OutputDir, bucket.Key)
		if err := os.MkdirAll(targetDir, 0755); err != nil {
			fmt.Printf(
				"Something went wrong while creating directory %s",
				targetDir,
			)
			os.Exit(1)
		}

		for _, item := range bucket.Items {
			fileName := item.FileInfo.CreatedAt.Format("2006_01_02_15_04_05")
			ext := filepath.Ext(item.FileInfo.Path)
			fullFileName := fileName + ext
			destPath := filepath.Join(targetDir, fullFileName)
			uniquePath := fs.GetUniquePath(destPath)

			if err := fs.CopyFile(item.FileInfo.Path, uniquePath); err != nil {
				fmt.Printf("Error copying %s: %v", filepath.Base(item.FileInfo.Path), err)
				os.Exit(1)
			}
			progress.Add(1)
		}
	}
}
