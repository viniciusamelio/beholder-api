import Elysia from "elysia";
import { createAppDto } from "./app_dtos";
import { decorators } from "@/core";

export const appRouter = new Elysia<"apps">({ prefix: "apps", precompile: true })
    .use(decorators)
    .post("/", async ({ set, body, appService }) => {
        const app = await appService.create(body);
        set.status = 201;
        return app;
    }, { body: createAppDto })