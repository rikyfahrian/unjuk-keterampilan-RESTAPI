package delivery

import (
	"fmt"
	"net/http"
	"restapi-altera/src/helper"

	"github.com/labstack/echo/v4"
)

func (d *delivery) GetUserByID(c echo.Context) error {

	ctx := c.Request().Context()

	cookie, err := c.Cookie("jwt")
	if err != nil {
		return helper.WriteResponse(c, http.StatusUnauthorized, "UNAUTHORIZED", nil)
	}

	data, err := d.usecase.GetUserByID(ctx, cookie.Value)
	if err != nil {
		return helper.WriteResponse(c, http.StatusInternalServerError, err.Error(), nil)
	}

	return helper.WriteResponse(c, http.StatusOK, fmt.Sprintf("hi %s", data.Name), data)

}
