package repository

import (
	"context"
	"database/sql"
	"log"
	"muxblog/schemas"
	"muxblog/utils"
)

type userRepository struct {
	ctx context.Context
	db  *sql.DB
}

func NewUserRepository(ctx context.Context, db *sql.DB) *userRepository {
	return &userRepository{ctx: ctx, db: db}
}

func (r *userRepository) GetAll() ([]schemas.Users, error) {
	var userSchema schemas.Users
	var usersModel []schemas.Users

	row, err := r.db.QueryContext(r.ctx, "SELECT id, firstname, lastname, email FROM users ORDER BY id DESC")

	if err != nil {
		log.Fatal(err.Error())
	}

	for row.Next() {
		if err = row.Scan(&userSchema.ID, &userSchema.FirstName, &userSchema.LastName, &userSchema.Email); err != nil && sql.ErrNoRows != nil {
			return nil, err
		}
		usersModel = append(usersModel, userSchema)
	}

	return usersModel, nil

}

func (r *userRepository) GetBYID(id int) (schemas.Users, error) {
	var userModel schemas.Users

	row, err := r.db.QueryContext(r.ctx, "SELECT id, firstname,lastname,email,password FROM users WHERE id = ?", id)

	if err != nil {
		log.Fatal("Error Query Category: " + err.Error())
		return userModel, err
	}

	for row.Next() {
		err := row.Scan(&userModel.ID, &userModel.FirstName, &userModel.LastName, &userModel.Email, &userModel.Password)
		if err != nil {
			return userModel, err
		}
	}

	return userModel, nil
}

func (r *userRepository) Create(input *schemas.Users) (schemas.Users, error) {
	var userModel schemas.Users

	// _, err := r.FindByEmail(input.Email)

	// if err != nil {
	// 	log.Fatal("Error Find By Email", err.Error())
	// }

	row, err := r.db.PrepareContext(r.ctx, "INSERT INTO users (firstname, lastname, email, password) VALUES (?, ?, ?, ?)")

	if err != nil {
		log.Fatal("Error dibagian Create: ", err.Error())
	}

	res, err := row.ExecContext(r.ctx, input.FirstName, input.LastName, input.Email, input.Password)

	if err != nil {
		log.Fatal("Error dibagian Create: ", err.Error())
	}

	rowID, err := res.LastInsertId()

	if err != nil {
		return userModel, err
	}

	result, err := r.GetBYID(int(rowID))

	if err != nil {
		return userModel, err
	}

	return result, nil
}

func (r *userRepository) Update(input *schemas.Users) (schemas.Users, error) {

	resbyid, err := r.GetBYID(input.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = utils.ComparePassword(resbyid.Password, input.Password)

	if err != nil {
		log.Fatal(err.Error())
	}

	row, err := r.db.PrepareContext(r.ctx, "UPDATE users SET firstname=?,lastname=?,email=?,password=? WHERE id=?")

	if err != nil && err != sql.ErrNoRows {
		log.Fatal(err.Error())
	}

	_, queryError := row.ExecContext(r.ctx, input.FirstName, input.LastName, input.Email, utils.HashPassword(input.Password), input.ID)

	if queryError != nil {
		log.Fatal(err.Error())
	}

	res, err := r.GetBYID(input.ID)

	if err != nil {
		log.Fatal(err.Error())
	}

	return res, err

}

func (r *userRepository) Delete(id int) error {
	row, err := r.db.PrepareContext(r.ctx, "DELETE FROM users WHERE id = ?")

	if err != nil {
		log.Fatal("Repository delete:", err.Error())
	}
	_, queryError := row.ExecContext(r.ctx, id)

	if queryError != nil {
		log.Fatal("Repository delete: ", err.Error())
	}
	return nil
}

func (r *userRepository) FindByEmail(email string) (schemas.Users, error) {
	var usersModel schemas.Users
	row, err := r.db.QueryContext(r.ctx, "SELECT id, firstname, lastname, email, password FROM users WHERE email = ?", email)

	if err != nil {
		log.Fatal("Error Query Users: " + err.Error())
	}

	for row.Next() {
		err := row.Scan(&usersModel.ID, &usersModel.FirstName, &usersModel.LastName, &usersModel.Email, &usersModel.Password)

		if err != nil {
			log.Fatal(err.Error())
		}
	}
	return usersModel, nil
}

func (r *userRepository) Login(input *schemas.AuthLogin) (schemas.Users, error) {
	res, err := r.FindByEmail(input.Email)

	if err != nil {
		log.Fatal(err.Error())
	}

	err = utils.ComparePassword(res.Password, input.Password)

	if err != nil {
		log.Fatal("Compare Password: ", err.Error())
	}

	return res, nil
}
