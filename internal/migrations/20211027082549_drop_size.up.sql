ALTER TABLE
    costumes DROP COLUMN size;

ALTER TABLE
    clothes
ADD
    COLUMN size int NOT NULL DEFAULT 0;