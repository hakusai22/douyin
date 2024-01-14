package svc

import (
	"douyin/v2/user/app/rpc/internal/config"
	"douyin/v2/user/app/rpc/model"
	"douyin/v2/user/app/rpc/model/redisCache"
	"gorm.io/gorm"
	"log"
)

type ServiceContext struct {
	Config config.Config
	Redis  *redisCache.RedisPool
	Db     *gorm.DB
}

func NewServiceContext(c config.Config) *ServiceContext {
	db, err := model.InitGorm(c.DbConfig)
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	pool := redisCache.NewRedisPool(c)
	conn := pool.NewRedisConn()
	_, err = conn.Do("PING")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
		return nil
	}

	return &ServiceContext{
		Config: c,
		Redis:  pool,
		Db:     db,
	}
}
