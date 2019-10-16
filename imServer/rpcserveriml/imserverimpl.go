package rpcserveriml

import (
	"encoding/json"
	"github.com/micro/go-micro/broker"
	"github.com/roggen-yang/IMService/common/errors"
	pb "github.com/roggen-yang/IMService/imServer/protos"
	"github.com/roggen-yang/IMService/imServer/server"
	"golang.org/x/net/context"
	"sync"
)

type ImRpcServerIml struct {
	sync.Mutex
	publisherServerMap map[string]*server.RabbitMqBroker
}

func NewImRpcServerIml(publisherServerMap map[string]*server.RabbitMqBroker) *ImRpcServerIml {
	return &ImRpcServerIml{publisherServerMap: publisherServerMap}
}

func (i *ImRpcServerIml) PublishMessage(ctx context.Context, req *pb.PublishMessageRequest, rsp *pb.PublishMessageResponse) error {
	body, err := json.Marshal(req)
	if err != nil {
		return errors.PublishMessageErr
	}
	key := req.ServerName + req.Topic
	publisher := i.publisherServerMap[key]
	publisher.Publisher(&broker.Message{
		Body: body,
	})
	return nil
}
