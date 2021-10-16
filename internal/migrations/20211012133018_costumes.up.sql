CREATE TYPE conditions AS ENUM ('нормальное', 'хорошее', 'плохое');

CREATE TABLE costumes (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    desciption varchar,
    is_decor bool NOT NULL DEFAULT false,
    condition conditions
);

CREATE TABLE colors (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX color_name_idx ON colors (name);

CREATE TABLE costumes_colors (
    costume_id int NOT NULL REFERENCES costumes,
    color_id int NOT NULL REFERENCES colors
);

CREATE INDEX ON costumes_colors (costume_id, color_id);

CREATE TABLE materials (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX materials_name_idx ON materials (name);

CREATE TABLE costumes_materials (
    costume_id int NOT NULL REFERENCES costumes,
    material_id int NOT NULL REFERENCES materials
);

CREATE INDEX ON costumes_materials (costume_id, material_id);

CREATE TABLE designers (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE UNIQUE INDEX designer_name_idx ON designers (name);