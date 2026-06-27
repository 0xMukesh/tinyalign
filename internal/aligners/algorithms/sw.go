package algorithms

import (
	"github.com/0xmukesh/tinyalign/internal/aligners"
	"github.com/0xmukesh/tinyalign/internal/helpers"
)

type Sw struct {
	Substitution aligners.SubstitutionMatrix
	GapPenalty   int
}

func NewSw(substitution aligners.SubstitutionMatrix, gapPenalty int) *Sw {
	return &Sw{
		Substitution: substitution,
		GapPenalty:   gapPenalty,
	}
}

func (sw *Sw) Align(seqA, seqB string) aligners.AlignmentResult {
	nRows := len(seqA) + 1
	nCols := len(seqB) + 1

	scoringMatrix := helpers.BuildMatrix[int](nRows, nCols)
	tracebackMatrix := helpers.BuildMatrix[tracebackDirection](nRows, nCols)

	tracebackMatrix[0][0] = dirEnd
	for i := 1; i < nRows; i++ {
		tracebackMatrix[i][0] = dirUp
	}
	for j := 1; j < nCols; j++ {
		tracebackMatrix[0][j] = dirLeft
	}

	for i := 1; i < nRows; i++ {
		for j := 1; j < nCols; j++ {
			diag := scoringMatrix[i-1][j-1] + sw.Substitution.Score(seqA[i-1], seqB[j-1])
			left := scoringMatrix[i][j-1] + sw.GapPenalty
			up := scoringMatrix[i-1][j] + sw.GapPenalty
			score := max(diag, left, up, 0)

			var dir tracebackDirection
			switch score {
			case diag:
				dir = dirDiag
			case left:
				dir = dirLeft
			case up:
				dir = dirUp
			case 0:
				dir = dirEnd
			}

			scoringMatrix[i][j] = score
			tracebackMatrix[i][j] = dir
		}
	}

	max := 0
	iMax, jMax := 0, 0

	for i := range scoringMatrix {
		for j := range scoringMatrix[i] {
			if scoringMatrix[i][j] > max {
				max = scoringMatrix[i][j]
				iMax = i
				jMax = j
			}
		}
	}

	return dpTraceback(scoringMatrix, tracebackMatrix, seqA, seqB, iMax, jMax, func(i, j int) bool {
		return scoringMatrix[i][j] == 0
	})
}
