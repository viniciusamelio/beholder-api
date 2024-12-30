import Elysia from "elysia";
import { SnowFlakeIdService } from "./services/snowflake_id_service";
import { AppService } from "../features/apps/app_service";

export const decorators = new Elysia()
    .decorate({
        idService: new SnowFlakeIdService(),
        appService: new AppService(new SnowFlakeIdService()),
    })