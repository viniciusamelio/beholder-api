CREATE TABLE `apps` (
	`id` text PRIMARY KEY NOT NULL,
	`title` text NOT NULL,
	`created_at` integer DEFAULT CURRENT_TIMESTAMP,
	`updated_at` integer,
	`deleted_at` integer
);
--> statement-breakpoint
CREATE TABLE `calls` (
	`id` text PRIMARY KEY NOT NULL,
	`session_id` integer,
	`endpoint_id` integer NOT NULL,
	`user` text,
	`system` text,
	`metadata` text,
	`headers` text,
	`created_at` integer DEFAULT CURRENT_TIMESTAMP,
	`updated_at` integer,
	`deleted_at` integer
);
--> statement-breakpoint
CREATE TABLE `endpoints` (
	`id` text PRIMARY KEY NOT NULL,
	`environment_id` integer NOT NULL,
	`title` text NOT NULL,
	`url` text NOT NULL,
	`method` text DEFAULT 'GET' NOT NULL,
	`created_at` integer DEFAULT CURRENT_TIMESTAMP,
	`updated_at` integer,
	`deleted_at` integer
);
--> statement-breakpoint
CREATE TABLE `environments` (
	`id` text PRIMARY KEY NOT NULL,
	`app_id` integer NOT NULL,
	`title` text NOT NULL,
	`base_url` text NOT NULL,
	`color` text NOT NULL,
	`created_at` integer DEFAULT CURRENT_TIMESTAMP,
	`updated_at` integer,
	`deleted_at` integer
);
--> statement-breakpoint
CREATE TABLE `responses` (
	`id` text PRIMARY KEY NOT NULL,
	`call_id` integer NOT NULL,
	`status` integer NOT NULL,
	`headers` text,
	`body` text,
	`created_at` integer DEFAULT CURRENT_TIMESTAMP,
	`updated_at` integer,
	`deleted_at` integer
);
--> statement-breakpoint
CREATE TABLE `sessions` (
	`id` text PRIMARY KEY NOT NULL,
	`environment_id` integer NOT NULL,
	`title` text NOT NULL,
	`user` text,
	`system` text,
	`metadata` text,
	`created_at` integer DEFAULT CURRENT_TIMESTAMP,
	`updated_at` integer,
	`deleted_at` integer
);
