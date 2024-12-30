import type { IdService } from "@/core/protocols";
import { db } from "@/db";
import { apps } from "@/db/schema";
import type { CreateAppDto } from "@/features/apps/app_dtos";

export class AppService {
    constructor(
        private readonly idService: IdService
    ) { }
    async create(input: CreateAppDto) {
        const app = await db.insert(apps).values({
            id: this.idService.generate(),
            createdAt: new Date(),
            ...input,
        }).returning();
        return app;
    }
}