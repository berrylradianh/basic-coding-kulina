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

		if role != 2 {
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

		if role != 1 {
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

func (ah *AuthHandler) ForgotPassword() echo.HandlerFunc {
	return func(c echo.Context) error {

		var request ue.ForgotPassRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		email, err := ah.authUsecase.ForgotPassword(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Berhasil mengirim kode otp",
			"Email":   email,
		})
	}
}
func (ah *AuthHandler) VerifOtp() echo.HandlerFunc {
	return func(c echo.Context) error {

		var request ue.VerifOtp
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		err := ah.authUsecase.VerifOtp(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Berhasil memverifikasi",
		})
	}
}
func (ah *AuthHandler) ChangePassword() echo.HandlerFunc {
	return func(c echo.Context) error {

		var request ue.RecoveryRequest
		if err := c.Bind(&request); err != nil {
			return echo.NewHTTPError(http.StatusBadRequest, map[string]interface{}{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		err := ah.authUsecase.ChangePassword(request)
		if err != nil {
			return c.JSON(http.StatusBadRequest, echo.Map{
				"Status":  http.StatusBadRequest,
				"Message": err.Error(),
			})
		}

		return c.JSON(http.StatusOK, map[string]interface{}{
			"Status":  http.StatusOK,
			"Message": "Berhasil mengganti password",
		})
	}
}
