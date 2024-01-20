package mysql

import (
	"database/sql"
	"muxblog/pkg/logger"
)

func NewClient(logger logger.Logger) (*sql.DB, error) {
	con, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/muxblog")

	if err != nil {
		logger.Error(err.Error())
		return nil, err
	}

	con.SetMaxIdleConns(10)
	con.SetMaxOpenConns(100)
	con.SetConnMaxLifetime(10)

	return con, nil
}
