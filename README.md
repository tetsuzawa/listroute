# listroute

listroute displays all routes in the project files.


## vs grep

You can do the same thing with grep, but you need to write complex regular expressions and shell scripts.


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

### run

```
go vet -vettool=`which listroute` ./...
```

(When using go vet, you must specify the full path to the tool)

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

Source

```
go install github.com/tetsuzawa/listroute/cmd/listroute@latest
```

Binary releases

```
# see releases
curl -L -o listroute.tar.gz https://github.com/tetsuzawa/listroute/releases/download/vX.X.X/listroute_X.X.X_linux_amd64.tar.gz
tar xvf listroute.tar.gz
sudo install listroute /usr/local/bin/

```

# Options

```
`-listroute.matcherFunctions <list of functions>`
```

This option is used to explicitly specify a matcher function.
Multiple functions may be specified, separated by commas (,).
Function names are case-sensitive.
By default, the functions in [labstack/echo/v4](https://github.com/labstack/echo) is registered.

Example
```shell
go vet -vettool=(which listroute) -listroute.matcherFunctions=Get,Post,Handle
```


## Tips

Maybe it could be used for more than just routes :)