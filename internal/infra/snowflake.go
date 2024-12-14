package infra

import "github.com/godruoyi/go-snowflake"

func GenerateSnowflakeID() int {
	snowflake.SetMachineID(1)
	return int(snowflake.ID())
}
