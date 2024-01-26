package task

import (
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type Repository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) *Repository {
	return &Repository{
		db: db,
	}
}

func (r *Repository) List() (Tasks, error) {
	tasks := make([]*Task, 0)
	if err := r.db.Find(&tasks).Error; err != nil {
		return nil, err
	}
	return tasks, nil
}

func (r *Repository) Create(task *Task) (*Task, error) {
	if err := r.db.Create(task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *Repository) Read(id uuid.UUID) (*Task, error) {
	task := &Task{}
	if err := r.db.Where("id = ?", id).First(&task).Error; err != nil {
		return nil, err
	}
	return task, nil
}

func (r *Repository) Update(task *Task) (int64, error) {
	result := r.db.Model(&Task{}).
		Select("Title", "Description", "Created_date", "Completed_date", "Deleted_date").
		Where("id = ?", task.ID).
		Updates(task)
	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Task{})
	return result.RowsAffected, result.Error
}
