package aligners

type Sw struct {
	Substitution SubstitutionMatrix
	Gap          GapPenalty
}

func NewSw(substitution SubstitutionMatrix, gap GapPenalty) *Sw {
	return &Sw{
		Substitution: substitution,
		Gap:          gap,
	}
}

func (sw *Sw) Align(seqA, seqB string) AlignmentResult {
	return dpAlign(dpConfig{
		initMatrix:     sw.initMatrix,
		recurrence:     sw.recurrence,
		tracebackStart: sw.tracebackStart,
		tracebackDone:  sw.tracebackDone,
		substitution:   sw.Substitution,
		gap:            sw.Gap,
	}, seqA, seqB)
}

func (sw *Sw) initMatrix(_, _ int) int {
	return 0
}

func (sw *Sw) recurrence(diag, left, up int) (int, tracebackDirection) {
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

	return score, dir
}

func (sw *Sw) tracebackStart(scoring [][]int) (i, j int) {
	max := 0
	iMax, jMax := 0, 0

	for i := range scoring {
		for j := range scoring[i] {
			if scoring[i][j] > max {
				max = scoring[i][j]
				iMax, jMax = i, j
			}
		}
	}

	return iMax, jMax
}

func (sw *Sw) tracebackDone(_, _, score int) bool {
	return score == 0
}
