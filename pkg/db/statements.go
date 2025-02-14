package db

import (
	"fmt"
	"strings"
)

const stmtSchema = `
	CREATE TABLE IF NOT EXISTS browsers (
		product 		TEXT NOT NULL,
		version 		TEXT NOT NULL,
		release_date 	DATETIME NOT NULL,
		install_path 	TEXT NOT NULL,
		PRIMARY KEY (product, version)
	);
`

const stmtDropAll = `
	DROP TABLE browsers;
`

// SELECT all browsers.
// Scans into product string, version string, release_date string, install_path string
const stmtGetBrowsers = `
	SELECT product, version, release_date, install_path
	FROM browsers;
`

// SELECT a single browser into a browser.Browser.
// // Scans into product string, version string, release_date string, install_path string
// REQUIRES NAMED PARAMS: browsers.product, browsers.version
const stmtGetBrowser = `
	SELECT b.product, b.version, b.release_date, b.install_path
	FROM browsers b
	WHERE b.product = ?
	AND b.version = ?;
`

// REQUIRES NAMED PARAMS: install_path, browsers.product, browsers.version
const stmtUpdateBrowserInstallPath = `
	UPDATE browsers SET install_path = ?
	WHERE product = ?
	AND version = ?;
`

// REQUIRES {rows} OF NAMED PARAMS: product, version, release_date, install_path.
func stmtInsertBrowsers(rows int) string {
	return fmt.Sprintf(`
		INSERT INTO browsers (product, version, release_date, install_path)
		VALUES %s
		ON CONFLICT DO NOTHING;
		`,
		multiRowBindParams(4, rows),
	)
}

func multiRowBindParams(cols, rows int) string {
	row := "("
	row += strings.Repeat("?, ", cols-1)
	row += "?),"
	out := strings.Repeat(row, rows)
	return out[:len(out)-1] // Trim last comma
}
