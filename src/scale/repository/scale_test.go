package repository

import (
	"reflect"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/scale/src/domain"
	"github.com/scale/src/helper"
	"github.com/stretchr/testify/assert"
)

func init() {
	helper.InitTime()
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	repo := &scaleRepository{}

	type args struct {
		param *domain.Scale
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				param: &domain.Scale{
					Date: date,
				},
			},
			wantErr: false,
			mock:    func() {},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := repo.Create(test.args.param)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestGetScales(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	repo := &scaleRepository{}

	type args struct {
		param domain.Scale
	}

	tests := []struct {
		name       string
		wantResult []domain.Scale
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			wantResult: []domain.Scale{
				{
					Date:       date,
					Min:        45,
					Max:        50,
					Difference: 5,
				},
				{
					Date:       date.AddDate(0, 0, -1),
					Min:        45,
					Max:        50,
					Difference: 5,
				},
				{
					Date:       date.AddDate(0, 0, -3),
					Min:        45,
					Max:        50,
					Difference: 5,
				},
			},
			wantErr: false,
			mock: func() {
				repo.scales = []domain.Scale{
					{
						Date:       date.AddDate(0, 0, -3),
						Min:        45,
						Max:        50,
						Difference: 5,
					},
					{
						Date:       date.AddDate(0, 0, -1),
						Min:        45,
						Max:        50,
						Difference: 5,
					},
					{
						Date:       date,
						Min:        45,
						Max:        50,
						Difference: 5,
					},
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			got, err := repo.GetScales()
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.wantResult, got)
		})
	}
}

func TestGetScale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	repo := &scaleRepository{}

	type args struct {
		date time.Time
	}
	tests := []struct {
		name       string
		args       args
		wantResult []domain.Scale
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			args: args{
				date: date,
			},
			wantResult: []domain.Scale{
				{
					Date:       date,
					Min:        45,
					Max:        50,
					Difference: 5,
				},
			},
			wantErr: false,
			mock: func() {
				repo.scales = []domain.Scale{
					{
						Date:       date,
						Min:        45,
						Max:        50,
						Difference: 5,
					},
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := repo.GetScale(test.args.date)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.wantResult, got)
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	repo := &scaleRepository{}

	type args struct {
		param *domain.Scale
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				param: &domain.Scale{
					Date:       date,
					Min:        45,
					Max:        50,
					Difference: 5,
				},
			},
			wantErr: false,
			mock: func() {
				repo.scales = []domain.Scale{
					{
						Date:       date,
						Min:        45,
						Max:        51,
						Difference: 6,
					},
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := repo.Update(test.args.param)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	repo := &scaleRepository{}

	type args struct {
		date time.Time
	}
	tests := []struct {
		name    string
		args    args
		wantErr bool
		mock    func()
	}{
		{
			name: "success",
			args: args{
				date: date,
			},
			wantErr: false,
			mock: func() {
				repo.scales = []domain.Scale{
					{
						Date:       date,
						Min:        45,
						Max:        50,
						Difference: 5,
					},
				}
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := repo.Delete(test.args.date)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func Test_scaleRepository_GetScales(t *testing.T) {
	type fields struct {
		scales []domain.Scale
	}
	tests := []struct {
		name    string
		fields  fields
		want    []domain.Scale
		wantErr bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &scaleRepository{
				scales: tt.fields.scales,
			}
			got, err := s.GetScales()
			if (err != nil) != tt.wantErr {
				t.Errorf("scaleRepository.GetScales() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("scaleRepository.GetScales() = %v, want %v", got, tt.want)
			}
		})
	}
}
