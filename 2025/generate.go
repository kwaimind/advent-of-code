//go:build ignore

package main

import (
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"strconv"
)

func main() {
	// Find the next day number
	files, err := filepath.Glob("day*.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error scanning files: %v\n", err)
		os.Exit(1)
	}

	maxDay := 0
	dayPattern := regexp.MustCompile(`^day(\d+)\.go$`)

	for _, file := range files {
		matches := dayPattern.FindStringSubmatch(filepath.Base(file))
		if len(matches) == 2 {
			day, _ := strconv.Atoi(matches[1])
			if day > maxDay {
				maxDay = day
			}
		}
	}

	nextDay := maxDay + 1
	dayFileName := fmt.Sprintf("day%d.go", nextDay)
	dayFuncName := fmt.Sprintf("day%d", nextDay)

	// Check if file already exists
	if _, err := os.Stat(dayFileName); err == nil {
		fmt.Fprintf(os.Stderr, "Error: %s already exists\n", dayFileName)
		os.Exit(1)
	}

	// Create new day file
	dayContent := fmt.Sprintf(`package main

func %s(filepath string) string {
	return "Not implemented yet"
}
`, dayFuncName)

	if err := os.WriteFile(dayFileName, []byte(dayContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error creating %s: %v\n", dayFileName, err)
		os.Exit(1)
	}

	fmt.Printf("Created %s\n", dayFileName)

	// Update main.go
	mainContent, err := os.ReadFile("main.go")
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error reading main.go: %v\n", err)
		os.Exit(1)
	}

	content := string(mainContent)

	// Find the solutions map and add new entry before closing brace
	mapPattern := regexp.MustCompile(`(var solutions = map\[string\]Solution\{[^}]*)(}\s*)`)
	newEntry := fmt.Sprintf("\t\"day%d\": day%d,\n", nextDay, nextDay)

	updatedContent := mapPattern.ReplaceAllString(content, "${1}"+newEntry+"$2")

	if err := os.WriteFile("main.go", []byte(updatedContent), 0644); err != nil {
		fmt.Fprintf(os.Stderr, "Error updating main.go: %v\n", err)
		os.Exit(1)
	}

	fmt.Printf("Updated main.go with %s\n", dayFuncName)
}
