package usecase

import (
	"errors"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/scale/src/common"
	"github.com/scale/src/domain"
	"github.com/scale/src/helper"
	mock_domain "github.com/scale/src/mock"
	"github.com/stretchr/testify/assert"
)

func init() {
	helper.InitTime()
}

func TestNewScaleUsecase(t *testing.T) {
	NewScaleUsecase(nil)
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date, _ := time.Parse(common.TimeLayout, "2022-02-01")

	scaleMock := mock_domain.NewMockScaleRepository(ctrl)

	uc := &scaleUsecase{
		scaleRepository: scaleMock,
	}

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
					Min:  45,
					Max:  50,
				},
			},
			wantErr: false,
			mock: func() {
				scaleMock.EXPECT().Create(&domain.Scale{
					Date:       date,
					Min:        45,
					Max:        50,
					Difference: 5,
				}).Return(nil)
			},
		},
		{
			name: "error",
			args: args{
				param: &domain.Scale{
					Date: date,
					Min:  45,
					Max:  50,
				},
			},
			wantErr: true,
			mock: func() {
				scaleMock.EXPECT().Create(&domain.Scale{
					Date:       date,
					Min:        45,
					Max:        50,
					Difference: 5,
				}).Return(errors.New("some error"))
			},
		},
		{
			name: "error",
			args: args{
				param: &domain.Scale{
					Date: date,
					Min:  50,
					Max:  45,
				},
			},
			wantErr: true,
			mock:    func() {},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()

			err := uc.Create(test.args.param)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestGetScales(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	scaleMock := mock_domain.NewMockScaleRepository(ctrl)

	uc := &scaleUsecase{
		scaleRepository: scaleMock,
	}

	type args struct {
		param domain.Scale
	}
	tests := []struct {
		name       string
		wantResult *domain.ScaleResponse
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			wantResult: &domain.ScaleResponse{
				Scales: []domain.Scale{
					{
						Date:       date,
						Min:        45,
						Max:        50,
						Difference: 5,
					},
					{
						Date:       date,
						Min:        47,
						Max:        52,
						Difference: 5,
					},
				},
				Average: &domain.ScaleAverrage{
					Min:        46,
					Max:        51,
					Difference: 5,
				},
			},
			wantErr: false,
			mock: func() {
				scaleMock.EXPECT().GetScales().Return([]domain.Scale{
					{
						Date:       date,
						Min:        45,
						Max:        50,
						Difference: 5,
					},
					{
						Date:       date,
						Min:        47,
						Max:        52,
						Difference: 5,
					},
				}, nil)
			},
		},
		{
			name:       "error",
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				scaleMock.EXPECT().GetScales().Return(nil, errors.New("some error"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			t.Run(test.name, func(t *testing.T) {
				test.mock()

				got, err := uc.GetScales()
				assert.Equal(t, test.wantErr, err != nil)
				assert.Equal(t, test.wantResult, got)
			})
		})
	}
}

func TestGetScale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date, _ := time.Parse(common.TimeLayout, "2022-02-01")

	scaleMock := mock_domain.NewMockScaleRepository(ctrl)

	uc := &scaleUsecase{
		scaleRepository: scaleMock,
	}

	type args struct {
		date string
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
				date: "2022-02-01",
			},
			wantResult: []domain.Scale{
				{
					Date:       date,
					Min:        45,
					Max:        45,
					Difference: 5,
				},
			},
			wantErr: false,
			mock: func() {
				scaleMock.EXPECT().GetScale(date).Return([]domain.Scale{
					{
						Date:       date,
						Min:        45,
						Max:        45,
						Difference: 5,
					},
				}, nil)
			},
		},
		{
			name: "error",
			args: args{
				date: "2022-02-01",
			},
			wantResult: nil,
			wantErr:    true,
			mock: func() {
				scaleMock.EXPECT().GetScale(date).Return(nil, errors.New("some error"))
			},
		},
		{
			name: "error param",
			args: args{
				date: "date",
			},
			wantResult: nil,
			wantErr:    true,
			mock:       func() {},
		},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			got, err := uc.GetScale(test.args.date)
			assert.Equal(t, test.wantErr, err != nil)
			assert.Equal(t, test.wantResult, got)
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())

	scaleMock := mock_domain.NewMockScaleRepository(ctrl)

	uc := &scaleUsecase{
		scaleRepository: scaleMock,
	}

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
					Min:        47,
					Max:        50,
					Difference: 3,
				},
			},
			wantErr: false,
			mock: func() {
				scaleMock.EXPECT().Update(&domain.Scale{
					Date:       date,
					Min:        47,
					Max:        50,
					Difference: 3,
				}).Return(nil)
			},
		},
		{
			name: "error",
			args: args{
				param: &domain.Scale{
					Date:       date,
					Min:        47,
					Max:        50,
					Difference: 3,
				},
			},
			wantErr: true,
			mock: func() {
				scaleMock.EXPECT().Update(&domain.Scale{
					Date:       date,
					Min:        47,
					Max:        50,
					Difference: 3,
				}).Return(errors.New("some error"))
			},
		},
		{
			name: "error",
			args: args{
				param: &domain.Scale{
					Date: date,
					Min:  50,
					Max:  45,
				},
			},
			wantErr: true,
			mock:    func() {},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := uc.Update(test.args.param)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}

func TestDelete(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()
	date, _ := time.Parse(common.TimeLayout, "2022-02-01")

	scaleMock := mock_domain.NewMockScaleRepository(ctrl)

	uc := &scaleUsecase{
		scaleRepository: scaleMock,
	}

	type args struct {
		date string
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
				date: "2022-02-01",
			},
			wantErr: false,
			mock: func() {
				scaleMock.EXPECT().Delete(date).Return(nil)
			},
		},
		{
			name: "error",
			args: args{
				date: "2022-02-01",
			},
			wantErr: true,
			mock: func() {
				scaleMock.EXPECT().Delete(date).Return(errors.New("some error"))
			},
		},
		{
			name: "error",
			args: args{
				date: "date",
			},
			wantErr: true,
			mock:    func() {},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			test.mock()
			err := uc.Delete(test.args.date)
			assert.Equal(t, test.wantErr, err != nil)
		})
	}
}
