package router

import (
	"beholder-api/internal/application/models"
	"beholder-api/internal/data/repositories"
	"beholder-api/internal/dtos"
	"beholder-api/internal/utils"

	"github.com/labstack/echo/v4"
)

func CallRouter(r *echo.Echo, repo repositories.CallRepository) {
	g := r.Group("/call")

	g.GET("", PaginationMiddleware(func(c echo.Context) error {
		context := c.(*PaginationContext)
		repo.Get(context.Pagination).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, *f.Code(), f.Message())
			},
			func(calls *[]*models.Call) {
				Response(c, 200, calls)
			},
		)
		return nil
	}))

	g.POST("", func(c echo.Context) error {
		input := dtos.CreateCallDto{}
		err := c.Bind(&input)
		if err != nil {
			return ErrorResponse(c, 400, err.Error())
		}
		repo.Create(input.ToModel()).Fold(
			func(f utils.Failure) {
				ErrorResponse(c, 400, f.Message())
			},
			func(call *models.Call) {
				Response(c, 201, call)
			},
		)
		return nil
	})
}
