package repository

import (
	"context"
	"database/sql"
	"fmt"
	"log"
	"muxblog/schemas"
)

type commentRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewCommentRepository(ctx context.Context, db *sql.DB) *commentRepository {
	return &commentRepository{ctx: ctx, db: db}
}

func (r *commentRepository) GetAll() ([]schemas.Comment, error) {
	var comment schemas.Comment
	var commentModel []schemas.Comment

	row, err := r.db.QueryContext(r.ctx, "SELECT * FROM comments ORDER BY id DESC")

	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		if err = row.Scan(&comment.ID, &comment.ID, &comment.IDPOSTCOMMENT, &comment.USERNAMECOMMENT, &comment.COMMENT); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}

		commentModel = append(commentModel, comment)
	}

	return commentModel, nil
}

func (r *commentRepository) GetID(id int) (schemas.Comment, error) {
	var comment schemas.Comment

	row, err := r.db.QueryContext(r.ctx, "SELECT id, id_post_comment,user_name_comment,comment FROM comments WHERE id = ?", id)

	if err != nil {
		log.Fatal("Error Query Comment: " + err.Error())
		return comment, err
	}

	for row.Next() {
		err := row.Scan(&comment.ID, &comment.IDPOSTCOMMENT, &comment.USERNAMECOMMENT, &comment.COMMENT)
		if err != nil {
			return comment, err
		}
	}

	return comment, nil
}

func (r *commentRepository) Create(comment *schemas.Comment) (schemas.Comment, error) {
	var commentModel schemas.Comment

	row, err := r.db.PrepareContext(r.ctx, "INSERT INTO comments (id_post_comment,user_name_comment,comment) VALUES (?,?,?)")

	fmt.Println(comment)

	if err != nil {
		return commentModel, err
	}

	res, err := row.ExecContext(r.ctx, &comment.IDPOSTCOMMENT, &comment.USERNAMECOMMENT, &comment.COMMENT)

	if err != nil {
		return commentModel, err
	}

	rowID, err := res.LastInsertId()

	if err != nil {
		return commentModel, err
	}

	result, err := r.GetID(int(rowID))

	if err != nil {
		return commentModel, err
	}

	return result, nil
}

func (r *commentRepository) Update(input *schemas.Comment) (schemas.Comment, error) {
	row, err := r.db.PrepareContext(r.ctx, "UPDATE comments SET id_post_comment=?,user_name_comment=?,comment=? WHERE id=?")

	var commentModel schemas.Comment

	if err != nil {
		log.Fatal("Error dibagian update: ", err.Error())
		return commentModel, err
	}

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err)
		return commentModel, err
	}

	_, queryError := row.ExecContext(r.ctx, &input.IDPOSTCOMMENT, &input.USERNAMECOMMENT, &input.COMMENT, &input.ID)
	fmt.Println(input)

	if queryError != nil {
		log.Fatal(queryError.Error())
		return commentModel, err
	}

	res, err := r.GetID(input.ID)

	if err != nil {
		return commentModel, err
	}

	return res, nil
}

func (r *commentRepository) Delete(id int) error {

	row, err := r.db.PrepareContext(r.ctx, "DELETE FROM comments WHERE id = ?")

	if err != nil {
		return err
	}

	_, queryError := row.ExecContext(r.ctx, id)

	if queryError != nil {
		return err
	}

	return nil

}
