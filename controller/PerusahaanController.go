package controller

import (
	"encoding/base64"
	"fmt"
	"net/http"
	"os"
	"strconv"

	"github.com/Saident/Mini-Project-Averil-Primayuda/config"
	"github.com/Saident/Mini-Project-Averil-Primayuda/model"
	"github.com/labstack/echo"
)

func GetPerusahaansController(c echo.Context) error {
	var perusahaans []model.Perusahaan

	if err := config.DB.Find(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success get all perusahaans",
		"perusahaans": perusahaans,
	})
}

func GetPerusahaanController(c echo.Context) error {
	claims, err := GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var perusahaans model.Perusahaan
		if err := config.DB.First(&perusahaans, perusahaan_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "success get perusahaans by id",
			"perusahaans": perusahaans,
		})
	}
	return echo.ErrForbidden
}

func CreatePerusahaanController(c echo.Context) error {
	perusahaans := model.Perusahaan{}
	c.Bind(&perusahaans)

	if err := config.DB.Save(&perusahaans).Error; err != nil {
		return echo.NewHTTPError(http.StatusBadRequest, err.Error())
	}
	return c.JSON(http.StatusOK, map[string]interface{}{
		"message":     "success create new perusahaans",
		"perusahaans": perusahaans,
	})
}

func UpdatePerusahaanController(c echo.Context) error {
	claims, err := GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var perusahaans model.Perusahaan
		if err := config.DB.First(&perusahaans, perusahaan_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusNotFound, err.Error())
		}

		c.Bind(&perusahaans)

		if err := config.DB.Save(&perusahaans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":     "success update perusahaan by id",
			"perusahaans": perusahaans,
		})
	}
	return echo.ErrForbidden
}

func DeletePerusahaanController(c echo.Context) error {
	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var perusahaan model.Perusahaan
		if err := config.DB.First(&perusahaan, perusahaan_id).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		if err := config.DB.Delete(&perusahaan).Error; err != nil {
			return echo.NewHTTPError(http.StatusInternalServerError, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success delete user by id",
		})
	}
	return echo.ErrForbidden
}

func PostJobsController(c echo.Context) error {
	jobs := model.Jobs{}

	claims, err := GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		jobs.PerusahaanID = int(perusahaan_id)
		jobs.Status = "Belum Divalidasi"
		c.Bind(&jobs)

		if err := config.DB.Save(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success create new jobs",
			"jobs":    jobs,
		})
	}
	return echo.ErrForbidden
}

func GetJobByPerusahaanController(c echo.Context) error {
	claims, err := GetJwtClaims(c)
	if !err {
		return echo.NewHTTPError(http.StatusBadRequest)
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	var jobs []model.Jobs
	if role == "perusahaan" {
		if err := config.DB.Where("perusahaan_id = ?", perusahaan_id).Find(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success get jobs by perusahaan id",
			"jobs":    jobs,
		})
	}
	return echo.ErrForbidden
}

func UpdateJobByPerusahaanController(c echo.Context) error {
	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	job_id, err := strconv.Atoi(c.Param("job_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	if role == "perusahaan" {
		var jobs model.Jobs
		if err := config.DB.Where("job_id = ? AND perusahaan_id = ?", job_id, perusahaan_id).First(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Bind(&jobs)

		if err := config.DB.Save(&jobs).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message": "success update job by id",
			"job":     jobs,
		})
	}
	return echo.ErrForbidden
}

func GetAllLamaranByPerusahaanController(c echo.Context) error {
	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var lamarans []model.Lamaran
		if err := config.DB.Where("perusahaan_id = ?", perusahaan_id).Find(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success get all lamarans",
			"lamarans": lamarans,
		})
	}
	return echo.ErrForbidden
}

func GetLamaranByIdController(c echo.Context) error {
	lamaran_id, err := strconv.Atoi(c.Param("lamaran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var lamarans model.Lamaran
		var users model.User
		if err := config.DB.Where("perusahaan_id = ? AND id = ?", perusahaan_id, lamaran_id).First(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		if err := config.DB.First(&users, lamarans.UserID).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}
		return c.JSON(http.StatusOK, map[string]interface{}{
			"lamaran": lamarans,
			"pelamar": users,
		})
	}
	return echo.ErrForbidden
}

func ValidateLamaranController(c echo.Context) error {
	lamaran_id, err := strconv.Atoi(c.Param("lamaran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := int(claims["id"].(float64))

	if role == "perusahaan" {
		var lamarans model.Lamaran
		if err := config.DB.Where("perusahaan_id = ? AND id = ?", perusahaan_id, lamaran_id).First(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		c.Bind(&lamarans)

		if err := config.DB.Save(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"message":  "success save lamarans",
			"lamarans": lamarans,
		})
	}
	return echo.ErrForbidden
}

func GetUserLampiranByPerusahaanController(c echo.Context) error {
	lamaran_id, err := strconv.Atoi(c.Param("lamaran_id"))
	if err != nil {
		echo.NewHTTPError(http.StatusBadRequest, "messages: invalid id parameter")
	}

	claims, bool := GetJwtClaims(c)
	if !bool {
		return echo.NewHTTPError(http.StatusBadRequest, "messages: invalid JWT")
	}
	role := claims["role"].(string)
	perusahaan_id := claims["id"].(float64)

	if role == "perusahaan" {
		var lamarans model.Lamaran
		if err := config.DB.Where("id = ? AND perusahaan_id = ?", lamaran_id, perusahaan_id).Find(&lamarans).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var users model.User
		if err := config.DB.Where("id = ?", lamarans.UserID).Find(&users).Error; err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, err.Error())
		}

		var lampirans model.Lampiran
		if err := config.DB.Where("user_id = ?", users.ID).Find(&lampirans).Error; err != nil {
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
			"message":   "success get all lamarans",
			"lampirans": lampirans.Lampiran_tipe,
			"content":   path,
		})
	}
	return echo.ErrForbidden
}
