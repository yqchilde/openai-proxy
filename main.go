package main

import (
	"fmt"
	"os"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/proxy"
)

func main() {
	app := fiber.New()
	app.All("/*", func(ctx *fiber.Ctx) error {
		proxyUrl := fmt.Sprintf("https://%s/%s", os.Getenv("PROXY_DOMAIN"), ctx.Params("*", ""))
		if err := proxy.Do(ctx, proxyUrl); err != nil {
			return err
		}
		ctx.Response().Header.Del(fiber.HeaderServer)
		return nil
	})
	if err := app.Listen(":5333"); err != nil {
		fmt.Printf("启动失败: %s", err.Error())
		os.Exit(1)
	}
}
