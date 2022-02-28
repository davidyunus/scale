package main

import (
	"net/http"
	"time"

	"github.com/labstack/echo"
	"github.com/scale/src/domain"
	"github.com/scale/src/helper"
	"github.com/scale/src/scale/handler"
	scalerepo "github.com/scale/src/scale/repository"
	scaleuc "github.com/scale/src/scale/usecase"
)

var (
	scaleUsecase    domain.ScaleUsecase
	scaleRepository domain.ScaleRepository
)

func init() {
	helper.InitTime()
}

func initRepo() {
	scaleRepository = scalerepo.NewScaleRepository()
}

func initUsecase() {
	scaleUsecase = scaleuc.NewScaleUsecase(scaleRepository)

	// for init data
	scaleUsecase.Create(&domain.Scale{
		Date: time.Date(2018, 8, 22, 0, 0, 0, 0, helper.GetLocation()),
		Min:  49,
		Max:  50,
	})
	scaleUsecase.Create(&domain.Scale{
		Date: time.Date(2018, 8, 21, 0, 0, 0, 0, helper.GetLocation()),
		Min:  49,
		Max:  49,
	})
	scaleUsecase.Create(&domain.Scale{
		Date: time.Date(2018, 8, 20, 0, 0, 0, 0, helper.GetLocation()),
		Min:  50,
		Max:  52,
	})
	scaleUsecase.Create(&domain.Scale{
		Date: time.Date(2018, 8, 19, 0, 0, 0, 0, helper.GetLocation()),
		Min:  50,
		Max:  51,
	})
	scaleUsecase.Create(&domain.Scale{
		Date: time.Date(2018, 8, 18, 0, 0, 0, 0, helper.GetLocation()),
		Min:  48,
		Max:  50,
	})
}

func initHTTP() error {
	e := echo.New()
	e.Debug = true

	handler.NewScaleHandler(e, scaleUsecase)
	e.GET("/ping", func(c echo.Context) error {
		return c.JSON(http.StatusOK, helper.Response(200, "Pong", nil, nil))
	})

	return e.Start(":8080")
}

func main() {
	initRepo()
	initUsecase()

	initHTTP()
}
