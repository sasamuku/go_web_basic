create table if not exists posts (
  id      serial primary key,
  content text,
  author  varchar(255)
);
