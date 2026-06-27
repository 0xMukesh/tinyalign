package config

import (
	"flag"
	"fmt"
	"maps"
	"os"
	"slices"
	"strconv"
	"strings"

	"github.com/0xmukesh/tinyalign/internal/aligners"
)

type TinyAlignAlgorithm int

const (
	AlgorithmNw TinyAlignAlgorithm = iota
	AlgorithmSw
)

type Config struct {
	SeqAPath           string
	SeqBPath           string
	Algorithm          TinyAlignAlgorithm
	Match              *int
	Mismatch           *int
	SubstitutionMatrix string
	GapPenalty         int
	AlignVizWidth      int
}

var (
	seqAPath, seqBPath    string
	rawAlgorithm          string
	rawMatch, rawMismatch string
	substitutionMatrix    string
	gapPenalty            int
	alignVizWidth         int

	algorithmNameMapping = map[string]TinyAlignAlgorithm{
		"nw": AlgorithmNw,
		"sw": AlgorithmSw,
	}
)

func Parse() (*Config, error) {
	validSubstitutionMatrixOptionsStr := strings.Join(slices.Collect(maps.Keys(aligners.SubstitutionMatrixNameMapping)), ", ")
	validAlgorithmOptionsStr := strings.Join(slices.Collect(maps.Keys(algorithmNameMapping)), ", ")

	flag.StringVar(&seqAPath, "seq-a", "", "path to sequence A's fasta file")
	flag.StringVar(&seqBPath, "seq-b", "", "path to sequence B's fasta file")
	flag.StringVar(&rawAlgorithm, "algorithm", "", fmt.Sprintf("algorithm which is to be used pairwise alignment. options: %s", validAlgorithmOptionsStr))
	flag.StringVar(&rawMatch, "match", "", "value which is to be used on match, if substitution matrix is not passed")
	flag.StringVar(&rawMismatch, "mismatch", "", "value which is to be used on mismatch, if substitution matrix is not passed")
	flag.StringVar(&substitutionMatrix, "matrix", "", fmt.Sprintf("name of the substitution matrix which is to be used. options: %s", validSubstitutionMatrixOptionsStr))
	flag.IntVar(&gapPenalty, "gap-penalty", -1, "value which is to be used for gap penalty")
	flag.IntVar(&alignVizWidth, "align-viz-width", 60, "number of characters per row in alignment visualization")
	flag.Parse()

	if seqAPath == "" {
		return nil, fmt.Errorf("missing sequence A's fasta file path")
	}
	if seqBPath == "" {
		return nil, fmt.Errorf("missing sequence B's fasta file path")
	}

	for _, path := range []string{seqAPath, seqBPath} {
		if _, err := os.Stat(path); err != nil {
			if os.IsNotExist(err) {
				return nil, fmt.Errorf("%s path doesn't exist", path)
			}

			return nil, fmt.Errorf("failed to open %s: %s", path, err)
		}
	}

	algorithm, ok := algorithmNameMapping[rawAlgorithm]
	if !ok {
		return nil, fmt.Errorf("invalid algorithm: %s. valid options: %s", rawAlgorithm, validAlgorithmOptionsStr)
	}

	var match, mismatch *int = nil, nil
	if substitutionMatrix != "" {
		if _, ok := aligners.SubstitutionMatrixNameMapping[substitutionMatrix]; !ok {
			return nil, fmt.Errorf("invalid substitution matrix: %s. valid options: %s", substitutionMatrix, validSubstitutionMatrixOptionsStr)
		}
	} else {
		if rawMatch != "" {
			parsed, err := strconv.Atoi(rawMatch)
			if err == nil {
				match = &parsed
			} else {
				return nil, fmt.Errorf("match is expected to be of type int. received: %s", rawMatch)
			}
		}

		if rawMismatch != "" {
			parsed, err := strconv.Atoi(rawMismatch)
			if err == nil {
				mismatch = &parsed
			} else {
				return nil, fmt.Errorf("mismatch is expected to be of type int. received: %s", rawMismatch)
			}
		}

		if match == nil || mismatch == nil {
			if match == nil {
				rawMatch = "<nil>"
			}

			if mismatch == nil {
				rawMismatch = "<nil>"
			}

			return nil, fmt.Errorf("both match and mismatch are required to be passed. received match: %s, mismatch: %s", rawMatch, rawMismatch)
		}
	}

	return &Config{
		SeqAPath:           seqAPath,
		SeqBPath:           seqBPath,
		Algorithm:          algorithm,
		Match:              match,
		Mismatch:           mismatch,
		SubstitutionMatrix: substitutionMatrix,
		GapPenalty:         gapPenalty,
		AlignVizWidth:      alignVizWidth,
	}, nil
}
