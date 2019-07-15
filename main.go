package main

import (
	"log"
	"net/http"
	"os"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"s32x.com/ipdata/ipdata"
)

var (
	port     = getenv("PORT", "8080")
	cityPath = getenv("CITY_PATH", "./db/city.tar.gz")
	asnPath  = getenv("ASN_PATH", "./db/asn.tar.gz")
)

func main() {
	// Create the ipdata client
	ic, err := ipdata.NewClient(cityPath, asnPath)
	if err != nil {
		log.Fatal(err)
	}
	defer ic.Close()

	// Create a new echo Echo and bind all middleware
	e := echo.New()
	e.HideBanner = true

	// Bind middleware
	e.Pre(middleware.RemoveTrailingSlashWithConfig(
		middleware.TrailingSlashConfig{
			RedirectCode: http.StatusPermanentRedirect,
		}))
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Pre(middleware.Secure())
	e.Use(middleware.Gzip())
	e.Use(middleware.CORS())

	// Serve the static web content on the base echo instance
	e.Static("*", "./static")

	// Bind all API endpoint handlers
	e.GET("/lookup/:ip", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ic.Lookup(c.Param("ip")))
	})
	e.GET("/healthcheck", func(c echo.Context) error {
		return c.NoContent(http.StatusOK)
	})

	// Listen on the passed port
	e.Logger.Fatal(e.Start(":" + port))
}

// getenv attempts to retrieve and return a variable from the environment. If it
// fails it will either crash or failover to a passed default value
func getenv(key string, def ...string) string {
	if v, ok := os.LookupEnv(key); ok {
		return v
	}
	if len(def) == 0 {
		log.Fatalf("%s not defined in environment", key)
	}
	return def[0]
}
