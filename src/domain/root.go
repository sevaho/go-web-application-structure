package domain

import (
	"github.com/sevaho/gowas/src/db"
	"github.com/sevaho/gowas/src/repositories"
)

type Domain struct {
	tasks *repositories.TaskRepository
}

func New() *Domain {
	db := db.SetUpDB()

	return &Domain{
		tasks: repositories.NewTaskRepository(db),
	}
}
