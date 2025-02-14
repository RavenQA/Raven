package db

import (
	"context"
	"database/sql"
	"path/filepath"
	"time"

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

func (d *Db) DropAll(ctx context.Context) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, stmtDropAll)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (d *Db) UpdateInstallPath(ctx context.Context, installPath, product, version string) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	_, err = tx.ExecContext(ctx, stmtUpdateBrowserInstallPath, installPath, product, version)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (d *Db) InsertBrowsers(ctx context.Context, bs []browser.Browser) error {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return err
	}
	defer tx.Rollback()
	var args []any
	for _, b := range bs {
		args = append(args, b.Product, b.Version, b.ReleaseDate.UTC().Format(time.RFC3339), b.InstallPath)
	}
	_, err = tx.ExecContext(ctx, stmtInsertBrowsers(len(bs)), args...)
	if err != nil {
		return err
	}
	return tx.Commit()
}

func (d *Db) GetBrowser(ctx context.Context, product, version string) (*browser.Browser, error) {
	tx, err := d.db.BeginTx(ctx, nil)
	if err != nil {
		return nil, err
	}
	defer tx.Rollback()
	rows, err := tx.QueryContext(ctx, stmtGetBrowser, product, version)
	if err != nil {
		return nil, err
	}
	var b browser.Browser
	for rows.Next() {
		var releaseDate string
		err = rows.Scan(&b.Product, b.Version, &releaseDate, &b.InstallPath)
		if err != nil {
			return nil, err
		}
		d, err := time.Parse(time.RFC3339, releaseDate)
		if err != nil {
			return nil, err
		}
		b.ReleaseDate = d.UTC()
		break
	}
	return &b, nil
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
		var releaseDate string
		err = rows.Scan(&b.Product, &b.Version, &releaseDate, &b.InstallPath)
		if err != nil {
			return nil, err
		}
		d, err := time.Parse(time.RFC3339, releaseDate)
		if err != nil {
			return nil, err
		}
		b.ReleaseDate = d.UTC()
		bs = append(bs, b)
	}
	return bs, tx.Commit()
}
