CREATE EXTENSION IF NOT EXISTS "uuid-ossp";

CREATE TABLE IF NOT EXISTS users (
   "id" uuid PRIMARY KEY NOT NULL DEFAULT (uuid_generate_v4()),
   "username" text UNIQUE NOT NULL,
   "email" text UNIQUE NOT NULL,
   "password" text NOT NULL
);