ALTER TABLE
    costumes DROP COLUMN condition;

ALTER TABLE
    clothes
ADD
    COLUMN condition conditions NOT NULL DEFAULT 'нормальное';