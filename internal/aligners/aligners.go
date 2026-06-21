package aligners

import (
	"fmt"
	"strings"
)

type AlignmentResult struct {
	SeqA            string
	SeqB            string
	Score           int
	Identities      int
	Gaps            int
	AlignmentLength int
}

func (ar AlignmentResult) Visualize(chunkSize int) {
	i := 0
	startA := 0
	startB := 0

	fmt.Printf("Identities: %d/%d\n", ar.Identities, ar.AlignmentLength)
	fmt.Printf("Gaps: %d/%d\n", ar.Gaps, ar.AlignmentLength)
	fmt.Printf("Score: %d\n\n", ar.Score)

	for i < len(ar.SeqA) {
		var sb strings.Builder

		chunkStart := i
		chunkEnd := min(i+chunkSize, len(ar.SeqA))
		chunkA := ar.SeqA[chunkStart:chunkEnd]
		chunkB := ar.SeqB[chunkStart:chunkEnd]

		matchLine, countA, countB := ar.buildMatchLine(chunkA, chunkB)
		endA := startA + countA - 1
		endB := startB + countB - 1

		fmt.Fprintf(&sb, "Query %8d %s %d\n", startA, chunkA, endA)
		fmt.Fprintf(&sb, "%s\n", matchLine)
		fmt.Fprintf(&sb, "Sbjct %8d %s %d", startB, chunkB, endB)
		if i+chunkSize < len(ar.SeqA) {
			sb.WriteString("\n\n")
		}

		fmt.Println(sb.String())

		i += chunkSize
		startA = endA
		startB = endB
	}
}

func (ar AlignmentResult) buildMatchLine(chunkA, chunkB string) (string, int, int) {
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
