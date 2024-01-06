package handler

import (
	"github.com/antfley/GOT/model"
	"github.com/antfley/GOT/view/user"
	"github.com/labstack/echo/v4"
)

type UserHandler struct{}

func (h UserHandler) HandleUserShow(c echo.Context) error {
	u := model.User{
		Email: "antfley@gg.com",
	}
	return render(c, user.Show(u))
}
