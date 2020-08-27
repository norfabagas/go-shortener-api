package main

import (
	"io/ioutil"
	"log"
	"os"
	"testing"
)

var a App

func TestMain(m *testing.M) {
	a = App{}
	os.Setenv("APP_ENV", "test")
	a.Initialize()

	ensureTableExists()

	code := m.Run()

	clearTable()

	os.Exit(code)
}

func ensureTableExists() {
	if _, err := a.DB.Exec(tableCreationQuery()); err != nil {
		log.Fatal(err)
	}
}

func clearTable() {
	a.DB.Exec("DELETE FROM urls")
	a.DB.Exec("ALTER TABLE urls AUTO_INCREMENT = 1")
}

func tableCreationQuery() string {
	content, err := ioutil.ReadFile("db/migrations/schema.txt")
	if err != nil {
		log.Fatal(err)
	}
	return string(content)
}
