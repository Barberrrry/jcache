package storage

import (
	"errors"
)

type Storage interface {
	Keys() []string
	Expire(key string, ttl uint64) error
	Get(key string) (string, error)
	Set(key, value string, ttl uint64) error
	Update(key, value string) error
	Delete(key string) error
	HashCreate(key string, ttl uint64) error
	HashGet(key, field string) (string, error)
	HashGetAll(key string) (map[string]string, error)
	HashSet(key, field, value string) error
	HashDelete(key, field string) error
	HashLen(key string) (int, error)
	HashKeys(key string) ([]string, error)
	ListCreate(key string, ttl uint64) error
	ListLeftPop(key string) (string, error)
	ListRightPop(key string) (string, error)
	ListLeftPush(key, value string) error
	ListRightPush(key, value string) error
	ListLen(key string) (int, error)
	ListRange(key string, start, stop int) ([]string, error)
}

var (
	KeyNotExistsError     = errors.New("Key does not exist")
	KeyAlreadyExistsError = errors.New("Key already exists")
	ListEmptyError        = errors.New("List is empty")
	FieldNotExistError    = errors.New("Field does not exist")
	KeyStringTypeError    = errors.New("Key type is not string")
	KeyHashTypeError      = errors.New("Key type is not hash")
	KeyListTypeError      = errors.New("Key type is not list")
)
