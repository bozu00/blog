package context

import (
	"github.com/labstack/echo"
)


type CustomContext struct {
	echo.Context
	UserId int
}

func AdaptCustomContext(h echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, -1} // -1は存在しないユーザ
			return h(cc)
		}
}
