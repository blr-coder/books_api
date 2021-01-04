-- +migrate Up
CREATE TYPE user_role_t AS ENUM('admin', 'user');

CREATE TABLE users
(
    id         BIGSERIAL PRIMARY KEY,
    email      VARCHAR(256) NOT NULL UNIQUE,
    password   VARCHAR(256) NOT NULL,

    role       user_role_t NOT NULL,

    created_at TIMESTAMP WITH TIME ZONE,
    updated_at TIMESTAMP WITH TIME ZONE,
    deleted_at TIMESTAMP WITH TIME ZONE
);

COMMENT ON TABLE  users            IS 'пользователи';
COMMENT ON COLUMN users.id         IS 'идентификатор';
COMMENT ON COLUMN users.email      IS 'почта';
COMMENT ON COLUMN users.password   IS 'пароль';

COMMENT ON COLUMN users.role       IS 'роль';

COMMENT ON COLUMN users.created_at IS 'дата создания';
COMMENT ON COLUMN users.updated_at IS 'дата обновления';
COMMENT ON COLUMN users.deleted_at IS 'дата удаления';

-- +migrate Down
DROP TABLE users;
DROP TYPE user_role_t;
