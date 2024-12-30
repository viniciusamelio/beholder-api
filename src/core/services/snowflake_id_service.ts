import { SnowflakeId } from "@akashrajpurohit/snowflake-id";

export const SnowFlakeIdService = {
	generate: () => {
		const snowflake = SnowflakeId();
		return snowflake.generate();
	},
};
