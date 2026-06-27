package aligners

import (
	"fmt"
	"maps"
	"slices"
	"strings"
)

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

type FromMatrix struct {
	Name string
}

func NewFromMatrix(name string) (*FromMatrix, error) {
	name = strings.ToLower(name)
	if _, ok := SubstitutionMatrixNameMapping[name]; !ok {
		validOptions := slices.Collect(maps.Keys(SubstitutionMatrixNameMapping))
		return nil, fmt.Errorf("invalid substitution matrix name: %s. valid options: %s", name, strings.Join(validOptions, ", "))
	}

	return &FromMatrix{
		Name: name,
	}, nil
}

func (fm *FromMatrix) Score(a, b byte) int {
	table := SubstitutionMatrixNameMapping[fm.Name]
	score, ok := table[[2]byte{a, b}]
	if ok {
		return score
	}

	score, ok = table[[2]byte{b, a}]
	if ok {
		return score
	}

	return 0
}
