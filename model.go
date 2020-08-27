package main

import (
	"database/sql"
	"errors"
	"time"
)

type URL struct {
	ID        int       `json:"id"`
	URL       string    `json:"url"`
	Code      string    `json:"code"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

func (u *URL) getUrl(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *URL) updateUrl(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *URL) deleteUrl(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *URL) createUrl(db *sql.DB) error {
	return errors.New("Not implemented")
}

func (u *URL) getUrls(db *sql.DB, start, count int) ([]URL, error) {
	return nil, errors.New("Not implemented")
}
