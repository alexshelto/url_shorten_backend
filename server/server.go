package server;

import (
    "github.com/gin-gonic/gin"
)

type Server struct {
    Router *gin.Engine
}

func NewServer(router *gin.Engine) *Server {
    return &Server{
        Router: router,
    }
}

func (s *Server) Start(address string) {
    s.Router.Run(address)
}
