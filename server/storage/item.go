package storage

import (
	"container/list"
	"time"
)

type Item struct {
	Value      interface{}
	ExpireTime time.Time
}

func NewItem(value interface{}, ttl uint64) *Item {
	return &Item{
		Value:      value,
		ExpireTime: getExpireTime(ttl),
	}
}

func (i *Item) CastString() (string, error) {
	if value, ok := i.Value.(string); ok {
		return value, nil
	} else {
		return "", KeyStringTypeError
	}
}

func (i *Item) CastHash() (Hash, error) {
	if hash, ok := i.Value.(Hash); ok {
		return hash, nil
	} else {
		return nil, KeyHashTypeError
	}
}

func (i *Item) CastList() (*list.List, error) {
	if list, ok := i.Value.(*list.List); ok {
		return list, nil
	} else {
		return nil, KeyListTypeError
	}
}

func (i *Item) IsAlive() bool {
	return i.ExpireTime.IsZero() || i.ExpireTime.After(time.Now())
}

func (i *Item) SetTTL(ttl uint64) {
	i.ExpireTime = getExpireTime(ttl)
}

func getExpireTime(ttl uint64) (expireTime time.Time) {
	if ttl > 0 {
		expireTime = time.Now().Add(time.Duration(ttl) * time.Second)
	}
	return
}
