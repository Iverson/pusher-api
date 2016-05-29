package testfixtures

import (
	"database/sql"
	"fmt"
)

// MySQLHelper is the MySQL helper for this package
type MySQLHelper struct{}

func (MySQLHelper) paramType() int {
	return paramTypeQuestion
}

func (MySQLHelper) quoteKeyword(str string) string {
	return fmt.Sprintf("`%s`", str)
}

func (MySQLHelper) databaseName(db *sql.DB) (dbName string) {
	db.QueryRow("SELECT DATABASE()").Scan(&dbName)
	return
}

func (MySQLHelper) whileInsertOnTable(tx *sql.Tx, tableName string, fn func() error) error {
	return fn()
}

func (h *MySQLHelper) disableReferentialIntegrity(db *sql.DB, loadFn loadFunction) error {
	// re-enable after load
	defer db.Exec("SET FOREIGN_KEY_CHECKS = 1")

	tx, err := db.Begin()
	if err != nil {
		return err
	}

	_, err = tx.Exec("SET FOREIGN_KEY_CHECKS = 0")
	if err != nil {
		return err
	}

	err = loadFn(tx)
	if err != nil {
		tx.Rollback()
		return err
	}

	err = tx.Commit()
	return err
}
