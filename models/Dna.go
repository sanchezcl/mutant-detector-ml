package models

type Dna struct {
	M []string
}

const (
	LetterCount  = 4
	LettersMatch = "ATCG"
)

func NewDna(m []string) *Dna{
	return &Dna{M: m}
}
