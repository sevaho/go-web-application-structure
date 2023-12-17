package repositories

import (
	"context"
	"time"

	"github.com/gofrs/uuid"
	"github.com/sevaho/gowas/src/db"
	"github.com/sevaho/gowas/src/logger"
	"github.com/sevaho/gowas/src/models"
)

type TaskRepository struct {
	db *db.Queries
}

func NewTaskRepository(db *db.Queries) *TaskRepository {
	return &TaskRepository{db: db}
}

func (r *TaskRepository) GetAll() []models.TaskQueryModel {
	results, err := r.db.SelectNotes(context.Background())
	if err != nil {
		logger.Logger.Error().Err(err).Msg("Unable to select all tasks")
		return nil
	}
	items := []models.TaskQueryModel{}
	for _, result := range results {
		items = append(
			items, models.TaskQueryModel{
				Title:     result.Title,
				Text:      result.Text,
				CreatedAt: result.CreatedAt.Time,
				ID:        db.PGUUIDToUUID(result.ID),
			},
		)
	}
	return items
}

func (r *TaskRepository) Store(data models.TaskCreateCmd) (int, error) {
	new_id, _ := uuid.NewV4()

	params := db.InsertNoteParams{
		ID:        db.UUIDToPGUUID(new_id),
		CreatedAt: db.TimeToTimestamp(time.Now().UTC()),
		Title:     data.Title,
		Text:      data.Text,
	}
	result, err := r.db.InsertNote(context.Background(), params)
	if err != nil {
		logger.Logger.Error().Err(err).Msgf("Unable to store note [request=%v]", params)
		return 0, err
	}
	return int(result), nil
}

func (r *TaskRepository) Update(id uuid.UUID, data models.TaskCreateCmd) error {
	params := db.UpdateNoteBySerialIDParams{
		NoteID:        db.UUIDToPGUUID(id),
		Title:         data.Title,
		Text:          data.Text,
		TitleDoUpdate: data.Title != "",
		TextDoUpdate:  data.Text != "",
	}
	_, err := r.db.UpdateNoteBySerialID(context.Background(), params)
	if err != nil {
		logger.Logger.Error().Err(err).Msgf("Unable to update note [request=%v]", params)
		return err
	}
	return nil
}

func (r *TaskRepository) Delete(id uuid.UUID) error {
	_, err := r.db.DeleteNote(context.Background(), db.UUIDToPGUUID(id))
	if err != nil {
		logger.Logger.Error().Err(err).Msgf("Unable to delete note [request=%v]", id)
		return err
	}
	return nil
}
