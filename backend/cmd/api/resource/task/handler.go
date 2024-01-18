package task

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/go-chi/chi/v5"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type API struct {
	repository *Repository
}

func New(db *gorm.DB) *API {
	return &API{
		repository: NewRepository(db),
	}
}

// List godoc
//
// @summary List items
// @description List items
// @tags items
// @accept json
// @produce json
// @success 200 {array} DTO
// @failure 500 {object} err.Error
// @router /items [get]
func (a *API) List(w http.ResponseWriter, r *http.Request) {
	tasks, err := a.repository.List()
	if err != nil {
		// handle later
		return
	}

	if len(tasks) == 0 {
		fmt.Fprint(w, "[]")
		return
	}

	if err := json.NewEncoder(w).Encode(tasks.toDTO()); err != nil {
		// handle later
		return
	}
}

// Create godoc
//
// @summary Create item
// @description Create item
// @tags item
// @accept json
// @produce json
// @param body body Form true "Item form"
// @success 201
// @failure 400 {object} err.Error
// @failure 422 {object} err.Errors
// @failure 500 {object} err.Error
// @router /items [post]
func (a *API) Create(w http.ResponseWriter, r *http.Request) {
	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// handle later
		return
	}

	newTask := form.ToModel()
	newTask.ID = uuid.New()

	_, err := a.repository.Create(newTask)
	if err != nil {
		// handle later
		return
	}

	w.WriteHeader(http.StatusCreated)
}

// Read godoc
//
// @summary Read item
// @description Read item
// @tags items
// @accept json
// @produce json
// @param id path string true "Item ID"
// @success 200 {object} DTO
// @failure 400 {object} err.Error
// @failure 404
// @failure 500 {object} err.Error
// @router /items/{id} [get]
func (a *API) Read(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// handle later
		return
	}

	task, err := a.repository.Read(id)
	if err != nil {
		if err == gorm.ErrRecordNotFound {
			w.WriteHeader(http.StatusNotFound)
			return
		}
		// handle later
		return
	}

	dto := task.ToDto()
	if err := json.NewEncoder(w).Encode(dto); err != nil {
		// handle later
		return
	}
}

// Update godoc
//
// @summary Delete item
// @description Delete item
// @tags items
// @accept json
// @produce json
// @param id path string true "Item ID"
// @param body body Form true "Item form"
// @success 200
// @failure 400 {object} err.Error
// @failure 404
// @failure 422 {object} err.Errors
// @failure 500 {object} err.Error
// @router /items/{id} [put]
func (a *API) Update(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// handle later
		return
	}

	form := &Form{}
	if err := json.NewDecoder(r.Body).Decode(form); err != nil {
		// handle later
		return
	}

	task := form.ToModel()
	task.ID = id

	rows, err := a.repository.Update(task)
	if err != nil {
		// handle later
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

// Delete godoc
//
// @summary Delete item
// @description
// @tags items
// @accept json
// @produce json
// @param id path string true "Item ID"
// @success 200
// @failure 400 {object} err.Error
// @failure 404
// @failure 500 {object} err.Error
// @router /items/{id} [delete]
func (a *API) Delete(w http.ResponseWriter, r *http.Request) {
	id, err := uuid.Parse(chi.URLParam(r, "id"))
	if err != nil {
		// handle later
		return
	}

	rows, err := a.repository.Delete(id)
	if err != nil {
		// handle later
		return
	}
	if rows == 0 {
		w.WriteHeader(http.StatusNotFound)
		return
	}
}

func (f *Form) ToModel() *Task {
	return &Task{
		Title:       f.Title,
		Description: f.Description,
		CreatedDate: time.Now(),
	}
}

func (t *Task) ToDto() *DTO {
	return &DTO{
		ID:            t.ID.String(),
		Title:         t.Title,
		Description:   t.Description,
		CreatedDate:   t.CreatedDate.Format("2006-01-02"),
		CompletedDate: t.CompletedDate.Format("2006-01-02"),
		DeletedDate:   t.DeletedDate.Format("2006-01-02"),
	}
}

func (ts Tasks) toDTO() []*DTO {
	dtos := make([]*DTO, len(ts))
	for i, v := range ts {
		dtos[i] = v.ToDto()
	}
	return dtos
}
