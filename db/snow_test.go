package db

import (
	"regexp"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"
)

func TestSimulate(t *testing.T) {
	db, mock, err := sqlmock.New()
	if err != nil {
		t.Fatalf("an error '%s' was not expected when opening a stub database connection", err)
	}
	defer db.Close()
	rows := sqlmock.NewRows([]string{"result"}).AddRow("Statement executed")
	mock.ExpectQuery(regexp.QuoteMeta("SELECT system$wait(10)")).WillReturnRows(rows)

	// now we execute our method
	if _, err = simulate(db); err != nil {
		t.Errorf("error was not expected while waiting stats: %s", err)
	}

	// we make sure that all expectations were met
	if err := mock.ExpectationsWereMet(); err != nil {
		t.Errorf("there were unfulfilled expectations: %s", err)
	}
}
