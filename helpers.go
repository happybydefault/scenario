package scenario

import (
	"os"
	"testing"
)

func WriteFile(t *testing.T, filename string, data []byte) {
	t.Helper()

	if err := os.WriteFile(filename, data, 0666); err != nil {
		t.Fatalf("could not write to file %q: %v", filename, err)
	}
}

func ReadFile(t *testing.T, filename string) []byte {
	t.Helper()

	data, err := os.ReadFile(filename)
	if err != nil {
		t.Fatalf("could not read file: %v", err)
	}

	return data
}
