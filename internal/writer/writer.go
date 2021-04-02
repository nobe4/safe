package writer

import (
	"bytes"
	"io"
	"os"
	"regexp"

	"github.com/nobe4/safe/internal/entropy"
	"github.com/nobe4/safe/internal/logger"
)

// safeWriter impements the Writer struct
// It writes to out after filtering elements via the regexp re.
// filter stores the filtering function.
type safeWriter struct {
	out    *os.File
	re     *regexp.Regexp
	filter func([]byte) []byte
	censor byte
}

// New creates a new safe Writer.
func New(filter *string, censor *string, dict *string, threshold *float64) (io.Writer, error) {
	logger.Debug("Create a new writer with filter:", filter)

	i := &safeWriter{
		out: os.Stdout,
	}

	// Defaults:
	// - entropy filter
	// - X censor
	i.filter = i.entropyFilter(dict, threshold)
	i.censor = byte('X')

	// Set the censor character
	if censor != nil && *censor != "" {
		i.censor = (*censor)[0]
	}

	// If regexp isn't nil and it compiles, use regex filtering.
	if filter != nil && *filter != "" {
		logger.Debug("Test the regexp filter:", *filter)

		re, err := regexp.Compile(*filter)
		if err != nil {
			return nil, err
		}

		logger.Info("Use the regexp filter:", *filter)

		i.re = re
		i.filter = i.regexpFilter
	}

	return i, nil
}

// nBytes create a slice of n 'X'.
func (i *safeWriter) nBytes(n int) []byte {
	logger.Debug("Create a slice of", n, i.censor)

	var o []byte

	for x := 0; x < n; x++ {
		o = append(o, i.censor)
	}

	return o
}

// replaceAll replace all the bytes with X.
func (i *safeWriter) replaceAll(in []byte) []byte {
	logger.Debug("Replace all bytes in input:", in)
	return i.nBytes(len(in))
}

// regexpFilter replace all matches with X.
func (i *safeWriter) regexpFilter(in []byte) []byte {
	logger.Debug("Use regexpFilter with input:", in)

	return i.re.ReplaceAllFunc(in, i.replaceAll)
}

// entropyFilter replace all the high-entropy strings with X.
func (i *safeWriter) entropyFilter(dictName *string, threshold *float64) func([]byte) []byte {
	logger.Debug("Select entropy parameters.")

	dict, thres := entropy.Select(dictName, threshold)

	return func(in []byte) []byte {
		logger.Debug("Use entropyFilter with input:", string(in))

		for _, found := range entropy.Check(in, dict, thres) {
			logger.Debug("Replace", found)

			foundLen := len(found)
			if foundLen == 0 {
				continue
			}
			// Build a replacement of similar length
			replacement := i.nBytes(foundLen)

			// Replace all instances
			in = bytes.ReplaceAll(in, found, replacement)
		}

		return in
	}
}

// Write writes to the defined output the input bytes after filtering.
func (i *safeWriter) Write(in []byte) (int, error) {
	logger.Debug("Use Write with input:", in)

	return i.out.Write(i.filter(in))
}
