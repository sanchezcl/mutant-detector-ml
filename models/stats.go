package models

type Stats struct {
	CountMutantDna int64   `json:"count_mutant_dna"`
	CountHumanDna  int64   `json:"count_human_dna"`
	Ratio          float32 `json:"ratio, float32"`
}

func NewStats(m, h int64) *Stats {
	s := &Stats{
		CountMutantDna: m,
		CountHumanDna:  h,
	}
	s.Ratio = s.calculateRatio()
	return s
}

func (s *Stats) calculateRatio() float32 {
	if s.CountHumanDna == 0 {
		return float32(s.CountMutantDna)
	}
	return float32(float32(s.CountMutantDna) / float32(s.CountHumanDna))
}
