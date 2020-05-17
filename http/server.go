package http

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"log"
)

type Server struct {
	port int
	router *gin.Engine
}

type InitRouterFunction func(group *gin.RouterGroup)

func BuildServer(port int, router *gin.Engine) *Server {
	return &Server{port: port, router: router}
}

func (s *Server) Run() {
	if err := s.router.Run(fmt.Sprintf(":%d", s.port)); err != nil {
		log.Fatalf("Cannot start server %v", err)
	}
}

func (s *Server) RegisterRouter(path string, fn InitRouterFunction) *Server {
	fn(s.router.Group(path))
	return s
}