package main

import (
	"fmt"
	"photo-extractor/internal/scanner"
	"strconv"
)

type PhotoEntry struct {
	FileInfo  scanner.FileInfo
	groupMode GroupMode
}

func (p PhotoEntry) GroupKey() string {
	if p.groupMode == ByYear {
		return strconv.Itoa(
			p.FileInfo.CreatedAt.Year(),
		)
	}

	return fmt.Sprintf(
		"%d-%d",
		p.FileInfo.CreatedAt.Year(),
		p.FileInfo.CreatedAt.Month(),
	)
}

func fromFileInfo(files []scanner.FileInfo, mode GroupMode) []PhotoEntry {
	entries := make([]PhotoEntry, len(files))
	for i, file := range files {
		entries[i] = PhotoEntry{
			FileInfo:  file,
			groupMode: mode,
		}
	}
	return entries
}
