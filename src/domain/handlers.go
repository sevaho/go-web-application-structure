package domain

import (
	"github.com/gofiber/contrib/websocket"
	"github.com/gofiber/fiber/v2"
	"github.com/gofrs/uuid"
	"github.com/sevaho/gowas/src/logger"
	"github.com/sevaho/gowas/src/models"
)

func (domain *Domain) HandleGetTasks(ctx *fiber.Ctx) error {
	tasks := domain.tasks.GetAll()

	return ctx.Status(200).Render("pages/tasks", fiber.Map{
		"tasks": tasks,
	})
}

func (domain *Domain) HandleDeleteTask(ctx *fiber.Ctx) error {
	task_id, err := uuid.FromString(ctx.Params("id"))
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Unable to parse id to uuid")
		return ctx.Status(400).Render("pages/tasks", fiber.Map{"errors": []string{err.Error()}})

	}
	if err := domain.tasks.Delete(task_id); err != nil {
		return ctx.Status(400).Render("pages/tasks", fiber.Map{"errors": []string{err.Error()}})
	}

	return ctx.Status(200).SendString("<p>Deleted!</p>")
}

func (domain *Domain) HandlePatchTask(ctx *fiber.Ctx) error {
	task_id, err := uuid.FromString(ctx.Params("id"))
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Unable to parse id to uuid")
		return ctx.Status(400).Render("pages/tasks", fiber.Map{"errors": []string{err.Error()}})
	}

	task := models.TaskCreateCmd{Title: ctx.FormValue("title"), Text: ctx.FormValue("text")}

	if err := domain.tasks.Update(task_id, task); err != nil {
		return ctx.Status(400).Render("pages/tasks", fiber.Map{"errors": []string{err.Error()}})
	}

	return ctx.Status(200).SendString("<p>Updated!</p>")
}

func (domain *Domain) HandlePostTask(ctx *fiber.Ctx) error {
	task := models.TaskCreateCmd{Title: ctx.FormValue("title"), Text: ctx.FormValue("text")}

	if _, err := domain.tasks.Store(task); err != nil {
		return ctx.Status(400).Render("pages/tasks", fiber.Map{"errors": []string{err.Error()}})
	}

	return ctx.Status(201).SendString("<p>Created!</p>")
}

func (domain *Domain) HandleWebsocket(ctx *websocket.Conn) {
}
