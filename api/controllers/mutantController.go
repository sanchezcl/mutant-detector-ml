package controllers

import (
	"fmt"
	"github.com/labstack/echo/v4"
	"mutantDetector/api/requests"
	"mutantDetector/services"
	"net/http"
)

func MutantDetector(c echo.Context) error {
	// - MutantDetectorRequest
	// 	* validate parameters
	//	* response errors
	req := requests.NewMutantDetectorRequest()
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}

	// - Call mutant service with data
	mds, err := services.NewMutantDetectorService(req.ToModel())
	if err != nil {
		fmt.Println(struct {message string}{err.Error()})
		return echo.NewHTTPError(http.StatusBadRequest, err.Error(),)
	}
	result, err := mds.AnalyzeDna()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(),)
	}
	//Todo: Implement Persistence

	if !result {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	// - Response json
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func MutantStats(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "ok"})
}
