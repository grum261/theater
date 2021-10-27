ALTER TABLE
    costumes DROP COLUMN designer,
    DROP COLUMN location;

ALTER TABLE
    clothes
ADD
    COLUMN designer varchar,
ADD
    COLUMN location varchar;