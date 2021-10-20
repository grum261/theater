CREATE TYPE conditions AS enum ('плохое', 'нормальное', 'хорошее');

CREATE TABLE colors (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX color_name_idx ON colors (name);

CREATE TABLE materials (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX material_name_idx ON materials (name);

CREATE TABLE designers (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX designer_name_idx ON designers (name);

CREATE TABLE clothes_types (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX clothes_types_name_idx ON clothes_types (name);

CREATE TABLE clothes (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    type_id serial NOT NULL REFERENCES clothes_types
);

CREATE UNIQUE INDEX cloth_name_idx ON clothes (name);

CREATE TABLE clothes_colors (
    cloth_id serial NOT NULL REFERENCES clothes ON DELETE CASCADE,
    color_id serial NOT NULL REFERENCES colors ON DELETE CASCADE
);

CREATE INDEX ON clothes_colors (cloth_id, color_id);

CREATE TABLE clothes_materials (
    cloth_id serial NOT NULL REFERENCES clothes ON DELETE CASCADE,
    material_id serial NOT NULL REFERENCES materials ON DELETE CASCADE
);

CREATE INDEX ON clothes_materials (cloth_id, material_id);

CREATE TABLE costumes (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    size int NOT NULL DEFAULT 0,
    condition conditions NOT NULL DEFAULT 'нормальное',
    location varchar,
    description varchar,
    image_front varchar NOT NULL,
    image_back varchar NOT NULL,
    image_sideway varchar NOT NULL,
    image_details varchar NOT NULL,
    is_archived bool NOT NULL DEFAULT false,
    designer varchar NOT NULL
);

CREATE UNIQUE INDEX costume_name_idx ON costumes (name);

CREATE TABLE costumes_clothes (
    costume_id serial NOT NULL REFERENCES costumes ON DELETE CASCADE,
    cloth_id serial NOT NULL REFERENCES clothes ON DELETE CASCADE
);

CREATE INDEX ON costumes_clothes (costume_id, cloth_id);

CREATE TABLE performances (
    id serial PRIMARY KEY,
    name varchar NOT NULL,
    starting_at timestamp without time zone NOT NULL,
    duration int NOT NULL DEFAULT 0,
    location varchar NOT NULL
);

CREATE TABLE performances_costumes (
    performance_id serial NOT NULL REFERENCES performances ON DELETE CASCADE,
    costume_id serial NOT NULL REFERENCES costumes ON DELETE CASCADE
);

CREATE INDEX ON performances_costumes (performance_id, costume_id);