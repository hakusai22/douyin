package svc

import (
	"douyin/v2/api/internal/config"
	"douyin/v2/jwt/app/rpc/jwtrpc"
	"douyin/v2/user/app/rpc/userrpc"
	"douyin/v2/video/app/rpc/videorpc"
	"github.com/segmentio/kafka-go"
	"github.com/zeromicro/go-zero/zrpc"
	"time"
)

type ServiceContext struct {
	Config      config.Config
	UserRpc     userrpc.UserRpc
	JwtRpc      jwtrpc.JwtRpc
	KafkaWriter *kafka.Writer
	VideoRpc    videorpc.VideoRpc
}

func NewServiceContext(c config.Config) *ServiceContext {
	return &ServiceContext{
		Config:  c,
		UserRpc: userrpc.NewUserRpc(zrpc.MustNewClient(c.UserRpc)),
		JwtRpc:  jwtrpc.NewJwtRpc(zrpc.MustNewClient(c.JwtRpc)),
		KafkaWriter: getKafkaWriter(c.KafkaConfig.Host,
			c.KafkaConfig.Topic,
			c.KafkaConfig.BatchTimeout,
			c.KafkaConfig.BatchSize,
			c.KafkaConfig.BatchBytes,
		),
		VideoRpc: videorpc.NewVideoRpc(zrpc.MustNewClient(c.VideoRpc)),
	}
}

func getKafkaWriter(host, topic string, timeout int, size int, bytes int64) *kafka.Writer {
	return &kafka.Writer{
		Addr:         kafka.TCP(host),
		Topic:        topic,
		Balancer:     &kafka.LeastBytes{},
		BatchTimeout: time.Millisecond * time.Duration(timeout),
		BatchSize:    size,
		BatchBytes:   bytes,
	}
}
