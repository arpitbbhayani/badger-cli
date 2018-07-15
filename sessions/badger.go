package sessions

import (
	"github.com/dgraph-io/badger"
)

var (
	// BadgerDBSessions holds the global BadgerDB sessions
	BadgerDBSessions = make(map[string]*badger.DB)
)

// InitBadgerDB initializes the BadgerDB session for a key if it is nil
func InitBadgerDB(sessionKey string, databaseDir string) {
	opts := badger.DefaultOptions
	opts.Dir = databaseDir
	opts.ValueDir = databaseDir

	db, isAlreadyInitialized := BadgerDBSessions[sessionKey]
	if !isAlreadyInitialized {
		var err error
		db, err = badger.Open(opts)
		if err != nil {
			panic(err)
		}

		BadgerDBSessions[sessionKey] = db
	}
}

// UninitBadgerDB uninitializes the BadgerDB session for a key if it is not nil
func UninitBadgerDB(sessionKey string) {
	db, isInitialized := BadgerDBSessions[sessionKey]
	if isInitialized && db != nil {
		db.Close()
		delete(BadgerDBSessions, sessionKey)
	}
}
