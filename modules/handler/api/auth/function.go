package auth

import (
	"net/http"

	ue "basic-coding-kulina/modules/entity/user"

	"github.com/labstack/echo/v4"
)

func (ah *AuthHandler) Register() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.RegisterRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		err := ah.authUsecase.Register(request)
		if err != nil {
			return c.JSON(http.StatusInternalServerError, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusInternalServerError,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Register Sukses",
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthHandler) LoginUser() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		data, role, err := ah.authUsecase.Login(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		if role != "b29112fc-c7b5-4386-a31e-f2c040de7fcb" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Email atau password salah",
				"Status":  http.StatusUnauthorized,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil login",
			"Data":    data,
			"Status":  http.StatusOK,
		})
	}
}

func (ah *AuthHandler) LoginAdmin() echo.HandlerFunc {
	return func(c echo.Context) error {
		var request *ue.LoginRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		data, role, err := ah.authUsecase.Login(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Message": err.Error(),
				"Status":  http.StatusBadRequest,
			})
		}

		if role != "419a8a2d-0abe-4413-ac49-39d33cf9838d" {
			return c.JSON(http.StatusUnauthorized, echo.Map{
				"Message": "Email atau password salah",
				"Status":  http.StatusUnauthorized,
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Message": "Berhasil login",
			"Data":    data,
			"Status":  http.StatusOK,
		})
	}
}
