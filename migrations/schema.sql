CREATE TABLE IF NOT EXISTS "schema_migration" (
"version" TEXT PRIMARY KEY
);
CREATE UNIQUE INDEX "schema_migration_version_idx" ON "schema_migration" (version);
CREATE TABLE IF NOT EXISTS "models" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"civitai_id" INTEGER,
"name" TEXT,
"description" TEXT NOT NULL,
"allow_no_credit" NUMERIC,
"allow_derivatives" NUMERIC,
"allow_different_license" NUMERIC,
"type" TEXT,
"minor" NUMERIC,
"poi" NUMERIC,
"nsfw" NUMERIC,
"nsfw_level" INTEGER,
"cosmetic" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE sqlite_sequence(name,seq);
CREATE TABLE IF NOT EXISTS "allow_commercial_uses" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"model_id" INTEGER,
"use_type" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_id) REFERENCES models (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "stats" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"model_id" INTEGER,
"download_count" INTEGER,
"favorite_count" INTEGER,
"thumbs_up_count" INTEGER,
"thumbs_down_count" INTEGER,
"comment_count" INTEGER,
"rating_count" INTEGER,
"rating" REAL,
"tipped_amount_count" INTEGER,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_id) REFERENCES models (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "model_version_stats" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"model_version_id" INTEGER,
"download_count" INTEGER,
"thumbs_up_count" INTEGER,
"thumbs_down_count" INTEGER,
"rating_count" INTEGER,
"rating" REAL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_version_id) REFERENCES model_versions (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "creators" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"username" TEXT,
"image" TEXT NOT NULL,
"model_id" INTEGER,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_id) REFERENCES models (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "tags" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"name" TEXT,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL
);
CREATE TABLE IF NOT EXISTS "model_tags" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"model_id" INTEGER,
"tag_id" INTEGER,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_id) REFERENCES models (id) ON DELETE cascade,
FOREIGN KEY (tag_id) REFERENCES tags (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "model_versions" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"trained_words" TEXT,
"stats" TEXT,
"civitai_id" INTEGER,
"model_id" INTEGER,
"index" INTEGER,
"name" TEXT,
"base_model" TEXT,
"base_model_type" TEXT,
"published_at" DATETIME,
"availability" TEXT,
"nsfw_level" INTEGER,
"description" TEXT,
"download_url" TEXT NOT NULL,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_id) REFERENCES models (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "files" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"civitai_id" INTEGER,
"model_version_id" INTEGER,
"size_kb" REAL,
"name" TEXT,
"type" TEXT,
"pickle_scan_result" TEXT NOT NULL,
"pickle_scan_message" TEXT NOT NULL,
"virus_scan_result" TEXT NOT NULL,
"virus_scan_message" TEXT NOT NULL,
"scanned_at" DATETIME NOT NULL,
"metadata" TEXT NOT NULL,
"hashes" TEXT NOT NULL,
"download_url" TEXT NOT NULL,
"primary" NUMERIC,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_version_id) REFERENCES model_versions (id) ON DELETE cascade
);
CREATE TABLE IF NOT EXISTS "images" (
"id" INTEGER PRIMARY KEY AUTOINCREMENT,
"model_version_id" INTEGER,
"url" TEXT,
"nsfw_level" INTEGER,
"width" INTEGER,
"height" INTEGER,
"hash" TEXT NOT NULL,
"type" TEXT,
"has_meta" NUMERIC,
"on_site" NUMERIC,
"created_at" DATETIME NOT NULL,
"updated_at" DATETIME NOT NULL,
FOREIGN KEY (model_version_id) REFERENCES model_versions (id) ON DELETE cascade
);