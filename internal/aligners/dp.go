package aligners

import "github.com/0xmukesh/tinyalign/internal/helpers"

type tracebackDirection int

const (
	dirDiag tracebackDirection = iota
	dirLeft
	dirUp
	dirEnd
)

type dpConfig struct {
	initMatrix     func(i, j int) int
	recurrence     func(diag, left, up int) (score int, dir tracebackDirection)
	tracebackStart func(scoring [][]int) (i, j int)
	tracebackDone  func(i, j, score int) bool

	substitution SubstitutionMatrix
	gap          GapPenalty
}

func dpAlign(cfg dpConfig, seqA, seqB string) AlignmentResult {
	nRows := len(seqA) + 1
	nCols := len(seqB) + 1

	scoringMatrix := helpers.BuildMatrix[int](nRows, nCols)
	tracebackMatrix := helpers.BuildMatrix[tracebackDirection](nRows, nCols)

	for i := range scoringMatrix {
		for j := range scoringMatrix[i] {
			scoringMatrix[i][j] = cfg.initMatrix(i, j)
		}
	}

	for i := range tracebackMatrix {
		tracebackMatrix[i][0] = dirUp
	}
	for j := range tracebackMatrix[0] {
		tracebackMatrix[0][j] = dirLeft
	}
	tracebackMatrix[0][0] = dirEnd

	for i := 1; i < nRows; i++ {
		for j := 1; j < nCols; j++ {
			diag := scoringMatrix[i-1][j-1] + cfg.substitution.Score(seqA[i-1], seqB[j-1])
			left := scoringMatrix[i][j-1] + cfg.gap.Penalty(1)
			up := scoringMatrix[i-1][j] + cfg.gap.Penalty(1)
			score, dir := cfg.recurrence(diag, left, up)

			scoringMatrix[i][j] = score
			tracebackMatrix[i][j] = dir
		}
	}

	i, j := cfg.tracebackStart(scoringMatrix)

	maxLen := len(seqA) + len(seqB)
	bufA := make([]byte, maxLen)
	bufB := make([]byte, maxLen)
	pos := maxLen - 1
	score := 0
	identities := 0
	gaps := 0

	for {
		if cfg.tracebackDone(i, j, scoringMatrix[i][j]) {
			break
		}

		switch tracebackMatrix[i][j] {
		case dirDiag:
			bufA[pos] = seqA[i-1]
			bufB[pos] = seqB[j-1]
			score += cfg.substitution.Score(seqA[i-1], seqB[j-1])
			if seqA[i-1] == seqB[j-1] {
				identities++
			}

			i--
			j--
			pos--
		case dirLeft:
			bufA[pos] = '-'
			bufB[pos] = seqB[j-1]
			score += cfg.gap.Penalty(1)
			gaps++

			j--
			pos--
		case dirUp:
			bufA[pos] = seqA[i-1]
			bufB[pos] = '-'
			score += cfg.gap.Penalty(1)
			gaps++

			i--
			pos--
		}
	}

	return AlignmentResult{
		SeqA:            string(bufA[pos+1:]),
		SeqB:            string(bufB[pos+1:]),
		Score:           score,
		Identities:      identities,
		Gaps:            gaps,
		AlignmentLength: len(string(bufA[pos+1:])),
	}
}
