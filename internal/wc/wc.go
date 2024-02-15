package wc

import (
	"bufio"
	"io"
	"os"
	"sync"
	"tayara/go-wc/internal/utils"
)

type runeHandler func(r rune, rSize int, eof bool)

type Stats struct {
	Bytes int
	Lines int
	Chars int
	Words int
}

type filenameStatsOutput struct {
	filename string
	err      error
	stats    *Stats
}

func WCFilenames(filenames ...string) (map[string]Stats, error) {
	var wg sync.WaitGroup

	channel := make(chan filenameStatsOutput, len(filenames))

	for _, filename := range filenames {
		wg.Add(1)

		f := filename

		go wcFilenameWorker(f, &wg, channel)
	}

	go func() {
		wg.Wait()
		close(channel)
	}()

	results := map[string]Stats{}

	for {
		fileStats, ok := <-channel

		if !ok {
			break
		}

		if fileStats.err != nil {
			return nil, fileStats.err
		}

		results[fileStats.filename] = *fileStats.stats
	}

	return results, nil

}

func wcFilenameWorker(filename string, wg *sync.WaitGroup, destChan chan<- filenameStatsOutput) {
	defer wg.Done()

	stats, err := WCFilename(filename)

	destChan <- filenameStatsOutput{
		err:      err,
		stats:    stats,
		filename: filename,
	}
}

func WCFilename(filename string) (*Stats, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}

	return WC(f)
}

func WC(src io.Reader) (*Stats, error) {
	reader := bufio.NewReader(src)

	stats := Stats{}

	handlers := []runeHandler{
		wordsCounter(&stats),
		linesCounter(&stats),
		charsCounter(&stats),
		bytesCounter(&stats),
	}

	emitRune := func(r rune, size int, eof bool) {
		for _, h := range handlers {
			h(r, size, eof)
		}
	}

	for {
		r, size, err := reader.ReadRune()

		if err == io.EOF {
			emitRune(' ', 0, true)
			break
		} else if err != nil {
			return nil, err
		}

		emitRune(r, size, false)
	}

	return &stats, nil
}

func wordsCounter(stats *Stats) runeHandler {
	insideWord := false
	return func(r rune, rSize int, eof bool) {
		if eof && insideWord {
			stats.Words++
			return
		}

		if utils.IsWhitespace(r) {
			if insideWord {
				stats.Words++
			}
			insideWord = false
		} else {
			insideWord = true
		}
	}
}

func linesCounter(stats *Stats) runeHandler {
	return func(r rune, rSize int, eof bool) {
		if eof {
			if stats.Lines == 0 {
				stats.Lines = 1
			} else {
				stats.Lines++
			}
			return
		}
		if r == '\n' {
			stats.Lines++
		}
	}
}

func bytesCounter(stats *Stats) runeHandler {
	return func(r rune, rSize int, eof bool) {
		stats.Bytes += rSize
	}
}

func charsCounter(stats *Stats) runeHandler {
	return func(r rune, rSize int, eof bool) {
		if !eof {
			stats.Chars++
		}
	}
}
