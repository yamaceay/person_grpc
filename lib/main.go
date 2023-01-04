package lib

import (
	"crypto/sha256"
	"fmt"
	"io"
	"os"
)

func ReadFile(file string) ([]byte, error) {
	fp, err := os.Open(file)
	if err != nil {
		return nil, fmt.Errorf("unable to open file: %w", err)
	}

	bytes, err := io.ReadAll(fp)
	if err != nil {
		return nil, fmt.Errorf("unable to read file: %w", err)
	}

	return bytes, nil
}

func Hash(v string) string {
	h := sha256.New()
	h.Write([]byte(v))
	return fmt.Sprintf("%x", h.Sum(nil))
}
