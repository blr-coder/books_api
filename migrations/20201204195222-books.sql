-- +migrate Up
CREATE TABLE books
(
    id         BIGSERIAL PRIMARY KEY,
    title      VARCHAR(256) NOT NULL,
    author     VARCHAR(256) NOT NULL,
    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON TABLE  books            IS 'книги';
COMMENT ON COLUMN books.title      IS 'название';
COMMENT ON COLUMN books.author     IS 'автор';
COMMENT ON COLUMN books.created_at IS 'дата создания';
COMMENT ON COLUMN books.updated_at IS 'дата обновления';
COMMENT ON COLUMN books.deleted_at IS 'дата удаления';

-- +migrate Down
DROP TABLE books;
