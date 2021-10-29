ALTER TABLE
    costumes
ADD
    COLUMN is_decor bool NOT NULL DEFAULT false,
ADD
    COLUMN is_archived bool NOT NULL DEFAULT false;

ALTER TABLE
    clothes DROP COLUMN is_decor bool,
    DROP is_archived bool;