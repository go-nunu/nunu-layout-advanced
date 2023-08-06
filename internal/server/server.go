package server

import (
	"github.com/gin-gonic/gin"
)

type Server struct {
	ServerHTTP *gin.Engine
	//ServerGRPC *grpc.Server
	//ServerWS   *ws.Server
}

func NewServer(serverHTTP *gin.Engine) *Server {
	return &Server{
		ServerHTTP: serverHTTP,
	}
}
