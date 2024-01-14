package svc

import (
	"douyin/v2/video/app/kafka/internal/config"
	"douyin/v2/video/app/kafka/model"
	"douyin/v2/video/app/kafka/model/redisCache"
	"github.com/aliyun/aliyun-oss-go-sdk/oss"
	"github.com/segmentio/kafka-go"
	"gorm.io/gorm"
	"log"
	"strings"
)

type ServiceContext struct {
	Config      config.Config
	Oss         *oss.Client
	Redis       *redisCache.RedisPool
	Db          *gorm.DB
	KafkaReader *kafka.Reader
}

func NewServiceContext(c config.Config) *ServiceContext {

	db, err := model.InitGorm(c.DbConfig)
	if err != nil {
		log.Fatalln(err)
	}

	pool := redisCache.NewRedisPool(c)
	conn := pool.NewRedisConn()
	_, err = conn.Do("PING")
	defer conn.Close()
	if err != nil {
		log.Fatalln(err)
	}

	reader := getKafkaReader(c.KafkaConfig.Host, c.KafkaConfig.Topic, c.KafkaConfig.MinBytes, c.KafkaConfig.MaxBytes)
	err = reader.SetOffset(kafka.LastOffset)
	if err != nil {
		log.Fatalln(err)
	}
	return &ServiceContext{
		Config:      c,
		Db:          db,
		Redis:       pool,
		KafkaReader: reader,
	}
}

func getKafkaReader(kafkaURL, topic string, minBytes, maxBytes int) *kafka.Reader {
	brokers := strings.Split(kafkaURL, ",")
	return kafka.NewReader(kafka.ReaderConfig{
		Brokers:  brokers,
		Topic:    topic,
		MinBytes: minBytes,
		MaxBytes: maxBytes,
	})
}
