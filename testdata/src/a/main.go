package main

import (
	"a/subpkg"

	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", GETHandler)          // want `e.GET\("/", GETHandler\)`
	e.POST("/", POSTHandler)        // want `e.POST\("/", POSTHandler\)`
	e.PUT("/", PUTHandler)          // want `e.PUT\("/", PUTHandler\)`
	e.DELETE("/", DELETEHandler)    // want `e.DELETE\("/", DELETEHandler\)`
	e.PATCH("/", PATCHHandler)      // want `e.PATCH\("/", PATCHHandler\)`
	e.Static("/foo.jpg", "/public") // want `e.Static\("/foo.jpg", "/public"\)`
	e.File("/bar", "/bar.txt")      // want `e.File\("/bar", "/bar.txt"\)`

	// TODO AssignStmtに対応する
	sub := e.Group("/sub")                   // want `e\.Group\("/sub"\)`
	sub.GET("/", subpkg.SubGETHandler)       // want `sub.GET\("/", subpkg\.SubGETHandler\)`
	sub.POST("/", subpkg.SubPOSTHandler)     // want `sub.POST\("/", subpkg\.SubPOSTHandler\)`
	sub.PUT("/", subpkg.SubPUTHandler)       // want `sub.PUT\("/", subpkg\.SubPUTHandler\)`
	sub.DELETE("/", subpkg.SubDELETEHandler) // want `sub.DELETE\("/", subpkg\.SubDELETEHandler\)`
	sub.PATCH("/", subpkg.SubPATCHHandler)   // want `sub.PATCH\("/", subpkg\.SubPATCHHandler\)`

}

func GETHandler(c echo.Context) error    { return nil }
func POSTHandler(c echo.Context) error   { return nil }
func PUTHandler(c echo.Context) error    { return nil }
func DELETEHandler(c echo.Context) error { return nil }
func PATCHHandler(c echo.Context) error  { return nil }
