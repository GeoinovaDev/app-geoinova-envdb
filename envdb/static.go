package envdb

import "gorm.io/gorm"

var current *EnvDB

func create() *EnvDB {
	if current == nil {
		current = NewEnvDB()
	}

	return current
}

func Add(name string, db *gorm.DB) *EnvDB {
	e := create()
	e.m[name] = db
	return e
}

func Get(name string) *gorm.DB {
	e := create()

	if db, ok := e.m[name]; ok {
		return db
	}

	if db, ok := e.m["default"]; ok {
		return db
	}

	return nil
}
