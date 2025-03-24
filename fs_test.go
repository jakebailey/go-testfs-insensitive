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

	fs := os.DirFS(filepath.Join(dir, "testdata"))

	if err := fstest.TestFS(fs, "foo/bar/file.txt"); err != nil {
		t.Fatalf("failed to test file system: %v", err)
	}

	if err := fstest.TestFS(fs, "FOO/bar/file.txt"); err != nil {
		t.Fatalf("failed to test file system case insensitively: %v", err)
	}
}
