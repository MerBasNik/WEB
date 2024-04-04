CREATE TABLE users
(
    id                  serial       PRIMARY KEY,
    password_hash       varchar(255) not null,
    email	            varchar(255) not null unique
);

CREATE TABLE users_profile
(
    id                  serial       PRIMARY KEY,
    name                varchar(255) not null,
    surname             varchar(255) not null,
    photo               varchar(255) not null,
    city	            varchar(255) not null,
    telegram	        varchar(255) not null,
    findstatus          boolean      not null default false
);

CREATE TABLE chat_lists
(
    id          serial       PRIMARY KEY,
    title       varchar(255) not null
    --photo 
    -- listchat references chat_items on delete cascade not null 
    -- description varchar(255)
);

CREATE TABLE users_hobby
(
    id          serial       PRIMARY KEY,
    description varchar(255) not null
);

CREATE TABLE users_profile_lists
(
    id           serial                                              PRIMARY KEY,
    user_id      int references users (id)         on delete cascade not null,
    profile_id   int references users_profile (id) on delete cascade not null
);

CREATE TABLE users_hobby_lists
(
    id           serial                                              PRIMARY KEY,
    user_id      int references users (id)         on delete cascade not null,
    userhobby_id int references users_hobby (id)    on delete cascade not null
);

CREATE TABLE users_chat_lists
(
    id           serial                                              PRIMARY KEY,
    user_id      int references users (id)         on delete cascade not null,
    chatlists_id int references chat_lists (id)    on delete cascade not null
);

CREATE TABLE chat_items
(
    id          serial       PRIMARY KEY,
    title       varchar(255) not null,
    description varchar(255)
    -- done        boolean      not null default false
);

CREATE TABLE items_lists
(
    id           serial                                           PRIMARY KEY,
    chatitems_id int references chat_items (id) on delete cascade not null,
    chatlists_id int references chat_lists (id) on delete cascade not null
);