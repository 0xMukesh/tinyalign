package aligners

import (
	"fmt"
	"io"
	"strings"
)

type AlignmentResult struct {
	SeqA            string
	SeqB            string
	StartA          int
	StartB          int
	Score           int
	Identities      int
	Gaps            int
	AlignmentLength int
}

func (ar AlignmentResult) Visualize(w io.Writer, vizWidth int) {
	i := 0
	posA := ar.StartA
	posB := ar.StartB

	var header []byte
	header = fmt.Appendf(header, "Identities: %d/%d\n", ar.Identities, ar.AlignmentLength)
	header = fmt.Appendf(header, "Gaps: %d/%d\n", ar.Gaps, ar.AlignmentLength)
	header = fmt.Appendf(header, "Score: %d\n\n", ar.Score)
	w.Write(header)

	for i < len(ar.SeqA) {
		var sb strings.Builder

		chunkStart := i
		chunkEnd := min(i+vizWidth, len(ar.SeqA))
		chunkA := ar.SeqA[chunkStart:chunkEnd]
		chunkB := ar.SeqB[chunkStart:chunkEnd]

		matchLine, countA, countB := ar.buildMatchLine(chunkA, chunkB)
		endA := posA + countA - 1
		endB := posB + countB - 1

		fmt.Fprintf(&sb, "Query %8d %s %d\n", posA, chunkA, endA)
		fmt.Fprintf(&sb, "%s\n", matchLine)
		fmt.Fprintf(&sb, "Sbjct %8d %s %d\n", posB, chunkB, endB)
		if i+vizWidth < len(ar.SeqA) {
			sb.WriteString("\n\n")
		}

		w.Write([]byte(sb.String()))

		i += vizWidth
		posA = endA
		posB = endB
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
		} else if chunkA[j] != '-' && chunkB[j] != '-' {
			fmt.Fprintf(&sb, ".")
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
