CREATE TABLE performances (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    start_date timestamp without time zone NOT NULL,
    created_at timestamp without time zone NOT NULL,
    updated_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP
);

CREATE UNIQUE INDEX performance_name_idx ON performances (name);

ALTER TABLE
    costumes
ADD
    COLUMN size int NOT NULL DEFAULT 0,
ADD
    COLUMN location varchar NOT NULL DEFAULT '',
ADD
    COLUMN created_at timestamp without time zone NOT NULL,
ADD
    COLUMN update_at timestamp without time zone NOT NULL DEFAULT CURRENT_TIMESTAMP;

CREATE TABLE performances_costumes (
    performance_id int NOT NULL REFERENCES performances,
    costume_id int NOT NULL REFERENCES costumes
);

CREATE INDEX ON performances_costumes (performance_id, costume_id);