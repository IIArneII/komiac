package dal

import (
	"context"
	"komiac/internal/app/entities"
	"komiac/internal/app/errors"
	"komiac/tests"
	"testing"

	uuid "github.com/satori/go.uuid"
)

func TestCreate(t *testing.T) {
	t.Run("Create empty application", func(t *testing.T) {
		app := &entities.Application{}

		cxt := context.Background()
		_, err := tests.Storage.Create(cxt, app)
		if err != errors.BadUUID {
			t.Errorf("Not received error: %s", errors.BadUUID)
		}
	})

	t.Run("Create application", func(t *testing.T) {
		app := CreateUniqueApplication()

		cxt := context.Background()
		dbApp, err := tests.Storage.Create(cxt, app)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		err = СompareApplication(app, dbApp)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("Create existing application", func(t *testing.T) {
		app := CreateUniqueApplication()

		cxt := context.Background()
		_, err := tests.Storage.Create(cxt, app)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		_, err = tests.Storage.Create(cxt, app)
		if err == nil {
			t.Error("Not received error")
		}
	})
}

func TestGet(t *testing.T) {
	t.Run("Get existing application", func(t *testing.T) {
		app := CreateUniqueApplication()

		cxt := context.Background()
		_, err := tests.Storage.Create(cxt, app)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		dbApp, err := tests.Storage.Get(cxt, app.UUID)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		err = СompareApplication(app, dbApp)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("Get not existing application", func(t *testing.T) {
		cxt := context.Background()
		_, err := tests.Storage.Get(cxt, uuid.NewV4())
		if err != errors.ErrNotFound {
			t.Errorf("Not received error: %s", errors.ErrNotFound.Error())
		}
	})

	t.Run("Create application by empty uuid", func(t *testing.T) {
		cxt := context.Background()
		_, err := tests.Storage.Get(cxt, uuid.NullUUID{}.UUID)
		if err != errors.BadUUID {
			t.Errorf("Not received error: %s", errors.BadUUID)
		}
	})
}

func TestUpdate(t *testing.T) {
	t.Run("Update existing application", func(t *testing.T) {
		app := CreateUniqueApplication()

		cxt := context.Background()
		_, err := tests.Storage.Create(cxt, app)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		app.Year = app.Year + 1

		dbApp, err := tests.Storage.Update(cxt, app)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}

		err = СompareApplication(app, dbApp)
		if err != nil {
			t.Errorf("Error: %s", err.Error())
		}
	})

	t.Run("Update not existing application", func(t *testing.T) {
		app := CreateUniqueApplication()

		cxt := context.Background()
		_, err := tests.Storage.Update(cxt, app)
		if err != errors.ErrNotFound {
			t.Errorf("Not received error: %s", errors.ErrNotFound.Error())
		}
	})

	t.Run("Update existing application with empty uuid", func(t *testing.T) {
		app := CreateUniqueApplication()
		app.UUID = uuid.NullUUID{}.UUID

		cxt := context.Background()
		_, err := tests.Storage.Update(cxt, app)
		if err != errors.BadUUID {
			t.Errorf("Not received error: %s", errors.BadUUID)
		}
	})
}
