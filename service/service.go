package service /* import "s32x.com/ipdata/service" */

import (
	"log"
	"net/http"
	"strings"

	packr "github.com/gobuffalo/packr/v2"
	"github.com/labstack/echo"
	"github.com/labstack/echo/middleware"
	"s32x.com/ipdata/ipdata"
)

// Start starts the ipdata API service using the passed params
func Start(port, env string) {
	// Create the ipdata client
	ic, err := ipdata.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	defer ic.Close()

	// Create a new echo Echo and bind all middleware
	e := echo.New()
	e.HideBanner = true
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.CORS())
	e.Use(middleware.Gzip())

	// Configure HTTP redirects and serve the web index if being hosted in prod
	if strings.Contains(strings.ToLower(env), "prod") {
		e.Pre(middleware.HTTPSNonWWWRedirect())

		// Serve the static web content
		wb := packr.New("web box", "./web")
		e.GET("*", echo.WrapHandler(http.FileServer(wb)))
	}

	// Bind all API endpoint handlers
	e.GET("/", func(c echo.Context) error {
		return c.Redirect(http.StatusTemporaryRedirect, "https://s32x.com/ipdata")
	})
	e.GET("/lookup/:ip", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ic.Lookup(c.Param("ip")))
	})

	// Listen on the passed port
	e.Logger.Fatal(e.Start(":" + port))
}
