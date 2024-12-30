import { relations, sql } from "drizzle-orm";
import { int, sqliteTable, text } from "drizzle-orm/sqlite-core";

const timestamps = {
	createdAt: int("created_at", {
		mode: "timestamp",
	}).default(sql`CURRENT_TIMESTAMP`),
	updatedAt: int("updated_at", {
		mode: "timestamp",
	}),
	deletedAt: int("deleted_at", {
		mode: "timestamp",
	}),
};

export const apps = sqliteTable("apps", {
	id: text("id").primaryKey().notNull(),
	title: text("title").notNull(),
	...timestamps,
});

export const environments = sqliteTable("environments", {
	id: text("id").primaryKey().notNull(),
	appId: int("app_id").notNull(),
	title: text("title").notNull(),
	baseUrl: text("base_url").notNull(),
	color: text("color").notNull(),
	...timestamps,
});

export const endpoints = sqliteTable("endpoints", {
	id: text("id").primaryKey().notNull(),
	environmentId: int("environment_id").notNull(),
	title: text("title").notNull(),
	path: text("url").notNull(),
	method: text("method").notNull().default("GET"),
	...timestamps,
});

export const sessions = sqliteTable("sessions", {
	id: text("id").primaryKey().notNull(),
	environmentId: int("environment_id").notNull(),
	title: text("title").notNull(),
	user: text("user"),
	system: text("system"),
	metadata: text("metadata", { mode: "json" }),
	...timestamps,
});

export const calls = sqliteTable("calls", {
	id: text("id").primaryKey().notNull(),
	sessionId: int("session_id"),
	endpointId: int("endpoint_id").notNull(),
	user: text("user"),
	system: text("system"),
	metadata: text("metadata", { mode: "json" }),
	headers: text("headers", { mode: "json" }),
	...timestamps,
});

export const responses = sqliteTable("responses", {
	id: text("id").primaryKey().notNull(),
	callId: int("call_id").notNull(),
	status: int("status").notNull(),
	headers: text("headers", { mode: "json" }),
	body: text("body"),
	...timestamps,
});

export const appEnvironments = relations(apps, ({ many }) => {
	return {
		environments: many(environments),
	};
});

export const envEndpoints = relations(environments, ({ many }) => {
	return {
		endpoints: many(endpoints),
	};
});

export const envSessions = relations(environments, ({ many }) => {
	return {
		sessions: many(sessions),
	};
});

export const sessionCalls = relations(sessions, ({ many }) => {
	return {
		calls: many(calls),
	};
});

export const endpointCalls = relations(endpoints, ({ many }) => {
	return {
		calls: many(calls),
	};
});

export const callResponse = relations(calls, ({ one }) => {
	return {
		response: one(responses),
	};
});
