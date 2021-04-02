package entropy

import (
	"bytes"
	"math"
	"strings"

	"github.com/nobe4/safe/internal/logger"
)

// References
var (
	references = map[string]struct {
		dict      []byte
		threshold float64
	}{
		"hex": {
			dict:      []byte("1234567890abcdefABCDEF"),
			threshold: 3.0,
		},
		"base64": {
			dict:      []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/="),
			threshold: 3.5,
		}, "ascii": {
			// Printable ASCII characters
			dict:      []byte(`ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789+/=!"#$%&'()*,-.:;<>?@[]^_\{|}~` + "`"),
			threshold: 3.5,
		},
	}
)

// List builds the list of all the names.
func List() string {
	r := []string{}
	for n := range references {
		r = append(r, n)
	}
	return strings.Join(r, " ")
}

func Select(name *string, threshold *float64) ([]byte, float64) {
	logger.Debug("Select dictionnary and threshold.")

	// Default to ascii
	dict := references["ascii"].dict
	thres := references["ascii"].threshold

	// Override the dict
	if name != nil && *name != "" {
		if ref, ok := references[*name]; ok {
			logger.Debug("Using name", *name)

			dict = ref.dict
			thres = ref.threshold
		} else {
			logger.Warn("Name not found:", *name, ", defaulting to ascii.")
		}
	}

	// Overwrite the threshold
	if threshold != nil {
		if *threshold > 0 {
			logger.Debug("Using threshold", *threshold)

			thres = *threshold
		} else {
			logger.Warn("Invalid threshold:", *threshold, ", defaulting to", thres, ".")
		}
	}

	logger.Debug("Using dict", string(dict), "and threshold", thres)
	return dict, thres
}

// Check extract all the parts of the input that are higher than a
// defined threshold for the corresponding dict.
func Check(in []byte, dict []byte, thres float64) [][]byte {
	found := [][]byte{}

	for _, field := range bytes.Fields(in) {
		logger.Debug("Working on new field", string(field))

		if compute(field, dict) > thres {
			found = append(found, field)
		}
	}

	logger.Debug("Found fields:", found)
	return found
}

func compute(in []byte, dict []byte) float64 {
	logger.Debug("Compute entropy for input", string(in))

	inLen := float64(len(in))
	entropy := 0.0

	if inLen > 0.0 {
		for _, b := range dict {
			count := bytes.Count(in, []byte{b})
			percent := float64(count) / inLen

			if percent > 0 {
				entropy += -percent * math.Log2(percent)
			}
		}
	}

	logger.Info("Entropy for", string(in), ":", entropy)
	return entropy
}
