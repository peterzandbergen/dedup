package hasher

import (
	"fmt"
	"hash"
	"io"
	"os"
)

type Hashable interface {
	Reader() io.Reader
	SetHash(hashString string)
}

type HashableError map[int]error

func (herr HashableError) Error() string {
	return ""
}

// hashReader creates a hash from hte reader and returns a string or an error.
func hashReader(h hash.Hash, r io.Reader) (string, error) {
	// Hash the reader, but do not close.
	_, err := io.Copy(h, r)
	if err != nil {
		return "", err
	}
	res := h.Sum(nil)
	return fmt.Sprintf("%x", res), nil
}

func HashHashable(h hash.Hash, hr Hashable) error {
	res, err := hashReader(h, hr.Reader())
	if err != nil {
		return err
	}
	hr.SetHash(res)
	return nil
}

func HashHasables(h hash.Hash, hs []Hashable) error {
	var herr HashableError = make(HashableError)
	for i, hh := range hs {
		h.Reset()
		if err := HashHashable(h, hh); err != nil {
			herr[i] = err
		}
	}
	if len(herr) > 0 {
		return herr
	}
	return nil
}

// HashFile returns the hash for the file.
func HashFile(h hash.Hash, filename string) (string, error) {
	// Open the file.
	f, err := os.Open(filename)
	if err != nil {
		return "", err
	}
	// Defer close.
	defer f.Close()

	// Hash the file.
	h.Reset()
	return hashReader(h, f)
}
