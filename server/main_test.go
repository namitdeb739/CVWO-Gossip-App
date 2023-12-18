package main_test
 
import (
     "os"
     "testing"
     "log"
     "net/http"
     "net/http/httptest"
     "strconv"
     "encoding/json"
     "bytes"
     "github.com/namitdeb739/CVWO-Gossip-App"
 )

var a main.App

func TestMain(m *testing.M) {
    a.Initialize(
        os.Getenv("APP_DB_USERNAME"),
        os.Getenv("APP_DB_PASSWORD"),
        os.Getenv("APP_DB_NAME"))

    ensureTableExists()
    code := m.Run()
    clearTable()
    os.Exit(code)
}

func ensureTableExists() {
    if _, err := a.DB.Exec(tableCreationQuery); err != nil {
        log.Fatal(err)
    }
}

func clearTable() {
    a.DB.Exec("DELETE FROM Users")
}

const tableCreationQuery = `CREATE TABLE IF NOT EXISTS gossip."Users"
(
    "User_ID" character varying(255) COLLATE pg_catalog."default" NOT NULL,
    CONSTRAINT PK_Users PRIMARY KEY ("User_ID")
);`