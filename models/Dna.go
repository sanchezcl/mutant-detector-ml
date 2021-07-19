package models

import (
	"crypto/sha1"
	"fmt"
	"strings"
)

type Dna struct {
	Base
	M        []string `gorm:"-"`
	Hash     string
	IsMutant bool
}

const (
	LetterCount  = 4
	LettersMatch = "ATCG"
)

func NewDna(m []string) *Dna {
	dna := &Dna{M: m}
	dna.Hash = dna.parseDnaHash()
	return dna
}

func (dna *Dna) parseDnaHash() string{
	str := strings.Join(dna.M, "")
	h := sha1.New()
	h.Write([]byte(str))
	return fmt.Sprintf("%x", h.Sum(nil))
}
