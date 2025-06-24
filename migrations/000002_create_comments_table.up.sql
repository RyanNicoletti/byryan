CREATE TABLE IF NOT EXISTS comments (
    id uuid DEFAULT gen_random_uuid(),
    name text NOT NULL,
    website text,
    content text NOT NULL,
    post_id uuid NOT NULL,
    FOREIGN KEY (post_id) REFERENCES posts(id) ON DELETE CASCADE,
    created timestamp(0) with time zone NOT NULL DEFAULT NOW()
);