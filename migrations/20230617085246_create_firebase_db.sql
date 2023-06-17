-- Create "firebase_token" table
CREATE TABLE "public"."firebase_token" ("id" bigserial NOT NULL, "token" text NOT NULL, "cid" uuid NOT NULL, "active" boolean NOT NULL, "expiresat" timestamp NOT NULL, PRIMARY KEY ("id"));
-- Create index "firebase_token_cid_token_idx" to table: "firebase_token"
CREATE INDEX "firebase_token_cid_token_idx" ON "public"."firebase_token" ("cid", "token");
