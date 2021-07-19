package requests

import (
	"github.com/go-playground/validator"
	"mutantDetector/models"
)

type MutantDetectorRequest struct {
	Dna []string `json:"dna" validate:"required,min=1"`
}

func NewMutantDetectorRequest() *MutantDetectorRequest {
	return &MutantDetectorRequest{}
}

func (mdr *MutantDetectorRequest) Validate() error {
	v := validator.New()
	return v.Struct(mdr)
}

func (mdr *MutantDetectorRequest) ToModel() *models.Dna {
	return models.NewDna(mdr.Dna)
}
