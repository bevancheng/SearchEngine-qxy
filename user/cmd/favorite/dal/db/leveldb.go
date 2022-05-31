package db

import (
	"github.com/syndtr/goleveldb/leveldb"
	"github.com/syndtr/goleveldb/leveldb/filter"
	"github.com/syndtr/goleveldb/leveldb/opt"
)

type LeveldbStorage struct {
	db   *leveldb.DB
	path string
}

func Open(path string) (*LeveldbStorage, error) {

	o := &opt.Options{
		Filter: filter.NewBloomFilter(10),
	}

	db, err := leveldb.OpenFile(path, o)
	if err != nil {
		return nil, err
	}
	return &LeveldbStorage{
		db:   db,
		path: path,
	}, nil
}

func (s *LeveldbStorage) Get(key []byte) ([]byte, bool) {
	buf, err := s.db.Get(key, nil)
	if err != nil {
		panic(err)
	}
	return buf, true
}

func (s *LeveldbStorage) Has(key []byte) bool {
	has, err := s.db.Has(key, nil)
	if err != nil {
		panic(err)
	}
	return has
}

func (s *LeveldbStorage) Set(key []byte, value []byte) error {
	err := s.db.Put(key, value, nil)
	if err != nil {
		panic(err)
	}
	return nil
}

func (s *LeveldbStorage) Delete(key []byte) error {
	return s.db.Delete(key, nil)
}

func (s *LeveldbStorage) Close() error {
	return s.db.Close()
}

func (s *LeveldbStorage) Count() int64 {
	var count int64
	iter := s.db.NewIterator(nil, nil)
	for iter.Next() {
		count++
	}
	iter.Release()
	return count
}
