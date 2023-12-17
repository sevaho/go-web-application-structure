package http

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofiber/fiber/v2/middleware/monitor"
	"github.com/sevaho/gowas/src/domain"
)

func setupRoutes(app *fiber.App, domain *domain.Domain) {
  // Redirects
	app.Get("/", func(c *fiber.Ctx) error { return c.Redirect("/tasks") })

	// PAGES
	app.Get("/tasks", domain.HandleGetTasks)

	// ACTIONS
	app.Delete("/tasks/:id", domain.HandleDeleteTask)
	app.Patch("/tasks/:id", domain.HandlePatchTask)
	app.Post("/tasks", domain.HandlePostTask)

	// SYSTEM ROUTES
	app.Get("/monitor", monitor.New())
	app.Get("/ws/:session_id", websocket.New(domain.HandleWebsocket, websocket.Config{RecoverHandler: RecoverWebsocketHandler}))
}
