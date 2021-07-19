package controllers

import (
	"fmt"
	"github.com/go-playground/validator"
	"github.com/labstack/echo/v4"
	"mutantDetector/api/requests"
	"mutantDetector/database"
	"mutantDetector/repositories"
	"mutantDetector/services"
	"net/http"
	"strings"
)

func MutantDetector(c echo.Context) error {
	req := requests.NewMutantDetectorRequest()
	if err := c.Bind(req); err != nil {
		return c.JSON(http.StatusBadRequest, err)
	}
	if err := req.Validate(); err != nil {
		e := err.(validator.ValidationErrors)
		out := map[string]string{
			"message": fmt.Sprintf("%s is %s", strings.ToLower(e[0].Field()), e[0].ActualTag()),
		}
		return c.JSON(http.StatusBadRequest, out)
	}

	dnaRepo := repositories.NewPgGormDnaRepository(database.NewDatabaseConn())
	mds, err := services.NewMutantDetectorService(req.ToModel(), dnaRepo)

	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error(), )
	}

	result, err := mds.AnalyzeDna()
	if err != nil {
		return echo.NewHTTPError(http.StatusInternalServerError, err.Error(), )
	}

	if !result {
		return echo.NewHTTPError(http.StatusForbidden, "Forbidden")
	}

	// - Response json
	return c.JSON(http.StatusOK, map[string]string{"message": "ok"})
}

func MutantStats(c echo.Context) error {
	return c.JSON(http.StatusOK, map[string]string{"data": "ok"})
}
