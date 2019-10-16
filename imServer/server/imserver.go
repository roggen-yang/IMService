package server

import (
	"encoding/json"
	"fmt"
	"github.com/go-acme/lego/log"
	"github.com/gorilla/websocket"
	"github.com/micro/go-micro/broker"
	"github.com/roggen-yang/IMService/common/errors"
	"github.com/roggen-yang/IMService/imServer/protocol"
	"net/http"
	"sync"
	"time"
)

type ImServerOptions func(im *ImServer)

type ImServer struct {
	rabbitMqBroker *RabbitMqBroker
	clients        map[string]*websocket.Conn
	Address        string
	lock           sync.Mutex
	upgrade        *websocket.Upgrader
}

func NewImServer(rabbitMqBroker *RabbitMqBroker, opts ImServerOptions) (*ImServer, error) {
	if err := broker.Init(); err != nil {
		return nil, err
	}

	if err := broker.Connect(); err != nil {
		return nil, err
	}

	imServer := &ImServer{
		rabbitMqBroker: rabbitMqBroker,
		clients:        make(map[string]*websocket.Conn, 0),
		upgrade: &websocket.Upgrader{
			HandshakeTimeout: 1024,
			ReadBufferSize:   1024,
			WriteBufferSize:  0,
			CheckOrigin: func(r *http.Request) bool {
				return true
			},
		},
	}

	if opts != nil {
		opts(imServer)
	}

	if imServer.Address == "" {
		imServer.Address = protocol.DefaultAddress
	}

	return imServer, nil
}

func (i *ImServer) SendMsg(r *protocol.SendMsgRequest) (*protocol.SendMsgResponse, error) {
	i.lock.Lock()
	defer i.lock.Unlock()

	log.Printf("send SendMsgRequest %+v\n", r)
	conn := i.clients[r.ToToken]
	if conn == nil {
		return nil, errors.UserNoLoginErr
	}

	r.Timestamp = time.Now().Unix()
	r.RemoteAddress = conn.RemoteAddr().String()
	bodyMsg, err := json.Marshal(r)
	if err != nil {
		return nil, errors.SendMessageErr
	}

	err = conn.WriteMessage(websocket.TextMessage, bodyMsg)
	if err != nil {
		log.Println("send message err %v", err)
		i.clients[r.ToToken] = nil
		return nil, err
	}

	log.Println("send message succ %v", r.Body)
	return &protocol.SendMsgResponse{}, nil
}

func (i *ImServer) Subscribe() {
	i.rabbitMqBroker.Subscribe(func(msg []byte) error {
		r := new(protocol.SendMsgRequest)
		err := json.Unmarshal(msg, r)
		if err != nil {
			log.Println("[Unmarshal msg err]: %+v", err)
			return err
		}

		i.SendMsg(r)
		if err != nil {
			log.Println("[SendMsg err]: %+v", err)
			return err
		}
		log.Println("had Subscribe msg %+v", string(msg))
		return nil
	})
}

func (i *ImServer) Run() {
	log.Println("websocket has listens at ", i.Address)
	http.HandleFunc(protocol.WebSocketPrefix, i.login)
	log.Fatal(http.ListenAndServe(i.Address, nil))
}

func (i *ImServer) login(w http.ResponseWriter, r *http.Request) {
	conn, err := i.upgrade.Upgrade(w, r, nil)
	if err != nil {
		log.Println(err)
		return
	}
	msgType, message, err := conn.ReadMessage()
	if err != nil {
		log.Println("read login message err ", err)
		return
	}
	if msgType != websocket.TextMessage {
		log.Println("read login msgType err ", err)
		return
	}
	fmt.Println(string(message))
	loginMsgRequest := new(protocol.LoginRequest)
	if err := json.Unmarshal(message, loginMsgRequest); err != nil {
		log.Println("json.Unmarshal msg err ", err)
		return
	}

	i.clients[loginMsgRequest.Token] = conn
	return
}
