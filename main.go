package main

import (
	"demo_gateway/pb"
	"demo_gateway/service"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	personService := service.NewServiceClient()

	e.GET("/person", func(c echo.Context) error {
		res, err := personService.Person.GetPerson(c.Request().Context(), &pb.GetPersonRequest{
			Name: "Tam",
			Age:  "25",
		})
		if err != nil {
			return c.JSON(500, "ไม่พร้อมให้บริการ")
		}
		return c.JSON(200, res)
	})
	e.Logger.Fatal(e.Start(":1323"))
}
