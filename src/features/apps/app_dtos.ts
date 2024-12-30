import { type Static, t } from "elysia";

export const createAppDto = t.Object({
    title: t.String(),
});
export type CreateAppDto = Static<typeof createAppDto>;

export const updateAppDto = t.Object({
    id: t.String(),
    title: t.String(),
});
export type UpdateAppDto = Static<typeof updateAppDto>;