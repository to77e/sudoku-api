-- +goose Up
create table public.sudoku
(
    uuid       uuid primary key   default gen_random_uuid(),
    sudoku     int[]     not null,
    k          int       not null,
    created_at timestamp not null default now(),
    update_at  timestamp not null default now()
);

-- +goose Down
drop table public.sudoku;

