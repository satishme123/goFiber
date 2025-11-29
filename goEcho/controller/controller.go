package controller

import (
	"fmt"
	"net/http"

	"github.com/labstack/echo/v4"
)

type User struct {
	Name  string `json:"name"`
	Email string `json:"email"`
}

type CustomContext struct {
	echo.Context
}

func (c *CustomContext) Foo() {
	println("foo")
}
func (c *CustomContext) Bar() {
	println("foo")
}
func UserRegister(c echo.Context) error {
	var u User
	err := c.Bind(&u)
	if err != nil {
		return err
	}
	cc := c.(*CustomContext)
	cc.Foo()
	cc.Bar()
	fmt.Println(u)
	return c.JSON(http.StatusOK, u)
}

func HelloWorld(c echo.Context) error {
	return c.String(http.StatusOK, "Hellow world")
}
func ParamsEcho(c echo.Context) error {
	val := c.Param("search")
	return c.String(http.StatusOK, val)
}
func QueryEcho(c echo.Context) error {
	val := c.QueryParam("name")
	fmt.Println(val)
	return c.String(http.StatusOK, val)
}
