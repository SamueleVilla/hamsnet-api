BEGIN;

CREATE TABLE IF NOT EXISTS roles (
    name VARCHAR(255) NOT NULL,
    description TEXT NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT roles_pkey PRIMARY KEY (name)
);

CREATE TABLE IF NOT EXISTS users (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    username VARCHAR(255) NOT NULL,
    email VARCHAR(255) NOT NULL,
    hashed_password VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT users_pkey PRIMARY KEY (id),
    CONSTRAINT users_email_key UNIQUE (email),
    CONSTRAINT users_username_key UNIQUE (username)
);

CREATE TABLE IF NOT EXISTS user_roles (
    user_id uuid NOT NULL,
    role_name VARCHAR(255) NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT user_roles_pkey PRIMARY KEY (user_id, role_name),
    CONSTRAINT user_roles_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id),
    CONSTRAINT user_roles_role_name_fkey FOREIGN KEY (role_name) REFERENCES roles(name)
);

CREATE TABLE IF NOT EXISTS hamster_posts (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    user_id uuid NOT NULL,
    content TEXT NOT NULL,
    image_key VARCHAR(255),
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT hamster_posts_pkey PRIMARY KEY (id),
    CONSTRAINT hamster_posts_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS hamster_post_likes (
    post_id uuid NOT NULL,
    user_id uuid NOT NULL,
    created_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT hamster_post_likes_pkey PRIMARY KEY (post_id, user_id),
    CONSTRAINT hamster_post_likes_post_id_fkey FOREIGN KEY (post_id) REFERENCES hamster_posts(id),
    CONSTRAINT hamster_post_likes_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
);

CREATE TABLE IF NOT EXISTS hamster_post_comments (
    id uuid NOT NULL DEFAULT gen_random_uuid(),
    post_id uuid NOT NULL,
    user_id uuid NOT NULL,
    content TEXT NOT NULL,
    reply_to uuid,
    created_at TIMESTAMP NOT NULL DEFAULT now(),
    updated_at TIMESTAMP NOT NULL DEFAULT now(),

    CONSTRAINT hamster_post_comments_pkey PRIMARY KEY (id),
    CONSTRAINT hamster_post_comments_post_id_fkey FOREIGN KEY (post_id) REFERENCES hamster_posts(id),
    CONSTRAINT hamster_post_comments_user_id_fkey FOREIGN KEY (user_id) REFERENCES users(id)
    CONSTRAINT hamster_post_comments_reply_to_fkey FOREIGN KEY (reply_to) REFERENCES hamster_post_comments(id)
);

-- Create view to get the count of likes for each post
create view hamster_post_likes_view as 
    (select post_id, count(*) as likes_count from hamster_post_likes hpl group by post_id);

-- Create view to get the count of comments for each post
create view hamster_post_comments_view as 
    (select post_id, count(*) as comments_count from hamster_post_comments hpl group by post_id);

-- Create view to get the count of replies for each comment
create view hamster_post_comments_replies_view as 
    (select hpc.reply_to as comment_id, count (*) as replies_count 
    from hamster_post_comments hpc, hamster_post_comments hpr
	where hpc.reply_to = hpr.id
	group by hpc.reply_to)


-- Create view to get the post with the author, likes count and comments count
create view hamster_post_feed as (select hp.id,
hp.user_id as author_id,
u.username as author, 
hp.content, hp.image_key, 
hp.created_at, 
coalesce(hpl.likes_count , 0) as likes_count, coalesce(hpc.comments_count, 0) as comments_count 
	from hamster_posts hp 
	inner join users u on hp.user_id = u.id 
	left join hamster_post_likes_view hpl on hp.id = hpl.post_id
	left join hamster_post_comments_view hpc on hp.id = hpc.post_id
    order by hpl.likes_count desc);
);


CREATE OR REPLACE FUNCTION is_post_liked(post_id uuid, user_id uuid)
RETURNS bool AS $$
BEGIN
  RETURN EXISTS (
    SELECT 1
    FROM hamster_post_likes hpl
    WHERE hpl.post_id = $1
      AND hpl.user_id = $2
  );
END;
$$ LANGUAGE plpgsql;


COMMIT;