package http

import (
	"fmt"
	"time"

	"net/http"

	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/compress"
	"github.com/gofiber/fiber/v2/middleware/filesystem"
	"github.com/gofiber/fiber/v2/middleware/helmet"
	"github.com/gofiber/fiber/v2/middleware/recover"

	"github.com/sevaho/gowas/src/assets"
	"github.com/sevaho/gowas/src/config"
	"github.com/sevaho/gowas/src/domain"
	"github.com/sevaho/gowas/src/logger"
	"github.com/sevaho/gowas/src/util"

	"github.com/gofiber/contrib/websocket"
)

type HttpIngress struct {
	port   int
	Server *fiber.App
}

func New(domain *domain.Domain) *HttpIngress {
	app := HttpIngress{
		Server: fiber.New(fiber.Config{
			Views:                 createEngine(&assets.Assets),
			DisableStartupMessage: true,
			IdleTimeout:           5 * time.Second,
			ProxyHeader:           "X-Forwarded-For",
			ServerHeader:          "APP",
			AppName:               "APP",
			ReadBufferSize:        10000,
			PassLocalsToViews:     true,
			ErrorHandler: func(c *fiber.Ctx, err error) error {
				c.Status(fiber.StatusInternalServerError)
				return c.SendString("500 INTERNAL SERVER ERROR " + err.Error())
			},
		}),
		port: config.Config.HTTP_SERVER_PORT,
	}

	// Middleware
	app.Server.Use("/public", filesystem.New(filesystem.Config{
		Root:       http.FS(&assets.Assets),
		PathPrefix: "public",
		Browse:     true,
	}))
	app.Server.Use(helmet.New())
	app.Server.Use(recover.New())
	app.Server.Use(compress.New())
	app.Server.Use("/ws", func(c *fiber.Ctx) error {
		// IsWebSocketUpgrade returns true if the client
		// requested upgrade to the WebSocket protocol.
		if websocket.IsWebSocketUpgrade(c) {
			c.Locals("allowed", true)
			return c.Next()
		}
		return fiber.ErrUpgradeRequired
	})

	// Routes
	setupRoutes(app.Server, domain)

	// 404 not found
	app.Server.Use(func(c *fiber.Ctx) error {
		return c.Status(fiber.StatusNotFound).Render("errors/404", fiber.Map{
			"detail": util.PrettyJson(fmt.Sprintf(`{"detail": "%s Not Found!"}`, c.Request().URI())),
		})
	})

	return &app
}

func (i *HttpIngress) Serve() {
	go func() {
		logger.Logger.Info().Msgf("[HTTP SERVER] Running on http://localhost:%d", i.port)
		err := i.Server.Listen(fmt.Sprint(":", i.port))
		if err != nil {
			logger.Logger.Error().Err(err).Stack().Msgf("[HTTP SERVER] closed unexpectedly: %s", err)
		}
	}()
}

func (i *HttpIngress) ShutDown() {
	defer logger.Logger.Info().Msg("[HTTP SERVER] shut down.")
	err := i.Server.Shutdown()
	if err != nil {
		logger.Logger.Error().Err(err).Stack().Msgf("[HTTP SERVER] Error shutting down the server: %s", err)
	}
}
