package handlers

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/dariuszkorolczukcom/musicGroupApi/pkg/preset"
	"github.com/labstack/echo/v4"
)

func GetPreset(c echo.Context) error {
	id := c.Param("id")
	if id != "" {
		// Get single preset
		result, err := preset.FetchPreset(id)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		response, err := json.Marshal(result)
		if err != nil {
			return c.JSON(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, string(response))
	}

	// Get list of presets
	result, err := preset.FetchPresets()
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	response, err := json.Marshal(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, string(response))
}

func CreatePreset(c echo.Context) error {
	var p preset.Preset
	err := c.Bind(&p)
	fmt.Println(p)

	result, err := preset.CreatePreset(p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	response, err := json.Marshal(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusCreated, string(response))
}

func UpdatePreset(c echo.Context) error {
	id := c.Param("id")

	var p preset.Preset
	err := c.Bind(&p)
	fmt.Println(p)
	result, err := preset.UpdatePreset(id, p)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	response, err := json.Marshal(result)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, string(response))
}

func DeletePreset(c echo.Context) error {
	id := c.Param("id")
	err := preset.DeletePreset(id)
	if err != nil {
		return c.JSON(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, "{\"deleted\":\""+id+"\"}")
}
