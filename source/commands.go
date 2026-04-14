package main

import (
	"os"
	"path/filepath"
	"runtime"
	"sort"
	"strings"
)

func GetCommands() []string {
	seen := make(map[string]struct{})
	var commands []string

	pathEnv := os.Getenv("PATH")

	for _, dir := range filepath.SplitList(pathEnv) {
		info, err := os.Stat(dir)
		if err != nil || !info.IsDir() {
			continue
		}

		entries, err := os.ReadDir(dir)
		if err != nil {
			continue
		}

		sort.Slice(entries, func(i, j int) bool {
			return strings.ToLower(entries[i].Name()) < strings.ToLower(entries[j].Name())
		})

		for _, entry := range entries {
			if entry.IsDir() {
				continue
			}

			name := entry.Name()
			if name == "" {
				continue
			}

			if runtime.GOOS == "windows" {
				ext := strings.ToLower(filepath.Ext(name))
				switch ext {
				case ".exe", ".cmd", ".bat", ".com":
				default:
					continue
				}
				name = strings.TrimSuffix(name, ext)
			} else {
				info, err := entry.Info()
				if err != nil {
					continue
				}
				if info.Mode()&0111 == 0 {
					continue
				}
			}

			lower := strings.ToLower(name)
			if _, exists := seen[lower]; !exists {
				seen[lower] = struct{}{}
				commands = append(commands, lower)
			}
		}
	}

	sort.Slice(commands, func(i, j int) bool {
		return strings.ToLower(commands[i]) < strings.ToLower(commands[j])
	})

	return commands
}
