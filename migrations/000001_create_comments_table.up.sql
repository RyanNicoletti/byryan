CREATE TABLE IF NOT EXISTS comments (
    id uuid DEFAULT gen_random_uuid(),
    name text NOT NULL,
    website text,
    content text NOT NULL,
    post_slug text NOT NULL,
    created timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    PRIMARY KEY (id)
);

CREATE INDEX idx_comments_post_slug ON comments(post_slug);
CREATE INDEX idx_comments_created ON comments(created DESC);