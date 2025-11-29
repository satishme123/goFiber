package main

import (
	cookies "myapp/Cookies"
	"myapp/controller"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &controller.CustomContext{
				Context: c,
			}
			return next(cc)
		}
	})
	e.GET("/", controller.HelloWorld)
	e.GET("/:search", controller.ParamsEcho)
	e.GET("/query", controller.QueryEcho)
	e.POST("/user", controller.UserRegister)
	e.GET("/set-cookie", cookies.WriteCookie)
	e.GET("/get-cookie", cookies.ReadCookie)
	e.GET("/get-all-cookie", cookies.ReadAllCookies)
	e.Logger.Fatal(e.Start(":1323"))
}
