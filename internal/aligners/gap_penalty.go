package aligners

type GapPenalty interface {
	Penalty(length int) int
}

type LinearGap struct {
	Gap int
}

func NewLinearGap(gap int) *LinearGap {
	return &LinearGap{
		Gap: gap,
	}
}

func (lg *LinearGap) Penalty(length int) int {
	return lg.Gap * length
}
