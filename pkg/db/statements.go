package db

const stmtSchema = `
	CREATE TABLE IF NOT EXISTS browsers (
		name TEXT NOT NULL,
		version_id INTEGER NOT NULL,
		path TEXT NOT NULL,
		PRIMARY KEY (name, version_id)
	);
	
	CREATE TABLE IF NOT EXISTS browser_versions (
		id INTEGER PRIMARY KEY,
		major INTEGER NOT NULL,
		minor INTEGER,
		build INTEGER,
		patch INTEGER,
		release_date TEXT NOT NULL
	)
`

// SELECT into a browser.Browser.
const stmtGetBrowsers = `
	SELECT b.name, v.major, v.minor, v.build, v.patch, v.release_date, b.path,
	FROM browsers b
	INNER JOIN browser_versions v
	ON b.version_id = v.id;
`

// INSERT major, minor, build, patch, release_date
// RETURNING id.
const stmtInsertVersionPreparedAndGetId = `
	INSERT INTO browser_versions (major, minor, build, patch, release_date)
	VALUES (?, ?, ?, ?, ?)
	RETURNING id;
`

// INSERT name, version_id, path, available.
const stmtInsertBrowserPrepared = `
	INSERT INTO browsers (name, version_id, path)
	VALUES (?, ?, ?);
`

// TODO: When do we need to be able to update browsers?
