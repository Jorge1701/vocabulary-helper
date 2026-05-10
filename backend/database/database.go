package database

import (
	"database/sql"
	"fmt"
	"unicode"
	"vocabulary-helper/model"

	_ "modernc.org/sqlite"

	"golang.org/x/text/unicode/norm"
)

const (
	WIKDICT_URL = "https://www.wikdict.com/es-pt/"
)

type DatabaseSearch struct {
	Found   bool
	Source  string
	Meaning []model.Meaning
}

type DatabaseEntry struct {
	Sense      string
	WrittenRep string
	TransList  string
}

func FindInDatabase(word string) DatabaseSearch {
	db, err := createDB()
	if err != nil {
		fmt.Println("Error opening connection to database:", err)
		return DatabaseSearch{Found: false}
	}
	defer db.Close()

	rows, err := db.Query(`
		SELECT sense, written_rep, trans_list
		FROM translation
		WHERE score > 1
		AND sense is not null
		AND trans_list is not null
		AND writte_rep_norm = ?
	`, word)
	if err != nil {
		fmt.Println("Error quering:", err)
		return DatabaseSearch{Found: false}
	}
	defer rows.Close()

	meanings := []model.Meaning{}
	for rows.Next() {
		var e DatabaseEntry
		if err := rows.Scan(&e.Sense, &e.WrittenRep, &e.TransList); err != nil {
			fmt.Println("Error parsing rows:", err)
			return DatabaseSearch{Found: false}
		}

		meanings = append(meanings, model.Meaning{
			Text:        e.Sense,
			Translation: e.TransList,
		})
	}

	if len(meanings) == 0 {
		return DatabaseSearch{Found: false}
	} else {
		return DatabaseSearch{Found: true, Meaning: meanings, Source: fmt.Sprint(WIKDICT_URL, word)}
	}
}

func createDB() (*sql.DB, error) {
	db, err := sql.Open("sqlite", "./databases/pt-es.sqlite3")
	if err != nil {
		return nil, err
	}
	return db, nil
}

func NormalizeDatabase() error {
	db, err := createDB()
	if err != nil {
		return fmt.Errorf("Could not create DB: %w", err)
	}
	defer db.Close()

	_, err = db.Exec(`ALTER TABLE translation ADD COLUMN writte_rep_norm TEXT`)
	if err != nil {
		fmt.Println("Column may already exist, continuing:", err)
	}

	rows, err := db.Query(`SELECT rowid, written_rep FROM translation`)
	if err != nil {
		return fmt.Errorf("select failed: %w", err)
	}
	defer rows.Close()

	type entry struct {
		rowid      int64
		writtenRep string
	}
	var entries []entry
	for rows.Next() {
		var e entry
		if err := rows.Scan(&e.rowid, &e.writtenRep); err != nil {
			return fmt.Errorf("scan failed: %w", err)
		}
		entries = append(entries, e)
	}
	if err := rows.Err(); err != nil {
		return err
	}

	tx, err := db.Begin()
	if err != nil {
		return fmt.Errorf("begin tx failed: %w", err)
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`UPDATE translation SET writte_rep_norm = ? WHERE rowid = ?`)
	if err != nil {
		return fmt.Errorf("prepare failed: %w", err)
	}
	defer stmt.Close()

	for _, e := range entries {
		if _, err := stmt.Exec(normalize(e.writtenRep), e.rowid); err != nil {
			return fmt.Errorf("update failed for rowid %d: %w", e.rowid, err)
		}
	}

	return tx.Commit()
}

func normalize(s string) string {
	t := norm.NFD.String(s)
	result := make([]rune, 0, len(t))
	for _, r := range t {
		if unicode.Is(unicode.Mn, r) {
			continue // strip diacritic
		}
		result = append(result, unicode.ToLower(r))
	}
	return string(result)
}
