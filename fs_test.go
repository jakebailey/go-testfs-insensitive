package main

import (
	"os"
	"path/filepath"
	"testing"
	"testing/fstest"
)

func TestFS(t *testing.T) {
	dir, err := os.Getwd()
	if err != nil {
		t.Fatalf("failed to get current working directory: %v", err)
	}
	dir = filepath.Join(dir, "testdata")

	const testFile = "foo/bar/file.txt"
	const testFileUpper = "FOO/BAR/FILE.TXT"

	caseInsensitive := false
	if _, err := os.Stat(filepath.Join(dir, testFileUpper)); err == nil {
		caseInsensitive = true
	}

	fs := os.DirFS(dir)

	if err := fstest.TestFS(fs, testFile); err != nil {
		t.Fatalf("failed to test file system: %v", err)
	}

	if caseInsensitive {
		if err := fstest.TestFS(fs, testFileUpper); err != nil {
			t.Fatalf("failed to test file system case insensitively: %v", err)
		}
	}
}
