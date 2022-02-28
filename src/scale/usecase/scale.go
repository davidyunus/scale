package usecase

import (
	"errors"
	"math"
	"time"

	"github.com/scale/src/common"
	"github.com/scale/src/domain"
)

type scaleUsecase struct {
	scaleRepository domain.ScaleRepository
}

func NewScaleUsecase(scaleRepository domain.ScaleRepository) domain.ScaleUsecase {
	return &scaleUsecase{
		scaleRepository: scaleRepository,
	}
}

func (s *scaleUsecase) Create(param *domain.Scale) error {
	if param.Max < param.Min {
		return errors.New("Min. greater than max.")
	}
	param.Difference = param.Max - param.Min

	err := s.scaleRepository.Create(param)
	if err != nil {
		return err
	}
	return nil
}

func (s *scaleUsecase) GetScales() (*domain.ScaleResponse, error) {
	scales, err := s.scaleRepository.GetScales()
	if err != nil {
		return nil, err
	}
	scaleResponse := &domain.ScaleResponse{
		Scales: scales,
	}

	var minTotal int
	var maxTotal int
	for _, scale := range scales {
		minTotal += scale.Min
		maxTotal += scale.Max
	}

	avgMin := float64(minTotal) / float64(len(scales))
	avgMax := float64(maxTotal) / float64(len(scales))
	diff := avgMax - avgMin
	avg := &domain.ScaleAverrage{
		Min:        math.Round(avgMin*10) / 10,
		Max:        math.Round(avgMax*10) / 10,
		Difference: math.Round(diff*10) / 10,
	}
	scaleResponse.Average = avg

	return scaleResponse, err
}

func (s *scaleUsecase) GetScale(date string) ([]domain.Scale, error) {
	d, err := time.Parse(common.TimeLayout, date)
	if err != nil {
		return nil, err
	}
	scales, err := s.scaleRepository.GetScale(d)
	if err != nil {
		return nil, err
	}

	return scales, err
}

func (s *scaleUsecase) Update(param *domain.Scale) error {
	if param.Max < param.Min {
		return errors.New("Min. greater than max.")
	}
	param.Difference = param.Max - param.Min

	err := s.scaleRepository.Update(param)
	if err != nil {
		return err
	}

	return nil
}

func (s *scaleUsecase) Delete(date string) error {
	d, err := time.Parse(common.TimeLayout, date)
	if err != nil {
		return err
	}
	err = s.scaleRepository.Delete(d)
	if err != nil {
		return err
	}
	return nil
}
