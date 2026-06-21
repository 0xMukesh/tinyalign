package aligners

type Nw struct {
	Substitution SubstitutionMatrix
	Gap          GapPenalty
}

func NewNw(substitution SubstitutionMatrix, gap GapPenalty) *Nw {
	return &Nw{
		Substitution: substitution,
		Gap:          gap,
	}
}

func (nw *Nw) Align(seqA, seqB string) AlignmentResult {
	return dpAlign(dpConfig{
		initMatrix:     nw.initMatrix,
		recurrence:     nw.recurrence,
		tracebackStart: nw.tracebackStart,
		tracebackDone:  nw.tracebackDone,
		substitution:   nw.Substitution,
		gap:            nw.Gap,
	}, seqA, seqB)
}

func (nw *Nw) initMatrix(i, j int) int {
	if i == 0 && j != 0 {
		return nw.Gap.Penalty(j)
	} else if i != 0 && j == 0 {
		return nw.Gap.Penalty(i)
	} else {
		return 0
	}
}

func (nw *Nw) recurrence(diag, left, up int) (int, tracebackDirection) {
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

	return score, dir
}

func (nw *Nw) tracebackStart(scoring [][]int) (i, j int) {
	return len(scoring) - 1, len(scoring[0]) - 1
}

func (nw *Nw) tracebackDone(i, j, _ int) bool {
	return i == 0 && j == 0
}
