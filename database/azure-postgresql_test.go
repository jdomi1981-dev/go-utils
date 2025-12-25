package database_test

import (
	"database/sql"
	"testing"

	"github.com/jdomi1981-dev/go-utils/database"
)

func TestConnectDB(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		dbHost       string
		dbUser       string
		dbPassword   string
		dbName       string
		resourceName string
		want         *sql.DB
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := database.ConnectDB(tt.dbHost, tt.dbUser, tt.dbPassword, tt.dbName, tt.resourceName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ConnectDB() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ConnectDB() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ConnectDB() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecuteQuery(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db           *sql.DB
		sqlStatement string
		resourceName string
		want         *sql.Rows
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := database.ExecuteQuery(tt.db, tt.sqlStatement, tt.resourceName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ExecuteQuery() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ExecuteQuery() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ExecuteQuery() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExecuteTransaction(t *testing.T) {
	tests := []struct {
		name string // description of this test case
		// Named input parameters for target function.
		db           *sql.DB
		sqlStatement string
		resourceName string
		want         int64
		wantErr      bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, gotErr := database.ExecuteTransaction(tt.db, tt.sqlStatement, tt.resourceName)
			if gotErr != nil {
				if !tt.wantErr {
					t.Errorf("ExecuteTransaction() failed: %v", gotErr)
				}
				return
			}
			if tt.wantErr {
				t.Fatal("ExecuteTransaction() succeeded unexpectedly")
			}
			// TODO: update the condition below to compare got with tt.want.
			if true {
				t.Errorf("ExecuteTransaction() = %v, want %v", got, tt.want)
			}
		})
	}
}
