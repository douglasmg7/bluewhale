-- drop table if exists entrance;
-- enable foreign keys
-- not working, reset to off when back to db
pragma foreign_keys = on;

create table entrance (
  id integer primary key autoincrement,
  user_id integer not null,
  school varchar(64) not null,
  created date not null,
  foreign key(user_id) references user(id)
);

create table user (
  id integer primary key autoincrement,
  name varchar(64) not null,
  -- carteira de identidade
  rg varchar(64) null,
  -- mobile number
  mobile varchar(64) null,
  email varchar(64) null,
  created date not null
);
create index idx_user_name on user(name);

-- create table parent(a primary key, b unique, c, d, e, f);
-- create unique index i1 on parent(c, d);
-- create index i2 on parent(e);
-- create unique index i3 on parent(f collate nocase);