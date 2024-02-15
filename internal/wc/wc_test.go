package wc

import (
	"io"
	"os"
	"path/filepath"
	"runtime"
	"strings"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestWC(t *testing.T) {
	tests := []struct {
		name     string
		input    io.Reader
		expected Stats
	}{
		{
			input: strings.NewReader("Hello \n I'm Abdulrahman"),
			expected: Stats{
				Bytes: 23,
				Chars: 23,
				Lines: 2,
				Words: 3,
			},
		},
		{
			input: strings.NewReader("       Hello \n I'm Abdulrahman"),
			expected: Stats{
				Bytes: 30,
				Chars: 30,
				Lines: 2,
				Words: 3,
			},
		},
		{
			input: strings.NewReader("僤A"), // utf-8 encoding
			expected: Stats{
				Bytes: 4,
				Chars: 2,
				Lines: 1,
				Words: 1,
			},
		},
		{
			input: func() io.Reader {
				_, f, _, _ := runtime.Caller(0)
				ff, err := os.Open(
					filepath.Join(filepath.Dir(f), "..", "..", "tests", "data", "7850_chars_1192_words_26_lines.txt"),
				)

				if err != nil {
					t.Error(err)
					return nil
				}

				return ff
			}(),
			expected: Stats{
				Bytes: 7850,
				Chars: 7850,
				Lines: 26,
				Words: 1192,
			},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			actual, err := WC(test.input)

			assert.NoError(t, err)

			if !assert.ObjectsAreEqual(test.expected, *actual) {
				t.Errorf("expected %v, actual %v", test.expected, *actual)
			}
		})
	}

}

func TestWCFileNames(t *testing.T) {
	_, f, _, _ := runtime.Caller(0)
	filename := filepath.Join(filepath.Dir(f), "..", "..", "tests", "data", "7850_chars_1192_words_26_lines.txt")

	stats, err := WCFilenames(filename)

	assert.NoError(t, err)
	assert.NotNil(t, stats)
}
