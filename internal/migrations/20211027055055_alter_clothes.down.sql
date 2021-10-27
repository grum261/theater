ALTER TABLE
    costumes ADD COLUMN condition conditions not null default 'нормальное';

ALTER TABLE
    clothes
drop
    COLUMN condition;