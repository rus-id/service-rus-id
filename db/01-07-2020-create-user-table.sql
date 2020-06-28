-- drop table if exists users;

create table if not exists users (
    id            uuid    primary key,
    first_name    text    not null,
    middle_name   text,
    last_name     text    not null,
    --
    passport           jsonb,
    driving_license    jsonb,
    --
    created_at timestamp with time zone default now() not null,
    updated_at timestamp with time zone default now() not null
);

create index on users using gin ((passport -> 'serial'));
create index on users using gin ((passport -> 'number'));
create index on users using gin ((driving_license -> 'serial'));
create index on users using gin ((driving_license -> 'number'));