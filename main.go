package main

import (
	"log"

	"github.com/0xmukesh/tinyalign/internal/aligners"
	"github.com/0xmukesh/tinyalign/internal/aligners/algorithms"
	"github.com/0xmukesh/tinyalign/internal/config"
	"github.com/0xmukesh/tinyalign/internal/helpers"
)

func main() {
	cfg, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse args: %s", err)
	}

	seqA, err := helpers.ParseFastaFile(cfg.SeqAPath)
	if err != nil {
		log.Fatalf("failed to parse fasta file at %s: %s", cfg.SeqAPath, err)
	}

	seqB, err := helpers.ParseFastaFile(cfg.SeqBPath)
	if err != nil {
		log.Fatalf("failed to parse fasta file at %s: %s", cfg.SeqBPath, err)
	}

	var substitution aligners.SubstitutionMatrix
	if cfg.Match != nil && cfg.Mismatch != nil {
		substitution = aligners.NewNaiveSubstitution(*cfg.Match, *cfg.Mismatch)
	} else if cfg.SubstitutionMatrix != "" {
		substitution, err = aligners.NewFromMatrix(cfg.SubstitutionMatrix)
		if err != nil {
			log.Fatalf("failed to construct substitution matrix: %s", err)
		}
	}

	var result aligners.AlignmentResult
	switch cfg.Algorithm {
	case config.AlgorithmNw:
		nw := algorithms.NewNw(substitution, cfg.GapPenalty)
		result = nw.Align(seqA, seqB)
	case config.AlgorithmSw:
		sw := algorithms.NewSw(substitution, cfg.GapPenalty)
		result = sw.Align(seqA, seqB)
	}

	result.Visualize(cfg.AlignVizWidth)
}
