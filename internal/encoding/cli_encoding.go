package encoding

import (
	"fmt"
	"strings"
	"tayara/gowc/internal/wc"
)

type CLIOptions struct {
	WriteLines bool
	WriteChars bool
	WriteWords bool
	WriteBytes bool
}

func (o *CLIOptions) AllFalse() bool {
	return !(o.WriteBytes || o.WriteChars || o.WriteWords || o.WriteLines)
}

func (o *CLIOptions) SetAllTrue() {
	o.WriteBytes = true
	o.WriteChars = true
	o.WriteLines = true
	o.WriteWords = true
}

func MarshalMultipleNamedStatsCLI(
	stats map[string]wc.Stats,
	opts CLIOptions,
) string {
	var builder strings.Builder

	for name, s := range stats {
		builder.WriteString(marshalNamedStats(name, &s, &opts))
		builder.WriteRune('\n')
	}

	return builder.String()
}

func marshalNamedStats(name string, stats *wc.Stats, opts *CLIOptions) string {
	var builder strings.Builder

	if opts.WriteBytes {
		builder.WriteString(fmt.Sprintf("bytes: %v\t", stats.Bytes))
	}
	if opts.WriteChars {
		builder.WriteString(fmt.Sprintf("chars: %v\t", stats.Chars))
	}
	if opts.WriteLines {
		builder.WriteString(fmt.Sprintf("lines: %v\t", stats.Lines))
	}
	if opts.WriteWords {
		builder.WriteString(fmt.Sprintf("words: %v\t", stats.Words))
	}

	builder.WriteString(name)

	return builder.String()
}

func MarshalStatsCLI(stats wc.Stats, opts CLIOptions) string {
	return marshalNamedStats("", &stats, &opts)
}
