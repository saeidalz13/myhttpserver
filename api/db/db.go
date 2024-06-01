package db

import (
	"errors"
	"myserver/api/models"
	"strings"
)

var (
	ErrKeyNotFound   = errors.New("key not found in json request")
	ErrValueNotFound = errors.New("value not found in json request")
	ErrValueNotInDb  = errors.New("value not found in database with provided key")

	ErrNoEmail       = errors.New("email is not provided")
	ErrNoPassword    = errors.New("password is not provided")
	ErrEmailNotFound = errors.New("user with this email does not exist")
)

type passwordType []byte
type emailType string

type Db struct {
	storage map[string]string
	users   map[emailType]passwordType
}

func NewDb() *Db {
	return &Db{
		storage: make(map[string]string),
		users:   make(map[emailType]passwordType),
	}
}

func (d *Db) InsertUser(email string, password []byte) error {
	if len(password) == 0 {
		return ErrNoPassword
	}

	if email == "" {
		return ErrNoEmail
	}

	d.users[emailType(email)] = passwordType(password)
	return nil
}

func (d *Db) SelectUser(email string) ([]byte, error) {
	password, prs := d.users[emailType(email)]
	if !prs {
		return nil, ErrEmailNotFound
	}

	return password, nil
}

func (d *Db) DeleteItem(key string) error {
	// Delete is a no-op if the key does not exist
	// so I look for it just to see if it actually exists
	_, err := d.SelectItem(key)
	if err != nil {
		return err
	}

	delete(d.storage, key)
	return nil
}

func (d *Db) SelectItem(key string) (string, error) {
	value, prs := d.storage[key]
	if !prs {
		return "", ErrValueNotInDb
	}

	return value, nil
}

func (d *Db) InsertItem(item *models.Item) (map[string]string, error) {
	if strings.TrimSpace(item.Key) == "" {
		return nil, ErrKeyNotFound
	}

	if strings.TrimSpace(item.Value) == "" {
		return nil, ErrValueNotFound
	}

	d.storage[item.Key] = item.Value
	return map[string]string{item.Key: item.Value}, nil
}
