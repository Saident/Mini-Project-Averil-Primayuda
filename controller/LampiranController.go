package controller

import (
	"encoding/base64"
	"fmt"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/Saident/Mini-Project-Averil-Primayuda/utils"

	"github.com/labstack/echo"
)

func PostLampiranController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := int(claims["id"].(float64))

	lampiran_tipe, err := strconv.Atoi(c.Param("lampiran_tipe"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	form, err := c.MultipartForm()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	file, bool := form.File["file"]
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	f, err := file[0].Open()
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}

	base64Data := base64.StdEncoding.EncodeToString(data)
	if role == "user" {
		lampirans := &model.Lampiran{
			Lampiran_tipe:    lampiran_tipe,
			Lampiran_content: base64Data,
			UserID:           user_id,
		}

		if err := config.DB.Save(&lampirans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "File uploaded successfully.",
		})
	}
	return echo.ErrForbidden
}

func GetLampiranListController(c echo.Context) error {
	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var lampirans []model.Lampiran
		if err := config.DB.Where("user_id = ?", user_id).Find(&lampirans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success get all lamarans",
			"lampiran": lampirans,
		})
	}
	return echo.ErrForbidden
}

func GetUserLampiranByUserController(c echo.Context) error {
	lampiran_id, err := strconv.Atoi(c.Param("lampiran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	user_id := claims["id"].(float64)

	if role == "user" {
		var lampirans model.Lampiran
		if err := config.DB.Where("user_id = ?", user_id).First(&lampirans, lampiran_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		decoded, err := base64.StdEncoding.DecodeString(lampirans.Lampiran_content)
		if err != nil {
			panic(err)
		}

		path := ("img/" + "users" + fmt.Sprintf("%f", user_id) + "lampiran" + fmt.Sprintf("%d", lampirans.ID) + ".jpg")

		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(decoded); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"lampirans": lampirans.Lampiran_tipe,
			"content":   path,
		})
	}
	return echo.ErrForbidden
}

func GetUserLampiranByPerusahaanController(c echo.Context) error {
	lamaran_id, err := strconv.Atoi(c.Param("lamaran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := utils.GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var lamarans model.Lamaran
		if err := config.DB.Where("id = ? AND perusahaan_id = ?", lamaran_id, perusahaan_id).First(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var users model.User
		if err := config.DB.Where("id = ?", lamarans.UserID).First(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var lampirans model.Lampiran
		if err := config.DB.Where("user_id = ?", users.ID).First(&lampirans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		decoded, err := base64.StdEncoding.DecodeString(lampirans.Lampiran_content)
		if err != nil {
			panic(err)
		}

		path := ("img/" + "users" + fmt.Sprintf("%d", users.ID) + "lampiran" + fmt.Sprintf("%d", lampirans.ID) + ".jpg")

		f, err := os.Create(path)
		if err != nil {
			panic(err)
		}
		defer f.Close()

		if _, err := f.Write(decoded); err != nil {
			panic(err)
		}
		if err := f.Sync(); err != nil {
			panic(err)
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"lampirans": lampirans.Lampiran_tipe,
			"content":   path,
		})
	}
	return echo.ErrForbidden
}
