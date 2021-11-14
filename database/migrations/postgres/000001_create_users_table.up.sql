CREATE TABLE users
(
    id         serial       not null unique,
    login      varchar(255) not null unique,
    password   varchar(255) not null,
    created_at timestamp
)