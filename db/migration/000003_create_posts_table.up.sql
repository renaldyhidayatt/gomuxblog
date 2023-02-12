
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