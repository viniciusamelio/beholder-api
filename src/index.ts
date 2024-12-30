import { desc } from "drizzle-orm";
import { Elysia } from "elysia";
import { db } from "./db";
import { apps } from "./db/schema";

const app = new Elysia()
	.get("/health", async ({ set }) => {
		set.status = 200;
		const allApps = await db.select().from(apps).orderBy(desc(apps.id)).limit(2);
		return {
			message: "Server is healthy",
			apps: allApps,
		};
	})
	.group("/v1", (router) => {
		return router;
	});


app.listen(process.env.PORT as string, () =>
	console.log(`🦊 Server started at ${app.server?.url.origin}`),
);
