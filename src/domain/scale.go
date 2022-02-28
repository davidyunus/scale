package domain

import "time"

type (
	ScaleUsecase interface {
		Create(param *Scale) error
		GetScales() (*ScaleResponse, error)
		GetScale(date string) ([]Scale, error)
		Update(param *Scale) error
		Delete(date string) error
	}

	ScaleRepository interface {
		Create(param *Scale) error
		GetScales() ([]Scale, error)
		GetScale(date time.Time) ([]Scale, error)
		Update(param *Scale) error
		Delete(date time.Time) error
	}
)

type Scale struct {
	Date       time.Time `json:"date"`
	Min        int       `json:"min"`
	Max        int       `json:"max"`
	Difference int       `json:"difference"`
}

type ScaleParam struct {
	Date string `json:"date"`
	Min  int    `json:"min"`
	Max  int    `json:"max"`
}

type ScaleAverrage struct {
	Min        float64 `json:"min"`
	Max        float64 `json:"max"`
	Difference float64 `json:"difference"`
}

type ScaleResponse struct {
	Scales  []Scale        `json:"scales"`
	Average *ScaleAverrage `json:"average"`
}
