## Go Chi Blog with Sqlc

Projek ini adalah implementasi blog sederhana menggunakan framework Go Chi dan SQLC untuk berinteraksi dengan database. Go Chi adalah sebuah framework web minimalis untuk Go (Golang), sementara SQLC adalah alat untuk menghasilkan kode SQL secara langsung dari query SQL.

## Install

```
go mod tidy
```

### Install go-migrate

https://github.com/golang-migrate/migrate/tree/master/cmd/migrate

### Install sqlc

https://docs.sqlc.dev/en/latest/overview/install.html

```
go install github.com/sqlc-dev/sqlc/cmd/sqlc@latest
```

### dbdocs

dbdocs.io is a free, simple tool to create web-based documentation for your database. This guide will help you setup and start using dbdocs in less than 5 minutes.
https://dbdocs.io/docs

```
npm install -g dbdocs
```

### sql2dbml and dbml2sql

DBML comes with a built-in CLI which can be used to convert between different formats from the command line

```
npm install -g @dbml/cli
```

## Struktur Database

<img src="./images/gomuxblog.png" alt="database" />
