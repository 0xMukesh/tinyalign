package main

import (
	"log"

	"github.com/0xmukesh/tinyalign/internal/aligners"
	"github.com/0xmukesh/tinyalign/internal/config"
	"github.com/0xmukesh/tinyalign/internal/helpers"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse args: %s", err)
	}

	seqA, err := helpers.ParseFastaFile(cfg.FastaFiles[0])
	if err != nil {
		log.Fatalf("failed to parse fasta file: %s", err)
	}

	seqB, err := helpers.ParseFastaFile(cfg.FastaFiles[1])
	if err != nil {
		log.Fatalf("failed to parse fasta file: %s", err)
	}

	var result aligners.AlignmentResult
	switch cfg.Algorithm {
	case config.AlgorithmNw:
		nw := aligners.NewNw(&aligners.Blosum62{}, aligners.NewLinearGap(-2))
		result = nw.Align(seqA, seqB)
	case config.AlgorithmSw:
		sw := aligners.NewSw(&aligners.Blosum62{}, aligners.NewLinearGap(-2))
		result = sw.Align(seqA, seqB)
	default:
		log.Fatalf("invalid algorithm: %d\n", cfg.Algorithm)
	}

	result.Visualize(cfg.ChunkSize)
}
