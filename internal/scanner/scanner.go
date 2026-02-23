package scanner

import (
	"io/fs"
	"path/filepath"
	"strings"
	"time"
)

type FileInfo struct {
	Path      string
	CreatedAt time.Time
	Extension string
}

func ScanFiles(inputDir string, maxDepth int, allowedExts []string) ([]FileInfo, error) {
	var results []FileInfo
	extMap := createExtensionMap(allowedExts)
	baseDepth := getPathDepth(inputDir)

	err := filepath.WalkDir(inputDir, func(path string, d fs.DirEntry, err error) error {
		if err != nil {
			return err
		}

		if d.IsDir() {
			return shouldBeSkipped(path, baseDepth, maxDepth)
		}

		ext := strings.ToLower(filepath.Ext(path))
		if !extMap[ext] {
			return nil
		}

		info, err := d.Info()
		if err != nil {
			return nil
		}

		results = append(results, FileInfo{
			Path:      path,
			CreatedAt: info.ModTime(),
			Extension: ext,
		})
		return nil
	})

	return results, err
}

func createExtensionMap(extensions []string) map[string]bool {
	extensionMap := make(map[string]bool)
	for _, ext := range extensions {
		extensionMap[strings.ToLower(ext)] = true
	}
	return extensionMap
}

func getPathDepth(path string) int {
	return len(
		strings.Split(path, string(filepath.Separator)),
	)
}

func shouldBeSkipped(path string, baseDepth int, maxDepth int) error {
	if getPathDepth(path)-baseDepth > maxDepth {
		return filepath.SkipDir
	}
	return nil
}
