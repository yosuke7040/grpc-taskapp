// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.26.0

package db

import (
	"context"
	"database/sql"
	"fmt"
)

type DBTX interface {
	ExecContext(context.Context, string, ...interface{}) (sql.Result, error)
	PrepareContext(context.Context, string) (*sql.Stmt, error)
	QueryContext(context.Context, string, ...interface{}) (*sql.Rows, error)
	QueryRowContext(context.Context, string, ...interface{}) *sql.Row
}

func New(db DBTX) *Queries {
	return &Queries{db: db}
}

func Prepare(ctx context.Context, db DBTX) (*Queries, error) {
	q := Queries{db: db}
	var err error
	if q.createTaskStmt, err = db.PrepareContext(ctx, createTask); err != nil {
		return nil, fmt.Errorf("error preparing query CreateTask: %w", err)
	}
	if q.deleteTaskStmt, err = db.PrepareContext(ctx, deleteTask); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteTask: %w", err)
	}
	if q.findTaskByIDStmt, err = db.PrepareContext(ctx, findTaskByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindTaskByID: %w", err)
	}
	if q.findTasksByUserIDStmt, err = db.PrepareContext(ctx, findTasksByUserID); err != nil {
		return nil, fmt.Errorf("error preparing query FindTasksByUserID: %w", err)
	}
	if q.findUserByEmailStmt, err = db.PrepareContext(ctx, findUserByEmail); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByEmail: %w", err)
	}
	if q.findUserByIDStmt, err = db.PrepareContext(ctx, findUserByID); err != nil {
		return nil, fmt.Errorf("error preparing query FindUserByID: %w", err)
	}
	if q.updateTaskStmt, err = db.PrepareContext(ctx, updateTask); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateTask: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createTaskStmt != nil {
		if cerr := q.createTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createTaskStmt: %w", cerr)
		}
	}
	if q.deleteTaskStmt != nil {
		if cerr := q.deleteTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteTaskStmt: %w", cerr)
		}
	}
	if q.findTaskByIDStmt != nil {
		if cerr := q.findTaskByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTaskByIDStmt: %w", cerr)
		}
	}
	if q.findTasksByUserIDStmt != nil {
		if cerr := q.findTasksByUserIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findTasksByUserIDStmt: %w", cerr)
		}
	}
	if q.findUserByEmailStmt != nil {
		if cerr := q.findUserByEmailStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByEmailStmt: %w", cerr)
		}
	}
	if q.findUserByIDStmt != nil {
		if cerr := q.findUserByIDStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findUserByIDStmt: %w", cerr)
		}
	}
	if q.updateTaskStmt != nil {
		if cerr := q.updateTaskStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateTaskStmt: %w", cerr)
		}
	}
	return err
}

func (q *Queries) exec(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (sql.Result, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).ExecContext(ctx, args...)
	case stmt != nil:
		return stmt.ExecContext(ctx, args...)
	default:
		return q.db.ExecContext(ctx, query, args...)
	}
}

func (q *Queries) query(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) (*sql.Rows, error) {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryContext(ctx, args...)
	default:
		return q.db.QueryContext(ctx, query, args...)
	}
}

func (q *Queries) queryRow(ctx context.Context, stmt *sql.Stmt, query string, args ...interface{}) *sql.Row {
	switch {
	case stmt != nil && q.tx != nil:
		return q.tx.StmtContext(ctx, stmt).QueryRowContext(ctx, args...)
	case stmt != nil:
		return stmt.QueryRowContext(ctx, args...)
	default:
		return q.db.QueryRowContext(ctx, query, args...)
	}
}

type Queries struct {
	db                    DBTX
	tx                    *sql.Tx
	createTaskStmt        *sql.Stmt
	deleteTaskStmt        *sql.Stmt
	findTaskByIDStmt      *sql.Stmt
	findTasksByUserIDStmt *sql.Stmt
	findUserByEmailStmt   *sql.Stmt
	findUserByIDStmt      *sql.Stmt
	updateTaskStmt        *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                    tx,
		tx:                    tx,
		createTaskStmt:        q.createTaskStmt,
		deleteTaskStmt:        q.deleteTaskStmt,
		findTaskByIDStmt:      q.findTaskByIDStmt,
		findTasksByUserIDStmt: q.findTasksByUserIDStmt,
		findUserByEmailStmt:   q.findUserByEmailStmt,
		findUserByIDStmt:      q.findUserByIDStmt,
		updateTaskStmt:        q.updateTaskStmt,
	}
}