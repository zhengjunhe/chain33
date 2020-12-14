// Code generated by mockery v1.0.0. DO NOT EDIT.

package mocks

import (
	db "github.com/33cn/dplatformos/common/db"
	lru "github.com/hashicorp/golang-lru"

	mock "github.com/stretchr/testify/mock"
)

// DB is an autogenerated mock type for the DB type
type DB struct {
	mock.Mock
}

// Begin provides a mock function with given fields:
func (_m *DB) Begin() {
	_m.Called()
}

// BeginTx provides a mock function with given fields:
func (_m *DB) BeginTx() (db.TxKV, error) {
	ret := _m.Called()

	var r0 db.TxKV
	if rf, ok := ret.Get(0).(func() db.TxKV); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.TxKV)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func() error); ok {
		r1 = rf()
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// Close provides a mock function with given fields:
func (_m *DB) Close() {
	_m.Called()
}

// Commit provides a mock function with given fields:
func (_m *DB) Commit() error {
	ret := _m.Called()

	var r0 error
	if rf, ok := ret.Get(0).(func() error); ok {
		r0 = rf()
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// CompactRange provides a mock function with given fields: start, limit
func (_m *DB) CompactRange(start []byte, limit []byte) error {
	ret := _m.Called(start, limit)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(start, limit)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Delete provides a mock function with given fields: _a0
func (_m *DB) Delete(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// DeleteSync provides a mock function with given fields: _a0
func (_m *DB) DeleteSync(_a0 []byte) error {
	ret := _m.Called(_a0)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte) error); ok {
		r0 = rf(_a0)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Get provides a mock function with given fields: key
func (_m *DB) Get(key []byte) ([]byte, error) {
	ret := _m.Called(key)

	var r0 []byte
	if rf, ok := ret.Get(0).(func([]byte) []byte); ok {
		r0 = rf(key)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).([]byte)
		}
	}

	var r1 error
	if rf, ok := ret.Get(1).(func([]byte) error); ok {
		r1 = rf(key)
	} else {
		r1 = ret.Error(1)
	}

	return r0, r1
}

// GetCache provides a mock function with given fields:
func (_m *DB) GetCache() *lru.ARCCache {
	ret := _m.Called()

	var r0 *lru.ARCCache
	if rf, ok := ret.Get(0).(func() *lru.ARCCache); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(*lru.ARCCache)
		}
	}

	return r0
}

// Iterator provides a mock function with given fields: start, end, reserver
func (_m *DB) Iterator(start []byte, end []byte, reserver bool) db.Iterator {
	ret := _m.Called(start, end, reserver)

	var r0 db.Iterator
	if rf, ok := ret.Get(0).(func([]byte, []byte, bool) db.Iterator); ok {
		r0 = rf(start, end, reserver)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Iterator)
		}
	}

	return r0
}

// NewBatch provides a mock function with given fields: sync
func (_m *DB) NewBatch(sync bool) db.Batch {
	ret := _m.Called(sync)

	var r0 db.Batch
	if rf, ok := ret.Get(0).(func(bool) db.Batch); ok {
		r0 = rf(sync)
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(db.Batch)
		}
	}

	return r0
}

// Print provides a mock function with given fields:
func (_m *DB) Print() {
	_m.Called()
}

// Rollback provides a mock function with given fields:
func (_m *DB) Rollback() {
	_m.Called()
}

// Set provides a mock function with given fields: key, value
func (_m *DB) Set(key []byte, value []byte) error {
	ret := _m.Called(key, value)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(key, value)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// SetCacheSize provides a mock function with given fields: size
func (_m *DB) SetCacheSize(size int) {
	_m.Called(size)
}

// SetSync provides a mock function with given fields: _a0, _a1
func (_m *DB) SetSync(_a0 []byte, _a1 []byte) error {
	ret := _m.Called(_a0, _a1)

	var r0 error
	if rf, ok := ret.Get(0).(func([]byte, []byte) error); ok {
		r0 = rf(_a0, _a1)
	} else {
		r0 = ret.Error(0)
	}

	return r0
}

// Stats provides a mock function with given fields:
func (_m *DB) Stats() map[string]string {
	ret := _m.Called()

	var r0 map[string]string
	if rf, ok := ret.Get(0).(func() map[string]string); ok {
		r0 = rf()
	} else {
		if ret.Get(0) != nil {
			r0 = ret.Get(0).(map[string]string)
		}
	}

	return r0
}
