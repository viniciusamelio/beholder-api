import { desc } from "drizzle-orm";
import { Elysia } from "elysia";
import { db } from "./db";
import { apps } from "./db/schema";
import { decorators } from "./core/decorators";
import { appRouter } from "./features/apps/app_router";

const app = new Elysia()
	.use(decorators)
	.get("/health", async ({ set }) => {
		set.status = 200;
		const allApps = await db.select().from(apps).orderBy(desc(apps.id)).limit(2);
		return {
			message: "Server is healthy",
			apps: allApps,
		};
	})
	.group("/v1", (router) => {
		router.use(appRouter);
		router.get("/", () => "ok")
		return router;
	});


app.listen(process.env.PORT as string, () =>
	console.log(`Server started at ${app.server?.url.origin}`),
);
