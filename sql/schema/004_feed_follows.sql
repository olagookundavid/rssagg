-- +goose Up
CREATE TABLE feed_follows (
    id UUID,
    created_at TIMESTAMP NOT NULL,
    updated_at TIMESTAMP NOT NULL,
    user_id UUID NOT NULL,
    feed_id UUID NOT NULL,
    PRIMARY KEY(id),
    UNIQUE(user_id, feed_id),
    FOREIGN KEY(user_id) REFERENCES users(id) ON DELETE CASCADE,
    FOREIGN KEY(feed_id) REFERENCES feeds(id) ON DELETE CASCADE
);
-- on delete cascade to delete field once that Userid is deleted
-- UNIQUE(user_id, feed_id) tomake sure the combo of both are unique
-- +goose Down
DROP TABLE feed_follows;
