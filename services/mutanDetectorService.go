package services

import (
	"fmt"
	"mutantDetector/models"
	"mutantDetector/repositories"
	"mutantDetector/validators"
	"reflect"
	"regexp"
	"strings"
)

const (
	charCount = 4
	minMatchMutantCount = 2
)

var (
	regexPattern = fmt.Sprintf("A{%[1]d}|T{%[1]d}|C{%[1]d}|G{%[1]d}", charCount)
)

type MutantDetectorService struct {
	dna *models.Dna
	dnaRepo repositories.DnaRepositoryInterface
}

func NewMutantDetectorService(dna *models.Dna, repo repositories.DnaRepositoryInterface) (*MutantDetectorService, error) {
	dnaVal := validators.NewDnaValidator(dna)
	_, err := dnaVal.Validate()
	if err != nil {
		return nil, err
	}
	return &MutantDetectorService{
		dna: dna,
		dnaRepo: repo,
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

	resDna, err := mds.findDnaByHash()
	if err != nil {
		return false , err
	}

	if !reflect.DeepEqual(models.Dna{}, *resDna) {
		return resDna.IsMutant, nil
	}

	dataGetters := []func()([]string, error){
		mds.getVerticals,
		mds.getHorizontals,
		mds.getDiagonals,
	}

	for _, getter := range dataGetters {
		r, err := getter()
		if err != nil {
			return false, err
		}
		m = append(m, r)
	}

	var result = 0
	for _, m := range m {
		str := strings.Join(m, "|")
		match, _ := regexp.MatchString(regexPattern, str)
		if match {
			result++
		}
	}

	if result >= minMatchMutantCount {
		mds.dna.IsMutant = true
	} else {
		resDna.IsMutant = false
	}

	mds.dna, err = mds.persistResults()
	if err != nil {
		return false, err
	}

	return mds.dna.IsMutant, nil
}

func (mds *MutantDetectorService) findDnaByHash() (*models.Dna, error) {
	return mds.dnaRepo.FindByHash(mds.dna)
}

func (mds *MutantDetectorService) persistResults() (*models.Dna,error) {
	return mds.dnaRepo.Create(mds.dna)
}