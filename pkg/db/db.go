package db

import (
	"context"
	"database/sql"
	"path/filepath"

	_ "github.com/mattn/go-sqlite3"
	"github.com/soikes/raven/pkg/browser"
)

type Db struct {
	db *sql.DB
}

const dbName = `raven.db`

func NewDB(path string) (*Db, error) {
	d, err := sql.Open("sqlite3", filepath.Join(path, dbName))
	if err != nil {
		return nil, err
	}
	return &Db{db: d}, nil
}

func (d *Db) Init(ctx context.Context) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, stmtSchema)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (d *Db) InsertBrowser(ctx context.Context, b browser.Browser) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	stmt, err := tx.Prepare(stmtInsertVersionPreparedAndGetId)
	if err != nil {
		return err
	}
	vid, err := stmt.QueryContext(ctx, b.Major, b.Minor, b.Build, b.Patch, b.ReleaseDate)
	stmt, err = tx.Prepare(stmtInsertBrowserPrepared)
	if err != nil {
		return err
	}
	_, err = stmt.ExecContext(ctx, b.Name, vid, b.Path, b.Available)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (d *Db) GetBrowsers(ctx context.Context) ([]browser.Browser, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.QueryContext(ctx, stmtGetBrowsers)
	if err != nil {
		return nil, err
	}
	var bs []browser.Browser
	for rows.Next() {
		var b browser.Browser
		err = rows.Scan(&b)
		if err != nil {
			return nil, err
		}
		bs = append(bs, b)
	}
	return bs, tx.Commit()
}
