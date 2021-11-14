CREATE TABLE tasks
(
    id         serial                                      not null unique,
    user_id    int references users (id) on delete cascade not null,
    name       varchar(255)                                not null,
    created_at timestamp,
    updated_at timestamp
)