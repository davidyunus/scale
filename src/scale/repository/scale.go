package repository

import (
	"sort"
	"time"

	"github.com/scale/src/common"
	"github.com/scale/src/domain"
)

type scaleRepository struct {
	scales []domain.Scale // asume this is db
}

func NewScaleRepository() domain.ScaleRepository {
	return &scaleRepository{}
}

func (s *scaleRepository) Create(param *domain.Scale) error {
	s.scales = append(s.scales, *param)
	return nil
}

func (s *scaleRepository) GetScales() ([]domain.Scale, error) {
	sort.Slice(s.scales, func(i, j int) bool {
		return s.scales[i].Date.After(s.scales[j].Date)
	})

	return s.scales, nil
}

func (s *scaleRepository) GetScale(date time.Time) ([]domain.Scale, error) {
	scaleResponse := []domain.Scale{}
	for _, scale := range s.scales {
		if scale.Date.Format(common.TimeLayout) == date.Format(common.TimeLayout) {
			scaleResponse = append(scaleResponse, scale)
		}
	}

	sort.Slice(scaleResponse, func(i, j int) bool {
		return scaleResponse[i].Date.After(scaleResponse[j].Date)
	})
	return scaleResponse, nil
}

func (s *scaleRepository) Update(param *domain.Scale) error {
	for i, scale := range s.scales {
		if scale.Date.Format(common.TimeLayout) == param.Date.Format(common.TimeLayout) {
			s.scales[i].Min = param.Min
			s.scales[i].Max = param.Max
			s.scales[i].Difference = param.Max - param.Min
		}
	}

	return nil
}

func (s *scaleRepository) Delete(date time.Time) error {
	for i, scale := range s.scales {
		if scale.Date.Format(common.TimeLayout) == date.Format(common.TimeLayout) {
			s.scales = remove(s.scales, i)
		}
	}

	return nil
}
