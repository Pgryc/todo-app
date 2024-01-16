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

func (r *Repository) List() (Items, error) {
	items := make([]*Item, 0)
	if err := r.db.Find(&items).Error; err != nil {
		return nil, err
	}
	return items, nil
}

func (r *Repository) Create(item *Item) (*Item, error) {
	if err := r.db.Create(item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *Repository) Read(id uuid.UUID) (*Item, error) {
	item := &Item{}
	if err := r.db.Where("id = ?", id).First(&item).Error; err != nil {
		return nil, err
	}
	return item, nil
}

func (r *Repository) Update(item *Item) (int64, error) {
	result := r.db.Model(&Item{}).
		Select("Title", "Description", "Created_date", "Completed_date", "Deleted_date").
		Where("id = ?", item.ID).
		Updates(item)
	return result.RowsAffected, result.Error
}

func (r *Repository) Delete(id uuid.UUID) (int64, error) {
	result := r.db.Where("id = ?", id).Delete(&Item{})
	return result.RowsAffected, result.Error
}
