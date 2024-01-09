package profile

import (
	"os"

	echojwt "github.com/labstack/echo-jwt"
	"github.com/labstack/echo/v4"
)

func (ph *ProfileHandler) RegisterRoute(e *echo.Echo) {
	jwtMiddleware := echojwt.JWT([]byte(os.Getenv("SECRET_KEY")))

	profileGroup := e.Group("/user")
	profileGroup.Use(jwtMiddleware)
	profileGroup.GET("", ph.GetUserProfile)
	profileGroup.PUT("/profile", ph.UpdateUserProfile)
	profileGroup.PUT("/add/profile", ph.UpdateUserProfile)
	profileGroup.POST("/address", ph.CreateAddressProfile)
	profileGroup.GET("/address", ph.GetAllAddressProfile)
	profileGroup.PUT("/address/:id", ph.UpdateAddressProfile)

	// Get Province and City
	e.GET("/province", ph.GetAllProvince, jwtMiddleware)
	e.GET("/city", ph.GetAllCityByProvince, jwtMiddleware)
}
