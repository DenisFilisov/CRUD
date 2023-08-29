CREATE TABLE users
(
    id            serial       not null unique,
    name          varchar(255) not null,
    username      varchar(255) not null unique,
    password_hash varchar(255) not null
);

CREATE TABLE news
(
    id          serial       not null unique,
    description varchar(255) not null
);

CREATE TABLE followers
(
    id      serial                                      not null unique,
    user_id int references users (id) on delete cascade not null,
    news_id int references news (id) on delete cascade  not null
)