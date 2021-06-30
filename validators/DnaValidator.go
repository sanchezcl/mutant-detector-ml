package validators

import (
	"errors"
	"fmt"
	"mutantDetector/models"
	"regexp"
	"strings"
)

type DnaValidator struct {
	dna *models.Dna
}

var regexPattern = fmt.Sprintf("[^%s]+", models.LettersMatch)

func NewDnaValidator(dna *models.Dna) *DnaValidator {
	return &DnaValidator{dna,}
}

func (dnaV *DnaValidator) Validate() (bool, error) {
	if _, err := dnaV.validateEmpty(); err != nil {
		return false, err
	}
	if _, err := dnaV.validateRowLen(); err != nil {
		return false, err
	}
	if _, err := dnaV.validateLetters(); err != nil {
		return false, err
	}
	return true, nil
}

func (dnaV *DnaValidator) validateEmpty() (bool, error) {
	if len(dnaV.dna.M) == 0 {
		return false, errors.New("dna is empty")
	}
	return true, nil
}

func (dnaV *DnaValidator) validateLetters() (bool, error) {
	s := strings.Join(dnaV.dna.M, "")
	match, _ := regexp.MatchString(regexPattern, s)
	if match {
		return false, errors.New("letters not match")
	}
	return true, nil
}

func (dnaV *DnaValidator) validateRowLen() (bool, error) {
	m := dnaV.dna.M
	for i := 1; i < len(m); i++ {
		if len(m[i]) != len(m[i-1]) || models.LetterCount > len(m[i]) {
			return false, errors.New("wrong matrix length")
		}
	}
	return true, nil
}