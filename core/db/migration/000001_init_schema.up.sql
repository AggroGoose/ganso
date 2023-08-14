CREATE TABLE "users" (
  "id" text PRIMARY KEY,
  "verified" boolean NOT NULL DEFAULT (FALSE),
  "banned" boolean NOT NULL DEFAULT (FALSE),
  "username" varchar UNIQUE,
  "image" text,
  "url" text,
  "url_verified" boolean,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "permissions" (
  "id" bigserial PRIMARY KEY,
  "title" varchar UNIQUE NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now())
);

CREATE TABLE "comments" (
  "id" bigserial PRIMARY KEY,
  "user_id" text NOT NULL,
  "post_id" text NOT NULL,
  "edited" boolean NOT NULL DEFAULT (FALSE),
  "date_time" timestamptz NOT NULL DEFAULT (now()),
  "content" text NOT NULL
);

CREATE TABLE "replies" (
  "id" bigserial PRIMARY KEY,
  "user_id" text NOT NULL,
  "comment_id" bigserial NOT NULL,
  "edited" boolean NOT NULL DEFAULT (false),
  "date_time" timestamptz NOT NULL DEFAULT (now()),
  "content" text NOT NULL
);

CREATE TABLE "posts" (
  "id" text PRIMARY KEY,
  "slug" varchar UNIQUE,
  "audio_url" text
);

CREATE TABLE "post_likes" (
  "user_id" text NOT NULL,
  "post_id" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "post_saves" (
  "user_id" text NOT NULL,
  "post_id" text NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "post_id")
);

CREATE TABLE "user_permissions" (
  "user_id" text NOT NULL,
  "permission_id" bigserial NOT NULL,
  "created_at" timestamptz NOT NULL DEFAULT (now()),
  PRIMARY KEY ("user_id", "permission_id")
);

CREATE INDEX ON "posts" ("slug");

ALTER TABLE "comments" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "comments" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "replies" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "replies" ADD FOREIGN KEY ("comment_id") REFERENCES "comments" ("id") ON DELETE CASCADE;

ALTER TABLE "post_likes" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "post_likes" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "post_saves" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "post_saves" ADD FOREIGN KEY ("post_id") REFERENCES "posts" ("id") ON DELETE CASCADE;

ALTER TABLE "user_permissions" ADD FOREIGN KEY ("user_id") REFERENCES "users" ("id") ON DELETE CASCADE;

ALTER TABLE "user_permissions" ADD FOREIGN KEY ("permission_id") REFERENCES "permissions" ("id") ON DELETE CASCADE;