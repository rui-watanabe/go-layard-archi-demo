package main

import (
	// "clean-archi-demo/infrastructure"
	// "fmt"

	"github.com/labstack/echo"
)

func main() {
	e := echo.New()
	// c := infrastructure.NewClient(server string, opts ...infrastructure.ClientOption)

	e.Logger.Fatal(e.Start(":1323"))
}
