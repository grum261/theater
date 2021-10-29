ALTER TABLE
    costumes
ADD
    COLUMN size int NOT NULL DEFAULT 0;

ALTER TABLE
    clothes DROP COLUMN size;