package config

import (
	"flag"
	"fmt"
	"maps"
	"slices"
	"strings"
)

type TinyAlignAlgorithm int

const (
	AlgorithmNw TinyAlignAlgorithm = iota
	AlgorithmSw
)

type arrayFlag []string

func (af *arrayFlag) Set(v string) error {
	*af = append(*af, v)
	return nil
}

func (af *arrayFlag) String() string {
	return strings.Join(*af, ", ")
}

type Config struct {
	FastaFiles []string
	Algorithm  TinyAlignAlgorithm
}

var (
	rawFastaFiles arrayFlag
	rawAlgorithm  string

	algorithmNameMapping = map[string]TinyAlignAlgorithm{
		"nw": AlgorithmNw,
		"sw": AlgorithmSw,
	}
)

func Parse() (Config, error) {
	flag.Var(&rawFastaFiles, "fasta", "path of the fasta files which need to be aligned")
	flag.StringVar(&rawAlgorithm, "algorithm", "", "algorithm which is to be used for sequence alignment")
	flag.Parse()

	if len(rawFastaFiles) != 2 {
		return Config{}, fmt.Errorf("expected 2 fasta files, got %d\n", len(rawFastaFiles))
	}

	validAlgorithmOptions := strings.Join(slices.Collect(maps.Keys(algorithmNameMapping)), ", ")
	algorithm, ok := algorithmNameMapping[rawAlgorithm]
	if !ok {
		return Config{}, fmt.Errorf("unknown algorithm %s. valid options: %s\n", validAlgorithmOptions)
	}

	return Config{
		FastaFiles: []string(rawFastaFiles),
		Algorithm:  algorithm,
	}, nil
}
