CREATE ROLE todoisgo SUPERUSER CREATEDB CREATEROLE NOINHERIT LOGIN;


CREATE DATABASE todoisdb
  WITH OWNER = todoisgo
       ENCODING = 'UTF8'
       TABLESPACE = pg_default
       CONNECTION LIMIT = -1;


-- Table: todo_user

-- DROP TABLE todo_user;

CREATE TABLE todo_user
(
  name character varying(50),
  id serial,
  email character varying(50) NOT NULL,
  active boolean,
  password character varying(200),
  CONSTRAINT pk_user PRIMARY KEY (id),
  CONSTRAINT unique_email UNIQUE (email)
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo_user
  OWNER TO postgres;
GRANT ALL ON TABLE todo_user TO postgres;
GRANT ALL ON TABLE todo_user TO public;


-- Table: todo_task

-- DROP TABLE todo_task;

CREATE TABLE todo_task
(
  description character varying(300),
  id serial,
  complete boolean DEFAULT false,
  created_at timestamp without time zone DEFAULT now(),
  complete_at timestamp without time zone DEFAULT '2000-01-01 00:00:00'::timestamp without time zone,
  user_id bigint,
  CONSTRAINT pk_task PRIMARY KEY (id),
  CONSTRAINT fk_task_user FOREIGN KEY (user_id)
      REFERENCES todo_user (id) MATCH SIMPLE
      ON UPDATE NO ACTION ON DELETE CASCADE
)
WITH (
  OIDS=FALSE
);
ALTER TABLE todo_task
  OWNER TO todoisgo;
GRANT ALL ON TABLE todo_task TO todoisgo;
GRANT ALL ON TABLE todo_task TO public;

