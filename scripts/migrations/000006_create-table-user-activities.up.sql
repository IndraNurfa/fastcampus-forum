CREATE TABLE IF NOT EXISTS user_activities (
    id INTEGER PRIMARY KEY AUTO_INCREMENT,
    post_id INTEGER NOT NULL,
    user_id BIGINT NOT NULL,
    is_liked BOOLEAN NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    updated_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP,
    created_by LONGTEXT NOT NULL,
    updated_by LONGTEXT NOT NULL,
    CONSTRAINT fk_post_id_user_activity FOREIGN KEY (post_id) REFERENCES posts(id),
    CONSTRAINT fk_user_id_user_activity FOREIGN KEY (user_id) REFERENCES users(id)
)