package handler

import (
	"fmt"
	"net/http"
	"os"

	"github.com/gofiber/adaptor/v2"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

// Handler is the main entry point of the application. Think of it like the main() method
func Handler(w http.ResponseWriter, r *http.Request) {
	r.RequestURI = r.URL.String()
	handler().ServeHTTP(w, r)
}

// building the fiber application
func handler() http.HandlerFunc {
	app := fiber.New()
	app.All("/*", func(ctx *fiber.Ctx) error {
		proxyUrl := fmt.Sprintf("https://%s/%s", os.Getenv("PROXY_DOMAIN"), ctx.Params("*", ""))
		if err := proxy.Do(ctx, proxyUrl); err != nil {
			return err
		}
		ctx.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	return adaptor.FiberApp(app)
}
