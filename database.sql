CREATE TABLE mst_users (
  id serial8 PRIMARY KEY,
  full_name varchar(60) NOT NULL,
  phone_number varchar(15) NOT NULL,
  password_hash varchar NOT NULL,
  created_by int8 default 0 NOT NULL,
  created_at timestamptz NOT NULL,
  updated_by int8,
  updated_at timestamptz,
  deleted_at timestamptz
);

CREATE TABLE trx_sessions (
  id serial8 PRIMARY KEY,
  user_id int8 NOT NULL,
  login_at timestamptz NOT NULL
);