// Code generated by sqlc. DO NOT EDIT.
// versions:
//   sqlc v1.16.0

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
	if q.createCategoryStmt, err = db.PrepareContext(ctx, createCategory); err != nil {
		return nil, fmt.Errorf("error preparing query CreateCategory: %w", err)
	}
	if q.createCommentStmt, err = db.PrepareContext(ctx, createComment); err != nil {
		return nil, fmt.Errorf("error preparing query CreateComment: %w", err)
	}
	if q.createPostStmt, err = db.PrepareContext(ctx, createPost); err != nil {
		return nil, fmt.Errorf("error preparing query CreatePost: %w", err)
	}
	if q.createUserStmt, err = db.PrepareContext(ctx, createUser); err != nil {
		return nil, fmt.Errorf("error preparing query CreateUser: %w", err)
	}
	if q.deleteCategoryStmt, err = db.PrepareContext(ctx, deleteCategory); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteCategory: %w", err)
	}
	if q.deleteCommentStmt, err = db.PrepareContext(ctx, deleteComment); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteComment: %w", err)
	}
	if q.deletePostStmt, err = db.PrepareContext(ctx, deletePost); err != nil {
		return nil, fmt.Errorf("error preparing query DeletePost: %w", err)
	}
	if q.deleteUserStmt, err = db.PrepareContext(ctx, deleteUser); err != nil {
		return nil, fmt.Errorf("error preparing query DeleteUser: %w", err)
	}
	if q.findByEmailUserStmt, err = db.PrepareContext(ctx, findByEmailUser); err != nil {
		return nil, fmt.Errorf("error preparing query FindByEmailUser: %w", err)
	}
	if q.getCategoriesStmt, err = db.PrepareContext(ctx, getCategories); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategories: %w", err)
	}
	if q.getCategoryStmt, err = db.PrepareContext(ctx, getCategory); err != nil {
		return nil, fmt.Errorf("error preparing query GetCategory: %w", err)
	}
	if q.getCommentStmt, err = db.PrepareContext(ctx, getComment); err != nil {
		return nil, fmt.Errorf("error preparing query GetComment: %w", err)
	}
	if q.getCommentsStmt, err = db.PrepareContext(ctx, getComments); err != nil {
		return nil, fmt.Errorf("error preparing query GetComments: %w", err)
	}
	if q.getPostStmt, err = db.PrepareContext(ctx, getPost); err != nil {
		return nil, fmt.Errorf("error preparing query GetPost: %w", err)
	}
	if q.getPostRelationStmt, err = db.PrepareContext(ctx, getPostRelation); err != nil {
		return nil, fmt.Errorf("error preparing query GetPostRelation: %w", err)
	}
	if q.getPostsStmt, err = db.PrepareContext(ctx, getPosts); err != nil {
		return nil, fmt.Errorf("error preparing query GetPosts: %w", err)
	}
	if q.getUserStmt, err = db.PrepareContext(ctx, getUser); err != nil {
		return nil, fmt.Errorf("error preparing query GetUser: %w", err)
	}
	if q.getUsersStmt, err = db.PrepareContext(ctx, getUsers); err != nil {
		return nil, fmt.Errorf("error preparing query GetUsers: %w", err)
	}
	if q.updateCategoryStmt, err = db.PrepareContext(ctx, updateCategory); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateCategory: %w", err)
	}
	if q.updateCommentStmt, err = db.PrepareContext(ctx, updateComment); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateComment: %w", err)
	}
	if q.updatePostStmt, err = db.PrepareContext(ctx, updatePost); err != nil {
		return nil, fmt.Errorf("error preparing query UpdatePost: %w", err)
	}
	if q.updateUserStmt, err = db.PrepareContext(ctx, updateUser); err != nil {
		return nil, fmt.Errorf("error preparing query UpdateUser: %w", err)
	}
	return &q, nil
}

