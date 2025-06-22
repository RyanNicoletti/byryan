CREATE TABLE IF NOT EXISTS posts (
    id uuid DEFAULT gen_random_uuid(),
    title text NOT NULL,
    slug text NOT NULL,
    tags text[],
    created timestamp(0) with time zone NOT NULL DEFAULT NOW(),
    updated timestamp(0) with time zone NOT NULL DEFAULT NOW()
);