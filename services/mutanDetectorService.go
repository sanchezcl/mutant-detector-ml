package services

import (
	"fmt"
	"mutantDetector/models"
	"mutantDetector/validators"
	"regexp"
	"strings"
)

const (
	charCount = 4
	minMatchMutantCount = 2
)

var regexPattern = fmt.Sprintf("A{%[1]d}|T{%[1]d}|C{%[1]d}|G{%[1]d}", charCount)

type MutantDetectorService struct {
	dna *models.Dna
}

func NewMutantDetectorService(dna *models.Dna) (*MutantDetectorService, error) {
	dnaVal := validators.NewDnaValidator(dna)
	_, err := dnaVal.Validate()
	if err != nil {
		return nil, err
	}
	return &MutantDetectorService{
		dna: dna,
	}, nil
}

func (mds *MutantDetectorService) getVerticals() ([]string, error) {
	m := mds.dna.M
	result := make([]string, len(m))
	for _, rows := range m {
		for i, elm := range rows {
			result[i] += string(elm)
		}
	}

	return result, nil
}

func (mds *MutantDetectorService) getHorizontals() ([]string, error) {
	return mds.dna.M, nil
}

func (mds *MutantDetectorService) getDiagonals() ([]string, error) {
	var (
		diagonal  string
		iDiagonal string
	)

	m := mds.dna.M
	vLen := len(mds.dna.M)

	for k, rows := range m {
		for i, elm := range rows {
			if k == i {
				diagonal += string(elm)
			}
			if i+k == vLen-1 {
				iDiagonal += string(elm)
			}
		}
	}

	return []string{diagonal, iDiagonal}, nil
}

func (mds *MutantDetectorService) AnalyzeDna() (bool, error) {
	var m [][]string
	mV, err := mds.getVerticals()
	if err != nil {
		return false, err
	}
	mH, err := mds.getHorizontals()
	if err != nil {
		return false, err
	}
	mD, err := mds.getDiagonals()
	if err != nil {
		return false, err
	}
	m = append(m, mV, mH, mD)
	var result = 0
	for _, m := range m {
		str := strings.Join(m, "|")
		match, _ := regexp.MatchString(regexPattern, str)
		if match {
			result++
		}
	}

	if result < minMatchMutantCount {
		return false, nil
	}
	return true, nil
}
