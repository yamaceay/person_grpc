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

func Hash(v string, limit ...int) string {
	h := sha256.New()
	h.Write([]byte(v))
	hash := fmt.Sprintf("%x", h.Sum(nil))
	if len(limit) > 0 {
		if maxchar := limit[0]; maxchar > 0 && maxchar < len(hash) {
			hash = hash[:]
		}
	}
	return hash
}
