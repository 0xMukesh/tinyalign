package algorithms

import (
	"github.com/0xmukesh/tinyalign/internal/aligners"
	"github.com/0xmukesh/tinyalign/internal/helpers"
)

type Nw struct {
	Substitution aligners.SubstitutionMatrix
	GapPenalty   int
}

func NewNw(substitution aligners.SubstitutionMatrix, gapPenalty int) *Nw {
	return &Nw{
		Substitution: substitution,
		GapPenalty:   gapPenalty,
	}
}

func (nw *Nw) Align(seqA, seqB string) aligners.AlignmentResult {
	nRows := len(seqA) + 1
	nCols := len(seqB) + 1

	scoringMatrix := helpers.BuildMatrix[int](nRows, nCols)
	tracebackMatrix := helpers.BuildMatrix[tracebackDirection](nRows, nCols)

	for i := range scoringMatrix {
		scoringMatrix[i][0] = nw.GapPenalty * i
	}
	for j := range scoringMatrix[0] {
		scoringMatrix[0][j] = nw.GapPenalty * j
	}

	tracebackMatrix[0][0] = dirEnd
	for i := 1; i < nRows; i++ {
		tracebackMatrix[i][0] = dirUp
	}
	for j := 1; j < nCols; j++ {
		tracebackMatrix[0][j] = dirLeft
	}

	for i := 1; i < nRows; i++ {
		for j := 1; j < nCols; j++ {
			diag := scoringMatrix[i-1][j-1] + nw.Substitution.Score(seqA[i-1], seqB[j-1])
			left := scoringMatrix[i][j-1] + nw.GapPenalty
			up := scoringMatrix[i-1][j] + nw.GapPenalty
			score := max(diag, left, up)

			var dir tracebackDirection
			switch score {
			case diag:
				dir = dirDiag
			case left:
				dir = dirLeft
			case up:
				dir = dirUp
			}

			scoringMatrix[i][j] = score
			tracebackMatrix[i][j] = dir
		}
	}

	return dpTraceback(scoringMatrix, tracebackMatrix, seqA, seqB, nRows-1, nCols-1, func(i, j int) bool {
		return i == 0 && j == 0
	})
}
