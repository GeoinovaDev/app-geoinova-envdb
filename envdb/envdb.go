package envdb

import "gorm.io/gorm"

type EnvDB struct {
	m map[string]*gorm.DB
}

func NewEnvDB() *EnvDB {
	e := &EnvDB{
		m: map[string]*gorm.DB{},
	}

	return e
}

func (e *EnvDB) Add(name string, db *gorm.DB) *EnvDB {
	e.m[name] = db
	return e
}

func (e *EnvDB) Get(name string) *gorm.DB {
	if db, ok := e.m[name]; ok {
		return db
	}

	if db, ok := e.m["default"]; ok {
		return db
	}

	return nil
}
