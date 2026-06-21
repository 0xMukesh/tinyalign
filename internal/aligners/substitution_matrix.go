package aligners

type SubstitutionMatrix interface {
	Score(a, b byte) int
}

type NaiveSubstitution struct {
	Match    int
	Mismatch int
}

func NewNaiveSubstitution(match, mismatch int) *NaiveSubstitution {
	return &NaiveSubstitution{
		Match:    match,
		Mismatch: mismatch,
	}
}

func (n *NaiveSubstitution) Score(a, b byte) int {
	if a == b {
		return n.Match
	}

	return n.Mismatch
}

type Blosum62 struct{}

func (b *Blosum62) Score(a, c byte) int {
	score, ok := blosum62Table[[2]byte{a, c}]
	if ok {
		return score
	}

	score, ok = blosum62Table[[2]byte{c, a}]
	if ok {
		return score
	}

	return 0
}

type Pam250 struct{}

func (p *Pam250) Score(a, b byte) int {
	score, ok := pam250Table[[2]byte{a, b}]
	if ok {
		return score
	}

	score, ok = pam250Table[[2]byte{a, b}]
	if ok {
		return score
	}

	return 0
}
