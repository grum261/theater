DROP TABLE clothes_colors;

DROP TABLE clothes_materials;

DROP TABLE colors;

DROP TABLE materials;

CREATE TABLE tags_types (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE
);

CREATE TABLE tags (
    id serial PRIMARY KEY,
    name varchar NOT NULL UNIQUE,
    type_id serial REFERENCES tags_types
);

CREATE UNIQUE INDEX tag_name_idx ON tags (name);

CREATE TABLE costumes_tags (
    costume_id serial NOT NULL REFERENCES costumes,
    tag_id serial NOT NULL REFERENCES tags
);

CREATE INDEX ON costumes_tags (costume_id, tag_id);