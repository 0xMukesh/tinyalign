package main

import (
	"log"

	"github.com/0xmukesh/tinyalign/internal/aligners"
	"github.com/0xmukesh/tinyalign/internal/config"
	"github.com/0xmukesh/tinyalign/internal/helpers"
)

func main() {
	config, err := config.Parse()
	if err != nil {
		log.Fatalf("failed to parse args: %s", err)
	}

	seqA, err := helpers.ParseFastaFile(config.FastaFiles[0])
	if err != nil {
		log.Fatalf("failed to parse 1st fasta file: %s", err)
	}

	seqB, err := helpers.ParseFastaFile(config.FastaFiles[1])
	if err != nil {
		log.Fatalf("failed to parse 2nd fasta file: %s", err)
	}

	nw := aligners.Nw{
		Match:    1,
		Mismatch: -1,
		Gap:      -2,
	}

	nw.Align(seqA, seqB)
}
