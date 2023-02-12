CREATE TABLE "users"(
    "id" serial PRIMARY KEY,
    "firstname" VARCHAR(100) NOT NULL,
    "lastname" varchar(100) NOT NULL,
    "email" varchar(100) UNIQUE NOT NULL,
    "password" varchar(100) NOT NULL
);

CREATE TABLE "categories" (
    "id" serial PRIMARY KEY,
    "name" varchar(200) NOT NULL
);

CREATE TABLE "comments" (
    "id" serial PRIMARY KEY,
    "id_post_comment" int NOT NULL,
    "user_name_comment" varchar(200) NOT NULL,
    "comment" varchar(200) NOT NULL,
    FOREIGN KEY (id_post_comment)
    REFERENCES posts(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);

CREATE TABLE "posts" (
    "id" serial PRIMARY KEY,
    "title" varchar(200) NOT NULL,
    "slug" varchar(200) NOT NULL,
    "img" varchar(100) NOT NULL,
    "body" text NOT NULL,
    "category_id" int NOT NULL,
    "user_id" int NOT NULL,
    "user_name" varchar(200) NOT NULL,
    FOREIGN KEY (category_id)
    REFERENCES categories(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE,
    FOREIGN KEY (user_id)
    REFERENCES users(id)
    ON UPDATE CASCADE
    ON DELETE CASCADE
);