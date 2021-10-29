ALTER TABLE
    clothes DROP COLUMN designer,
    DROP COLUMN location;

ALTER TABLE
    costumes
ADD
    COLUMN designer varchar,
ADD
    COLUMN location varchar;