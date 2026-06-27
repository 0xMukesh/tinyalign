package algorithms

import "github.com/0xmukesh/tinyalign/internal/aligners"

type tracebackDirection int

const (
	dirDiag tracebackDirection = iota
	dirLeft
	dirUp
	dirEnd
)

func dpTraceback(
	scoringMatrix [][]int, tracebackMatrix [][]tracebackDirection,
	seqA, seqB string,
	i, j int,
	toStop func(i, j int) bool,
) aligners.AlignmentResult {
	maxLen := len(seqA) + len(seqB)
	bufA := make([]byte, maxLen)
	bufB := make([]byte, maxLen)
	pos := maxLen - 1
	score := scoringMatrix[i][j]
	identities := 0
	gaps := 0

	for {
		if toStop(i, j) {
			break
		}

		switch tracebackMatrix[i][j] {
		case dirDiag:
			a, b := seqA[i-1], seqB[j-1]
			bufA[pos] = a
			bufB[pos] = b

			if a == b {
				identities++
			}

			i--
			j--
			pos--
		case dirLeft:
			b := seqB[j-1]
			bufA[pos] = '-'
			bufB[pos] = b

			gaps++
			j--
			pos--
		case dirUp:
			a := seqA[i-1]
			bufA[pos] = a
			bufB[pos] = '-'

			gaps++
			i--
			pos--
		}
	}

	return aligners.AlignmentResult{
		SeqA:            string(bufA[pos+1:]),
		SeqB:            string(bufB[pos+1:]),
		StartA:          i,
		StartB:          j,
		Score:           score,
		Identities:      identities,
		Gaps:            gaps,
		AlignmentLength: len(bufA[pos+1:]),
	}
}
