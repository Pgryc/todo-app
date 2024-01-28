package task_test

import (
	"testing"
	"time"

	"todo-backend/cmd/api/resource/task"
	mockDB "todo-backend/mock/db"
	testUtil "todo-backend/util/test"

	"github.com/DATA-DOG/go-sqlmock"
	"github.com/google/uuid"
)

func TestRepository_List(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := task.NewRepository(db)

	mockRows := sqlmock.NewRows([]string{"id", "title", "author"}).
		AddRow(uuid.New(), "Task1", "Description1").
		AddRow(uuid.New(), "Task2", "Description2")

	mock.ExpectQuery("^SELECT (.+) FROM \"tasks\"").WillReturnRows(mockRows)

	tasks, err := repo.List()
	testUtil.NoError(t, err)
	testUtil.Equal(t, len(tasks), 2)
}

func TestRepository_Create(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := task.NewRepository(db)

	id := uuid.New()
	mock.ExpectBegin()
	mock.ExpectExec("^INSERT INTO \"tasks\"").
		WithArgs(id, "Title", "Description", mockDB.AnyTime{}, mockDB.AnyTime{}, mockDB.AnyTime{}).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	task := &task.Task{ID: id, Title: "Title", Description: "Description", CreatedDate: time.Now()}
	_, err = repo.Create(task)
	testUtil.NoError(t, err)
}

func TestRepository_Update(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := task.NewRepository(db)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(id, "Task1", "Description1")

	mock.ExpectBegin()
	mock.ExpectExec("^UPDATE \"tasks\" SET").
		WithArgs("Title", "Description", id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	task := &task.Task{ID: id, Title: "Title", Description: "Description"}
	rows, err := repo.Update(task)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)
}

func TestRepository_Delete(t *testing.T) {
	t.Parallel()

	db, mock, err := mockDB.NewMockDB()
	testUtil.NoError(t, err)

	repo := task.NewRepository(db)

	id := uuid.New()
	_ = sqlmock.NewRows([]string{"id", "title", "description"}).
		AddRow(id, "Task1", "Description1")

	mock.ExpectBegin()
	mock.ExpectExec("^DELETE FROM \"tasks\" WHERE").
		WithArgs(id).
		WillReturnResult(sqlmock.NewResult(1, 1))
	mock.ExpectCommit()

	rows, err := repo.Delete(id)
	testUtil.NoError(t, err)
	testUtil.Equal(t, 1, rows)
}
