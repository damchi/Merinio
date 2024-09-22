CREATE TABLE branches (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  parent_id INT REFERENCES branches(id) ON DELETE SET NULL,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  state INTEGER DEFAULT 1,
  is_root BOOLEAN DEFAULT FALSE
);

CREATE TABLE requirements (
  id SERIAL PRIMARY KEY,
  name VARCHAR(255) NOT NULL,
  created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
  state INTEGER DEFAULT 1
);

CREATE TABLE branch_requirements (
 branch_id INT REFERENCES branches(id) ON DELETE CASCADE,
 requirement_id INT REFERENCES requirements(id) ON DELETE CASCADE,
 created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
 state INTEGER DEFAULT 1,
 PRIMARY KEY (branch_id, requirement_id)
);

CREATE TABLE restrictions (
id SERIAL PRIMARY KEY,
name VARCHAR(255) NOT NULL,
created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
state INTEGER DEFAULT 1
);

CREATE TABLE branch_restrictions (
id SERIAL PRIMARY KEY,
branch_id INT REFERENCES branches(id) ON DELETE CASCADE,
restriction_id INT REFERENCES restrictions(id) ON DELETE CASCADE,
created TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
modified TIMESTAMP DEFAULT CURRENT_TIMESTAMP,
UNIQUE (branch_id, restriction_id)
);


INSERT INTO branches (name, parent_id, is_root) VALUES('Branch A / Root branch', NULL, TRUE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch B', 1, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch C', 1, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch D', 1, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch E', 2, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch F', 2, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch G', 3, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch H', 4, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch I', 4, FALSE);
INSERT INTO branches (name, parent_id, is_root) VALUES('Branch J', 4, FALSE);


INSERT INTO requirements (name) VALUES ('Requirement A');
INSERT INTO requirements (name) VALUES ('Requirement B');
INSERT INTO requirements (name) VALUES ('Requirement C');
INSERT INTO requirements (name) VALUES ('Requirement D');
INSERT INTO requirements (name) VALUES ('Requirement E');
INSERT INTO requirements (name) VALUES ('Requirement F');
INSERT INTO requirements (name) VALUES ('Requirement G');
INSERT INTO requirements (name) VALUES ('Requirement H');
INSERT INTO requirements (name) VALUES ('Requirement I');
INSERT INTO requirements (name) VALUES ('Requirement J');

INSERT INTO branch_requirements (branch_id, requirement_id) VALUES (1, 1);
INSERT INTO branch_requirements (branch_id, requirement_id) VALUES(1, 2);
INSERT INTO branch_requirements (branch_id, requirement_id) VALUES(2, 3);
INSERT INTO branch_requirements (branch_id, requirement_id) VALUES(3, 4);


INSERT INTO restrictions (name) VALUES('Restriction A');
INSERT INTO restrictions (name) VALUES('Restriction B');
INSERT INTO restrictions (name) VALUES('Restriction C');
INSERT INTO restrictions (name) VALUES('Restriction D');
INSERT INTO restrictions (name) VALUES('Restriction E');
INSERT INTO restrictions (name) VALUES('Restriction F');
INSERT INTO restrictions (name) VALUES('Restriction G');
INSERT INTO restrictions (name) VALUES('Restriction H');
INSERT INTO restrictions (name) VALUES('Restriction I');
INSERT INTO restrictions (name) VALUES('Restriction J');

INSERT INTO branch_restrictions (branch_id, restriction_id) VALUES (1, 1);
INSERT INTO branch_restrictions (branch_id, restriction_id) VALUES(2, 2);
INSERT INTO branch_restrictions (branch_id, restriction_id) VALUES(2, 3);
INSERT INTO branch_restrictions (branch_id, restriction_id) VALUES(2, 4);
INSERT INTO branch_restrictions (branch_id, restriction_id) VALUES(3, 5);
