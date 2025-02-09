CREATE TABLE IF NOT EXISTS tests (
  id integer PRIMARY KEY,
  date timestamp NOT NULL,
  description text
);

CREATE TABLE IF NOT EXISTS students (
  id integer PRIMARY KEY,
  name text NOT NULL,
  class text NOT NULL,
  UNIQUE (name, class)
);

CREATE TABLE IF NOT EXISTS questions (
  id integer PRIMARY KEY,
  test_id integer NOT NULL,
  parent_id integer,
  max_points float NOT NULL,
  name text,
  rank integer NOT NULL,
  FOREIGN KEY (test_id) REFERENCES tests (id),
  FOREIGN KEY (parent_id) REFERENCES questions (id),
  UNIQUE (test_id, parent_id, rank)
);

CREATE TABLE IF NOT EXISTS results (
  id integer PRIMARY KEY,
  student_id integer NOT NULL,
  question_id integer NOT NULL,
  points float NOT NULL,
  FOREIGN KEY (student_id) REFERENCES students (id),
  FOREIGN KEY (question_id) REFERENCES questions (id),
  UNIQUE (student_id, question_id)
);


