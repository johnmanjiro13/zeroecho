# zeroecho
zeroecho is a [zerolog](https://github.com/rs/zerolog) wrapper for [Echo](https://github.com/labstack/echo).

# Installation
`go get github.com/johnmanjiro13/zeroecho`  

# Usage
```go
package main

import (
    "net/http"
    "os"

    "github.com/johnmanjiro13/zeroecho"
    "github.com/labstack/echo"
)

func main() {
    e := echo.New()

    // Setup logger
    logger := zeroecho.New(os.Stdout, "")
    e.Logger = logger

    // Use logging middleware
    e.Use(zeroecho.RequestLogger(zeroecho.Config{
        Logger:  logger,
    }))

    e.GET("/", func(c echo.Context) error {
        return c.String(http.StatusOK, "Hello, world!")
    })
    e.Logger.Fatal(e.Start(":1323"))
}
```
