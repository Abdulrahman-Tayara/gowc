package main

import "github.com/urfave/cli/v2"

func CLIApp(h cli.ActionFunc) *cli.App {
	app := &cli.App{
		Name: "Golang wc Tool",
		Authors: []*cli.Author{
			{
				Name:  "Abdulrahman Tayara",
				Email: "abdulrahman.tayara@gmail.com",
			},
		},
		Usage: "wc – word, line, character, and byte count",
		Description: `The wc utility displays the number of lines, words, and bytes contained in each input file, or standard input (if no file is specified) to the
		standard output.  A line is defined as a string of characters delimited by a ⟨newline⟩ character.  Characters beyond the final ⟨newline⟩ character
		will not be included in the line count.
   
		A word is defined as a string of characters delimited by white space characters.  White space characters are the set of characters for which the
		iswspace(3) function returns true.  If more than one input file is specified, a line of cumulative counts for all the files is displayed on a separate
		line after the output for the last file.`,
		Flags: []cli.Flag{
			&cli.BoolFlag{
				Name:    "lines",
				Aliases: []string{"l"},
				Usage:   "Count the inputs's lines",
			},
			&cli.BoolFlag{
				Name:    "words",
				Aliases: []string{"w"},
				Usage:   "Count the input's words",
			},
			&cli.BoolFlag{
				Name:    "chars",
				Aliases: []string{"m"},
				Usage:   "Count the input's chars",
			},
			&cli.BoolFlag{
				Name:    "bytes",
				Aliases: []string{"c"},
				Usage:   "Count the inputs's bytes",
			},
		},
		Args:   true,
		Action: h,
	}

	return app
}
