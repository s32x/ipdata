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
	e.Pre(middleware.RemoveTrailingSlashWithConfig(
		middleware.TrailingSlashConfig{
			Skipper:      middleware.DefaultSkipper,
			RedirectCode: http.StatusPermanentRedirect,
		}))
	e.Pre(middleware.Secure())

	// Configure HTTP redirects and serve the web index if being hosted in prod
	if strings.Contains(strings.ToLower(env), "prod") {
		e.Pre(middleware.HTTPSNonWWWRedirect())
	}

	// Bind remaining middleware
	e.Use(middleware.Logger())
	e.Use(middleware.Recover())
	e.Use(middleware.Gzip())

	// Serve the static web content on the base echo instance
	wb := packr.New("web box", "./web")
	e.GET("*", echo.WrapHandler(http.FileServer(wb)))

	// Create the API group with separate middlewares
	api := e.Group("/lookup")
	api.Use(middleware.CORS())

	// Bind all API endpoint handlers
	api.GET("/:ip", func(c echo.Context) error {
		return c.JSON(http.StatusOK, ic.Lookup(c.Param("ip")))
	})

	// Listen on the passed port
	e.Logger.Fatal(e.Start(":" + port))
}
