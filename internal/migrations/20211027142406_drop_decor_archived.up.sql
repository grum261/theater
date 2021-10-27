ALTER TABLE
    costumes DROP COLUMN is_decor,
    DROP COLUMN is_archived;

ALTER TABLE
    clothes
ADD
    COLUMN is_decor bool NOT NULL DEFAULT false,
ADD
    is_archived bool NOT NULL DEFAULT false;