package db

import (
	"errors"
	"myserver/models"
	"strings"
)

var (
	ErrKeyNotFound   = errors.New("key not found in json request")
	ErrValueNotFound = errors.New("value not found in json request")
	ErrValueNotInDb  = errors.New("value not found in database with provided key")
)

type Db struct {
	storage map[string]string
	users   map[string]string
}

func NewDb() *Db {
	return &Db{
		storage: make(map[string]string),
		users:   make(map[string]string),
	}
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
