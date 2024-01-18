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
	items := make([]*Task, 0)
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) Create(item *Task) (*Task, error) {
	if err := r.db.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *Repository) Read(id uuid.UUID) (*Task, error) {
	item := &Task{}
	if err := r.db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *Repository) Update(item *Task) (int64, error) {
	result := r.db.Model(&Task{}).
		Select("Title", "Description", "Created_date", "Completed_date", "Deleted_date").
		Where("id = ?", item.ID).
		Updates(item)
	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Task{})
	return result.RowsAffected, result.Error
}
