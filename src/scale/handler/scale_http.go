package handler

import (
	"net/http"
	"time"

	"github.com/scale/src/common"
	"github.com/scale/src/domain"
	"github.com/scale/src/helper"

	"github.com/labstack/echo"
)

type scaleHandler struct {
	scaleUsecase domain.ScaleUsecase
}

func NewScaleHandler(e *echo.Echo, scaleUsecase domain.ScaleUsecase) {
	handler := &scaleHandler{
		scaleUsecase: scaleUsecase,
	}

	e.POST("/scale", handler.Create)
	e.GET("/scale", handler.GetScale)
	e.GET("/scales", handler.GetScales)
	e.DELETE("/scale", handler.DeleteScale)
	e.PATCH("/scale", handler.Update)
}

func (h *scaleHandler) Create(c echo.Context) error {
	c.Echo().Validator = helper.NewValidator()
	payload := &domain.ScaleParam{}
	err := c.Bind(payload)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.Response(code, "Failed create scale", nil, err.Error()))
	}
	date, err := time.Parse(common.TimeLayout, payload.Date)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.Response(code, "Failed create scale", nil, err.Error()))
	}
	err = c.Validate(payload)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.Response(code, "Failed create scale", nil, err.Error()))
	}

	err = h.scaleUsecase.Create(&domain.Scale{
		Date: date,
		Min:  payload.Min,
		Max:  payload.Max,
	})
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed create scale", nil, err.Error()))
	}

	data := helper.Response(200, "Success create scale", nil, nil)
	return c.JSON(http.StatusOK, data)
}

func (h *scaleHandler) GetScales(c echo.Context) error {
	scales, err := h.scaleUsecase.GetScales()
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed get scales", nil, err.Error()))
	}

	data := helper.Response(200, "Success get scales", scales, nil)
	return c.JSON(http.StatusOK, data)
}

func (h *scaleHandler) GetScale(c echo.Context) error {
	query := c.Request().URL.Query()
	date := query.Get("date")
	_, err := time.Parse(common.TimeLayout, date)
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed delete scale", nil, err.Error()))
	}

	scales, err := h.scaleUsecase.GetScale(date)
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed get scale", nil, err.Error()))
	}

	data := helper.Response(200, "Success get scale", scales, nil)
	return c.JSON(http.StatusOK, data)
}

func (h *scaleHandler) Update(c echo.Context) error {
	c.Echo().Validator = helper.NewValidator()
	payload := &domain.ScaleParam{}
	err := c.Bind(payload)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.Response(code, "Failed update scale", nil, err.Error()))
	}
	date, err := time.Parse(common.TimeLayout, payload.Date)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.Response(code, "Failed update scale", nil, err.Error()))
	}
	err = c.Validate(payload)
	if err != nil {
		code := http.StatusBadRequest
		return c.JSON(code, helper.Response(code, "Failed update scale", nil, err.Error()))
	}

	err = h.scaleUsecase.Update(&domain.Scale{
		Date: date,
		Min:  payload.Min,
		Max:  payload.Max,
	})
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed update scale", nil, err.Error()))
	}

	data := helper.Response(200, "Success update scale", nil, nil)
	return c.JSON(http.StatusOK, data)
}

func (h *scaleHandler) DeleteScale(c echo.Context) error {
	query := c.Request().URL.Query()
	date := query.Get("date")
	_, err := time.Parse(common.TimeLayout, date)
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed delete scales", nil, err.Error()))
	}

	err = h.scaleUsecase.Delete(date)
	if err != nil {
		code := helper.GetStatusCode(err)
		return c.JSON(code, helper.Response(code, "Failed delete scales", nil, err.Error()))
	}

	data := helper.Response(200, "Success delete scale", nil, nil)
	return c.JSON(http.StatusOK, data)
}
