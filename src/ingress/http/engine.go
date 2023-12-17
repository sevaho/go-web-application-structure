package http

import (
	"embed"
	"net/http"
	"time"

	"github.com/gofiber/template/django/v3"
	"github.com/sevaho/gowas/src/config"
)

func createEngine(efs *embed.FS) *django.Engine {
	engine := django.NewPathForwardingFileSystem(http.FS(efs), "templates", ".html")
	if config.Config.DEVELOPMENT {
		engine.Reload(true)
	}

	engine.AddFunc("_prettytime", func(t time.Time) string {
		return t.Format(time.DateTime)
	})

	return engine
}
