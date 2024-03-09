package main

import (
	"bufio"
	"errors"
	"fmt"
	"log"
	"os"
	"tayara/gowc/internal/encoding"
	"tayara/gowc/internal/wc"

	"github.com/urfave/cli/v2"
)

func main() {
	app := CLIApp(handler)

	if err := app.Run(os.Args); err != nil {
		log.Fatal(err)
	}
}

func handler(c *cli.Context) error {
	writeLines := c.Bool("lines")
	writeChars := c.Bool("chars")
	writeBytes := c.Bool("bytes")
	writeWords := c.Bool("words")

	opts := encoding.CLIOptions{
		WriteLines: writeLines,
		WriteBytes: writeBytes,
		WriteWords: writeWords,
		WriteChars: writeChars,
	}

	if opts.AllFalse() {
		opts.SetAllTrue()
	}

	cliEncoded, err := processInput(c, opts)
	if err != nil {
		return err
	}

	fmt.Println(cliEncoded)

	return nil
}

func processInput(c *cli.Context, opts encoding.CLIOptions) (string, error) {
	var cliEncoded string

	if canReadFromStdin() { // Stdin source
		stats, err := handleStdinInput(opts)
		if err != nil {
			return "", err
		}
		cliEncoded = encoding.MarshalStatsCLI(*stats, opts)
	} else {
		filesStats, err := handleFilesInput(c, opts)
		if err != nil {
			return "", err
		}
		cliEncoded = encoding.MarshalMultipleNamedStatsCLI(filesStats, opts)
	}
	return cliEncoded, nil
}

func canReadFromStdin() bool {
	stat, _ := os.Stdin.Stat()
	return (stat.Mode() & os.ModeCharDevice) == 0
}

func handleStdinInput(opts encoding.CLIOptions) (*wc.Stats, error) {
	var stdin []byte

	{ // Read from stdin
		scanner := bufio.NewScanner(os.Stdin)
		for scanner.Scan() {
			stdin = append(stdin, scanner.Bytes()...)
		}
		if err := scanner.Err(); err != nil {
			return nil, err
		}
	}

	return wc.WCContent(stdin)
}

func handleFilesInput(c *cli.Context, opts encoding.CLIOptions) (map[string]wc.Stats, error) {
	var filenames []string

	for i := 0; i < c.Args().Len(); i++ {
		filenames = append(filenames, c.Args().Get(i))
	}

	if len(filenames) == 0 {
		return nil, errors.New("required at least one file")
	}

	result, _ := wc.WCFilenames(filenames...)
	return result, nil
}
