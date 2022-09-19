# listroute

listroute displays all routes in the project files.

## Example

```go
package main

import (
	"github.com/labstack/echo/v4"
)

func main() {
	e := echo.New()

	e.GET("/", GETHandler)
	e.POST("/", POSTHandler)
	e.PUT("/", PUTHandler)
	e.DELETE("/", DELETEHandler)
	e.PATCH("/", PATCHHandler)
	e.Static("/foo.jpg", "/public")
	e.File("/bar", "/bar.txt")

}

func GETHandler(c echo.Context) error    { return nil }
func POSTHandler(c echo.Context) error   { return nil }
func PUTHandler(c echo.Context) error    { return nil }
func DELETEHandler(c echo.Context) error { return nil }
func PATCHHandler(c echo.Context) error  { return nil }
```

### output

```
./main.go:10:2: e.GET("/", GETHandler)
./main.go:11:2: e.POST("/", POSTHandler)
./main.go:12:2: e.PUT("/", PUTHandler)
./main.go:13:2: e.DELETE("/", DELETEHandler)
./main.go:14:2: e.PATCH("/", PATCHHandler)
./main.go:15:2: e.Static("/foo.jpg", "/public")
./main.go:16:2: e.File("/bar", "/bar.txt")
```

# Install

```
go install github.com/tetsuzawa/listroute/cmd/listroute@latest
```

# Usage

## bash

```
go vet -vettool=`which listroute` ./...
```

## fish

```
go vet -vettool=(which listroute) ./...
```

## options

```
`-listroute.matcherFunctions <list of functions>`
```

This option is used to explicitly specify a matcher function.
Multiple functions may be specified, separated by commas (,).
By default, the functions in [labstack/echo/v4](https://github.com/labstack/echo) is registered.

Example
```shell
go vet -vettool=(which listroute) -listroute.matcherFunctions=Get,Post,Handle
```


## Tips

Maybe it could be used for more than just routes :)