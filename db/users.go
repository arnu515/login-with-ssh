package db

import (
	"context"
	"database/sql"

	"github.com/oklog/ulid/v2"
)

func SignupUser(ctx context.Context, db *sql.DB, username, email, sshKeyType, sshKeyId string) (iid string, err error) {
	iid = ulid.Make().String()
	tx, err := db.BeginTx(ctx, nil)
	if err != nil {
		return
	}
	defer func() {
		if err != nil {
			if rollbackErr := tx.Rollback(); rollbackErr != nil {
				println("SignupUser rollback error: ", rollbackErr)
			}
			return
		}
		err = tx.Commit()
	}()

	_, err = tx.ExecContext(ctx, "INSERT INTO users (id, username, email) VALUES (?, ?, ?)", iid, username, email)
	if err != nil {
		return
	}
	_, err = tx.ExecContext(ctx, "INSERT INTO ssh_public_keys (user_id, key_type, key) VALUES (?, ?, ?)", iid, sshKeyType, sshKeyId)
	if err != nil {
		return
	}

	return iid, nil
}

func AddToSSHQueue(ctx context.Context, db *sql.DB, id string) error {
	// drop existing items
	db.ExecContext(ctx, "DELETE FROM ssh_queue WHERE user_id = ?", id)
	_, err := db.ExecContext(ctx, "INSERT INTO ssh_queue (user_id) VALUES (?)", id)
	return err
}
