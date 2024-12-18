package echo

import (
	"net/http"
	"time"

	"github.com/labstack/echo/v4"
)

func cacheWithRevalidation(next echo.HandlerFunc) echo.HandlerFunc {
	return func(c echo.Context) error {
		// Call the next handler
		err := next(c)

		// Set Cache-Control header to enable caching and revalidation with a maximum age of 120 seconds
		c.Response().Header().Set("Cache-Control", "no-cache, max-age=120, must-revalidate")

		// Set Expires header to a future time to indicate the expiration time
		expiresTime := time.Now().Add(120 * time.Second)
		c.Response().Header().Set("Expires", expiresTime.UTC().Format(http.TimeFormat))

		// Set Last-Modified header to the current time
		c.Response().Header().Set("Last-Modified", time.Now().UTC().Format(http.TimeFormat))

		return err
	}
}
