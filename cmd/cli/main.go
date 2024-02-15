package main

import (
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

	var filenames []string

	for i := 0; i < c.Args().Len(); i++ {
		filenames = append(filenames, c.Args().Get(i))
	}

	if len(filenames) == 0 {
		return errors.New("required at least one file")
	}

	result, _ := wc.WCFilenames(filenames...)

	opts := encoding.CLIOptions{
		WriteLines: writeLines,
		WriteBytes: writeBytes,
		WriteWords: writeWords,
		WriteChars: writeChars,
	}

	if opts.AllFalse() {
		opts.SetAllTrue()
	}

	cliEncoded := encoding.MarshalMultipleNamedStatsCLI(result, opts)

	fmt.Println(cliEncoded)

	return nil
}
