import { SnowflakeId } from "@akashrajpurohit/snowflake-id";
import type { IdService } from "../protocols";


// @Service<IdService>()
export class SnowFlakeIdService implements IdService {
	generate(): string {
		return SnowflakeId().generate();
	}

}
