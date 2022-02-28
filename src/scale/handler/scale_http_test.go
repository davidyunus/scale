package handler

import (
	"errors"
	"fmt"
	"net/http"
	"net/http/httptest"
	"strings"
	"testing"
	"time"

	"github.com/golang/mock/gomock"
	"github.com/labstack/echo"
	"github.com/scale/src/common"
	"github.com/scale/src/domain"
	"github.com/scale/src/helper"
	mock_domain "github.com/scale/src/mock"
	"github.com/stretchr/testify/assert"
)

func init() {
	helper.InitTime()
}

func TestCreate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())
	scaleMock := mock_domain.NewMockScaleUsecase(ctrl)

	tests := []struct {
		name       string
		args       string
		wantResult string
		mock       func()
	}{
		{
			name: "success",
			args: `{"date":"2022-02-01","min":45,"max":50}`,
			wantResult: `{"code":200,"message":"Success create scale","data":null,"errors":null}
`,
			mock: func() {
				date, _ = time.Parse(common.TimeLayout, "2022-02-01")
				scaleMock.EXPECT().Create(&domain.Scale{
					Date: date,
					Min:  45,
					Max:  50,
				}).Return(nil)
			},
		},
		{
			name: "error",
			wantResult: `{"code":500,"message":"Failed create scale","data":null,"errors":"some error"}
`,
			args: `{"date":"2022-02-01","min":45,"max":50}`,
			mock: func() {
				date, _ = time.Parse(common.TimeLayout, "2022-02-01")
				scaleMock.EXPECT().Create(&domain.Scale{
					Date: date,
					Min:  45,
					Max:  50,
				}).Return(errors.New("some error"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPost, "/scale", strings.NewReader(test.args))
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := scaleHandler{
				scaleUsecase: scaleMock,
			}

			test.mock()

			if assert.NoError(t, h.Create(c)) {
				assert.Equal(t, test.wantResult, rec.Body.String())
			}
		})
	}
}

func TestGetScales(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())
	scaleMock := mock_domain.NewMockScaleUsecase(ctrl)

	tests := []struct {
		name       string
		wantResult string
		mock       func()
	}{
		{
			name: "success",
			wantResult: `{"code":200,"message":"Success get scales","data":{"scales":[{"date":"2022-02-01T00:00:00+07:00","min":47,"max":50,"difference":3},{"date":"2022-02-01T00:00:00+07:00","min":50,"max":53,"difference":3}],"average":{"min":48.5,"max":51.5,"difference":3}},"errors":null}
`,
			mock: func() {
				scaleMock.EXPECT().GetScales().Return(&domain.ScaleResponse{
					Scales: []domain.Scale{
						{
							Date:       date,
							Min:        47,
							Max:        50,
							Difference: 3,
						},
						{
							Date:       date,
							Min:        50,
							Max:        53,
							Difference: 3,
						},
					},
					Average: &domain.ScaleAverrage{
						Min:        48.5,
						Max:        51.5,
						Difference: 3,
					},
				}, nil)
			},
		},
		{
			name: "error",
			wantResult: `{"code":500,"message":"Failed get scales","data":null,"errors":"some error"}
`,
			mock: func() {
				scaleMock.EXPECT().GetScales().Return(nil, errors.New("some error"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, "/scales", nil)
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := scaleHandler{
				scaleUsecase: scaleMock,
			}

			test.mock()

			if assert.NoError(t, h.GetScales(c)) {
				assert.Equal(t, test.wantResult, rec.Body.String())
			}
		})
	}
}

func TestGetScale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	date := time.Date(2022, 2, 1, 0, 0, 0, 0, helper.GetLocation())
	scaleMock := mock_domain.NewMockScaleUsecase(ctrl)

	tests := []struct {
		name       string
		args       string
		wantResult string
		mock       func()
	}{
		{
			name: "success",
			args: `?date=2022-02-01`,
			wantResult: `{"code":200,"message":"Success get scale","data":[{"date":"2022-02-01T00:00:00+07:00","min":47,"max":50,"difference":3}],"errors":null}
`,
			mock: func() {
				scaleMock.EXPECT().GetScale("2022-02-01").Return([]domain.Scale{
					{
						Date:       date,
						Min:        47,
						Max:        50,
						Difference: 3,
					},
				}, nil)
			},
		},
		{
			name: "error",
			args: `?date=2022-02-01`,
			wantResult: `{"code":500,"message":"Failed get scale","data":null,"errors":"some error"}
`,
			mock: func() {
				scaleMock.EXPECT().GetScale("2022-02-01").Return(nil, errors.New("some error"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodGet, fmt.Sprintf("/scale%v", test.args), nil)
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := scaleHandler{
				scaleUsecase: scaleMock,
			}

			test.mock()

			if assert.NoError(t, h.GetScale(c)) {
				assert.Equal(t, test.wantResult, rec.Body.String())
			}
		})
	}
}

func TestUpdate(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	date, _ := time.Parse(common.TimeLayout, "2022-02-01")
	scaleMock := mock_domain.NewMockScaleUsecase(ctrl)

	tests := []struct {
		name       string
		args       string
		wantResult string
		mock       func()
	}{
		{
			name: "success",
			args: `{"date":"2022-02-01","min":45,"max":50}`,
			wantResult: `{"code":200,"message":"Success update scale","data":null,"errors":null}
`,
			mock: func() {
				scaleMock.EXPECT().Update(&domain.Scale{
					Date: date,
					Min:  45,
					Max:  50,
				}).Return(nil)
			},
		},
		{
			name: "error",
			args: `{"date":"2022-02-01","min":45,"max":50}`,
			wantResult: `{"code":500,"message":"Failed update scale","data":null,"errors":"some error"}
`,
			mock: func() {
				scaleMock.EXPECT().Update(&domain.Scale{
					Date: date,
					Min:  45,
					Max:  50,
				}).Return(errors.New("some error"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodPatch, "/scale", strings.NewReader(test.args))
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := scaleHandler{
				scaleUsecase: scaleMock,
			}

			test.mock()

			if assert.NoError(t, h.Update(c)) {
				assert.Equal(t, test.wantResult, rec.Body.String())
			}
		})
	}
}

func TestDeleteScale(t *testing.T) {
	ctrl := gomock.NewController(t)
	defer ctrl.Finish()

	scaleMock := mock_domain.NewMockScaleUsecase(ctrl)

	tests := []struct {
		name       string
		args       string
		wantResult string
		wantErr    bool
		mock       func()
	}{
		{
			name: "success",
			args: `?date=2022-02-01`,
			wantResult: `{"code":200,"message":"Success delete scale","data":null,"errors":null}
`,
			wantErr: false,
			mock: func() {
				scaleMock.EXPECT().Delete("2022-02-01").Return(nil)
			},
		},
		{
			name: "error",
			args: `?date=2022-02-01`,
			wantResult: `{"code":500,"message":"Failed delete scales","data":null,"errors":"some error"}
`,
			wantErr: true,
			mock: func() {
				scaleMock.EXPECT().Delete("2022-02-01").Return(errors.New("some error"))
			},
		},
	}
	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			e := echo.New()
			req := httptest.NewRequest(http.MethodDelete, fmt.Sprintf("/scale%v", test.args), nil)
			req.Header.Set("content-type", "application/json")
			rec := httptest.NewRecorder()
			c := e.NewContext(req, rec)
			h := scaleHandler{
				scaleUsecase: scaleMock,
			}

			test.mock()

			if assert.NoError(t, h.DeleteScale(c)) {
				assert.Equal(t, test.wantResult, rec.Body.String())
			}
		})
	}
}
