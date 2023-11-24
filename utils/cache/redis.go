package cache

import (
	"sampleprj/config"

	"github.com/redis/go-redis/v9"
)

func InitRedis(c config.AppConfig) *redis.Client {
	// rdb := redis.NewClient(
	// 	&redis.Options{
	// 		Addr:     fmt.Sprintf("%s:%d", c.RHOST, c.RPORT),
	// 		Password: c.RPASSWORD,
	// 		DB:       0,
	// 	},
	// )

	// if rdb == nil {
	// 	return nil
	// }
	// return rdb
	return nil
}