func (q *Queries) Close() error {
	var err error
	if q.createCategoryStmt != nil {
		if cerr := q.createCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCategoryStmt: %w", cerr)
		}
	}
	if q.createCommentStmt != nil {
		if cerr := q.createCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createCommentStmt: %w", cerr)
		}
	}
	if q.createPostStmt != nil {
		if cerr := q.createPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createPostStmt: %w", cerr)
		}
	}
	if q.createUserStmt != nil {
		if cerr := q.createUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing createUserStmt: %w", cerr)
		}
	}
	if q.deleteCategoryStmt != nil {
		if cerr := q.deleteCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCategoryStmt: %w", cerr)
		}
	}
	if q.deleteCommentStmt != nil {
		if cerr := q.deleteCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteCommentStmt: %w", cerr)
		}
	}
	if q.deletePostStmt != nil {
		if cerr := q.deletePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deletePostStmt: %w", cerr)
		}
	}
	if q.deleteUserStmt != nil {
		if cerr := q.deleteUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing deleteUserStmt: %w", cerr)
		}
	}
	if q.findByEmailUserStmt != nil {
		if cerr := q.findByEmailUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing findByEmailUserStmt: %w", cerr)
		}
	}
	if q.getCategoriesStmt != nil {
		if cerr := q.getCategoriesStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoriesStmt: %w", cerr)
		}
	}
	if q.getCategoryStmt != nil {
		if cerr := q.getCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCategoryStmt: %w", cerr)
		}
	}
	if q.getCommentStmt != nil {
		if cerr := q.getCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommentStmt: %w", cerr)
		}
	}
	if q.getCommentsStmt != nil {
		if cerr := q.getCommentsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getCommentsStmt: %w", cerr)
		}
	}
	if q.getPostStmt != nil {
		if cerr := q.getPostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostStmt: %w", cerr)
		}
	}
	if q.getPostRelationStmt != nil {
		if cerr := q.getPostRelationStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostRelationStmt: %w", cerr)
		}
	}
	if q.getPostsStmt != nil {
		if cerr := q.getPostsStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getPostsStmt: %w", cerr)
		}
	}
	if q.getUserStmt != nil {
		if cerr := q.getUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUserStmt: %w", cerr)
		}
	}
	if q.getUsersStmt != nil {
		if cerr := q.getUsersStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing getUsersStmt: %w", cerr)
		}
	}
	if q.updateCategoryStmt != nil {
		if cerr := q.updateCategoryStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCategoryStmt: %w", cerr)
		}
	}
	if q.updateCommentStmt != nil {
		if cerr := q.updateCommentStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateCommentStmt: %w", cerr)
		}
	}
	if q.updatePostStmt != nil {
		if cerr := q.updatePostStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updatePostStmt: %w", cerr)
		}
	}
	if q.updateUserStmt != nil {
		if cerr := q.updateUserStmt.Close(); cerr != nil {
			err = fmt.Errorf("error closing updateUserStmt: %w", cerr)
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
	db                  DBTX
	tx                  *sql.Tx
	createCategoryStmt  *sql.Stmt
	createCommentStmt   *sql.Stmt
	createPostStmt      *sql.Stmt
	createUserStmt      *sql.Stmt
	deleteCategoryStmt  *sql.Stmt
	deleteCommentStmt   *sql.Stmt
	deletePostStmt      *sql.Stmt
	deleteUserStmt      *sql.Stmt
	findByEmailUserStmt *sql.Stmt
	getCategoriesStmt   *sql.Stmt
	getCategoryStmt     *sql.Stmt
	getCommentStmt      *sql.Stmt
	getCommentsStmt     *sql.Stmt
	getPostStmt         *sql.Stmt
	getPostRelationStmt *sql.Stmt
	getPostsStmt        *sql.Stmt
	getUserStmt         *sql.Stmt
	getUsersStmt        *sql.Stmt
	updateCategoryStmt  *sql.Stmt
	updateCommentStmt   *sql.Stmt
	updatePostStmt      *sql.Stmt
	updateUserStmt      *sql.Stmt
}

func (q *Queries) WithTx(tx *sql.Tx) *Queries {
	return &Queries{
		db:                  tx,
		tx:                  tx,
		createCategoryStmt:  q.createCategoryStmt,
		createCommentStmt:   q.createCommentStmt,
		createPostStmt:      q.createPostStmt,
		createUserStmt:      q.createUserStmt,
		deleteCategoryStmt:  q.deleteCategoryStmt,
		deleteCommentStmt:   q.deleteCommentStmt,
		deletePostStmt:      q.deletePostStmt,
		deleteUserStmt:      q.deleteUserStmt,
		findByEmailUserStmt: q.findByEmailUserStmt,
		getCategoriesStmt:   q.getCategoriesStmt,
		getCategoryStmt:     q.getCategoryStmt,
		getCommentStmt:      q.getCommentStmt,
		getCommentsStmt:     q.getCommentsStmt,
		getPostStmt:         q.getPostStmt,
		getPostRelationStmt: q.getPostRelationStmt,
		getPostsStmt:        q.getPostsStmt,
		getUserStmt:         q.getUserStmt,
		getUsersStmt:        q.getUsersStmt,
		updateCategoryStmt:  q.updateCategoryStmt,
		updateCommentStmt:   q.updateCommentStmt,
		updatePostStmt:      q.updatePostStmt,
		updateUserStmt:      q.updateUserStmt,
	}
}
