package controller

import "github.com/labstack/echo/v4"

func Login(c echo.Context) error {
	req := loginReq{}
	if err := c.Bind(&req); err != nil {
		return err
	}

	return c.String(200, "Hello, World!")
}
