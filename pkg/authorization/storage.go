package authorization

import (
	"context"
	"encoding/json"
	"math/rand"
	"sync"
	"time"
)

type Storage interface {
	Get(ctx context.Context, key string) ([]byte, error)
	Set(ctx context.Context, key string, val []byte, expiration ...time.Duration) error
	Del(ctx context.Context, keys ...string) error
}

type memoryStorage struct {
	sync.Map
}

func (s *memoryStorage) Get(ctx context.Context, key string) ([]byte, error) {
	val, ok := s.Load(key)
	if !ok {
		return nil, ErrNotFound
	}
	return val.([]byte), nil
}
func (s *memoryStorage) Set(ctx context.Context, key string, val []byte, expiration ...time.Duration) error {
	s.Store(key, val)
	return nil
}

func (s *memoryStorage) Del(ctx context.Context, keys ...string) error {
	for _, key := range keys {
		s.Delete(key)
	}
	return nil
}

type secretItem struct {
	UID    string `json:"uid"`
	CID    string `json:"cid"`
	Secret string `json:"secret"`
}

type secretStorage struct {
	storage Storage
	max     int
	maxAge  time.Duration
}

func (s *secretStorage) get(ctx context.Context, uid string, cid string) ([]*secretItem, *secretItem, error) {
	data, err := s.storage.Get(ctx, uid)
	if err != nil {
		return nil, nil, err
	}
	if len(data) == 0 {
		return nil, nil, ErrNotFound
	}
	var items []*secretItem
	if err := json.Unmarshal(data, &items); err != nil {
		return nil, nil, err
	}
	for _, item := range items {
		if item.CID == cid {
			return items, item, nil
		}
	}
	return items, nil, ErrNotFound
}

func (s *secretStorage) Get(ctx context.Context, uid string, cid string) (*secretItem, error) {
	_, item, err := s.get(ctx, uid, cid)
	return item, err
}

func (s *secretStorage) Put(ctx context.Context, uid string, cid string) (*secretItem, error) {
	items, item, err := s.get(ctx, uid, cid)
	if err != nil && err != ErrNotFound {
		return nil, err
	}
	if item != nil {
		return item, nil
	}
	item = &secretItem{
		UID:    uid,
		CID:    cid,
		Secret: randString(16),
	}
	items = append(items, item)
	if len(items) > s.max {
		items = items[1:]
	}
	data, err := json.Marshal(item)
	if err != nil {
		return nil, err
	}
	if err := s.storage.Set(ctx, uid, data, s.maxAge); err != nil {
		return nil, err
	}
	return item, nil
}

func newSecretStorage(storage Storage, max int, maxAge time.Duration) *secretStorage {
	return &secretStorage{
		storage: storage,
		max:     max,
		maxAge:  maxAge,
	}
}

func randString(n int) string {
	const seed = "abcdefghijklmnopqrstuvwxyz0123456789"
	b := make([]byte, n)
	for i := range b {
		b[i] = seed[rand.Intn(len(seed))]
	}
	return string(b)
}
