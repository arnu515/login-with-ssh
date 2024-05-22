package main

import (
	"login-with-ssh/db"
)

func seed() (err error) {
	tx, err := db.DB.Begin()
	if err != nil {
		return err
	}
	defer func() {
		if err != nil {
			if rollbackError := tx.Rollback(); rollbackError != nil {
				println("Could not rollback: " + rollbackError.Error())
			}
			return
		}
	}()

	_, err = tx.Exec(`
CREATE TABLE users (
    id TEXT PRIMARY KEY,
    username TEXT NOT NULL UNIQUE,
    email TEXT NOT NULL UNIQUE
);
	`)
	if err != nil {
		return
	}
	_, err = tx.Exec(`
CREATE TABLE ssh_public_keys (
    user_id TEXT NOT NULL,
    key TEXT PRIMARY KEY,
    key_type TEXT NOT NULL,  -- ssh-rsa or ssh-ed25519
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);
	`)
	if err != nil {
		return
	}
	_, err = tx.Exec(`
CREATE TABLE ssh_queue (
	id INTEGER PRIMARY KEY AUTOINCREMENT,
    user_id TEXT NOT NULL,
    at TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE ON UPDATE CASCADE
);
	`)
	if err != nil {
		return
	}

	err = tx.Commit()
	if err != nil {
		return
	}
	return nil
}
