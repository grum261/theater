ALTER TABLE
    clothes
ADD
    COLUMN size varchar NOT NULL,
ADD
    COLUMN designer varchar NOT NULL,
ADD
    COLUMN location varchar,
ADD
    condition conditions NOT NULL DEFAULT 'нормальное',
ADD
    COLUMN is_archived bool NOT NULL DEFAULT false,
ADD
    COLUMN COMMENT varchar;

ALTER TABLE
    costumes DROP COLUMN size,
    DROP COLUMN location,
    DROP COLUMN designer,
    DROP COLUMN condition,
    add column is_decor bool not null default false;