package main

import (
	"log"
	"net/http"
	"os"
	"strconv"

	"bitbucket.org/sdwolfe32/ipdata/ipdata"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
)

var (
	// The port that the service will bind to
	port = getenv("PORT", "8080")
	// Whether or not to serve the web frontend
	web, _ = strconv.ParseBool(getenv("WEB", "false"))
)

func main() {
	// Create a new echo Echo and bind middleware
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())

	// Create the ipdata client
	ic, err := ipdata.New()
	if err != nil {
		log.Fatal(err)
	}
	defer ic.Close()

	// Bind all API endpoint handlers
	e.GET("/lookup/:ip", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ic.Lookup(c.Param("ip")))
	})

	// // Serve the static web content
	// e.Static("/", "web")
	// e.Static("/assets", "web/assets")

	// Listen on the passed port
	e.Logger.Fatal(e.Start(":" + port))
}

// getenv retrieves a variable from the environment and falls back to a passed
// default value if the key doesn't exist
func getenv(key, def string) string {
	if value, ok := os.LookupEnv(key); ok {
		return value
	}
	return def
}
