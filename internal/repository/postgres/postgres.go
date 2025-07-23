package postgres

import (
	"context"
	"database/sql"
	"github.com/folivorra/vroom/application"
	"time"
)

func NewPostgres(ctx context.Context, app *application.App, dsn string) (*sql.DB, error) {
	timeout, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
	defer cancel()

	db, err := sql.Open("pgx", dsn)
	if err != nil {
		return nil, err
	}
	if err = db.PingContext(timeout); err != nil {
		return nil, err
	}

	app.RegisterCleanup(func(ctx context.Context) error {
		timeout, cancel := context.WithTimeout(ctx, 500*time.Millisecond)
		defer cancel()

		if err := db.PingContext(timeout); err != nil {
			return err
		} else {
			return db.Close()
		}
	})

	return db, nil
}
