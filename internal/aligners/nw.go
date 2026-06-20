package aligners

import (
	"github.com/0xmukesh/tinyalign/internal/helpers"
)

type Nw struct {
	Match    int
	Mismatch int
	Gap      int
}

type tracebackDirection int

const (
	dirDiag tracebackDirection = iota
	dirLeft
	dirUp
	dirEnd
)

func (nw *Nw) Align(seqA, seqB string) (string, string) {
	nRows := len(seqA) + 1
	nCols := len(seqB) + 1

	scoringMatrix := helpers.BuildMatrix[int](nRows, nCols)
	tracebackMatrix := helpers.BuildMatrix[tracebackDirection](nRows, nCols)

	for i := range scoringMatrix {
		scoringMatrix[i][0] = i * nw.Gap
	}
	for j := range scoringMatrix[0] {
		scoringMatrix[0][j] = j * nw.Gap
	}

	for i := range tracebackMatrix {
		tracebackMatrix[i][0] = dirUp
	}
	for j := range tracebackMatrix[0] {
		tracebackMatrix[0][j] = dirLeft
	}
	tracebackMatrix[0][0] = dirEnd

	for i := 1; i < len(scoringMatrix); i++ {
		for j := 1; j < len(scoringMatrix[0]); j++ {
			var S int
			if seqA[i-1] == seqB[j-1] {
				S = nw.Match
			} else {
				S = nw.Mismatch
			}

			diag := scoringMatrix[i-1][j-1] + S
			left := scoringMatrix[i][j-1] + nw.Gap
			up := scoringMatrix[i-1][j] + nw.Gap
			score := max(diag, left, up)

			switch score {
			case diag:
				tracebackMatrix[i][j] = dirDiag
			case left:
				tracebackMatrix[i][j] = dirLeft
			case up:
				tracebackMatrix[i][j] = dirUp
			}

			scoringMatrix[i][j] = score
		}
	}

	toBreak := false
	i := len(tracebackMatrix) - 1
	j := len(tracebackMatrix[0]) - 1

	maxLen := len(seqA) + len(seqB)
	bufA := make([]byte, maxLen)
	bufB := make([]byte, maxLen)
	pos := maxLen - 1

	for !toBreak {
		switch tracebackMatrix[i][j] {
		case dirDiag:
			bufA[pos] = seqA[i-1]
			bufB[pos] = seqB[j-1]
			i--
			j--
			pos--
		case dirLeft:
			bufA[pos] = '-'
			bufB[pos] = seqB[j-1]
			j--
			pos--
		case dirUp:
			bufA[pos] = seqA[i-1]
			bufB[pos] = '-'
			i--
			pos--
		case dirEnd:
			toBreak = true
		}
	}

	return string(bufA[pos+1:]), string(bufB[pos+1:])
}
