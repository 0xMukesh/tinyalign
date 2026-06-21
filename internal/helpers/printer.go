package helpers

import (
	"fmt"
	"strings"
)

func VisualizeAlignedSeqs(alignedSeqA, alignedSeqB string, chunkSize int) {
	i := 0
	startA := 0
	startB := 0

	for i < len(alignedSeqA) {
		var sb strings.Builder

		chunkStart := i
		chunkEnd := min(i+chunkSize, len(alignedSeqA))
		chunkA := alignedSeqA[chunkStart:chunkEnd]
		chunkB := alignedSeqB[chunkStart:chunkEnd]

		matchLine, countA, countB := buildMatchLine(chunkA, chunkB)
		endA := startA + countA - 1
		endB := startB + countB - 1

		fmt.Fprintf(&sb, "Query %8d %s %d\n", startA, chunkA, endA)
		fmt.Fprintf(&sb, "%s\n", matchLine)
		fmt.Fprintf(&sb, "Sbjct %8d %s %d\n\n", startB, chunkB, endB)
		fmt.Println(sb.String())

		i += chunkSize
		startA = endA
		startB = endB
	}
}

func buildMatchLine(chunkA, chunkB string) (string, int, int) {
	var sb strings.Builder
	countA := 0
	countB := 0

	fmt.Fprintf(&sb, "%*s", 15, "")
	for j := 0; j < len(chunkA); j++ {
		if chunkA[j] == chunkB[j] {
			fmt.Fprintf(&sb, "|")
		} else {
			fmt.Fprintf(&sb, " ")
		}

		if chunkA[j] != '-' {
			countA++
		}
		if chunkB[j] != '-' {
			countB++
		}
	}

	return sb.String(), countA, countB
}
